package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"image/png"
	"log"

	"github.com/ebeeton/buddhabrot-go/plotter/buddhabrot"
	"github.com/ebeeton/buddhabrot-go/shared/database"
	"github.com/ebeeton/buddhabrot-go/shared/files"
	"github.com/ebeeton/buddhabrot-go/shared/parameters"
	"github.com/ebeeton/buddhabrot-go/shared/queue"
)

func main() {
	log.Println("Plotter starting.")
	queue.Dequeue(func(body []byte) (err error) {
		defer func() {
			// https://stackoverflow.com/a/25638915/2382333
			if r := recover(); r != nil {
				log.Printf("Recovered: %v", r)
				switch t := r.(type) {
				case error:
					err = t
				default:
					err = errors.New("unknown panic")
				}
			}
		}()
		r := bytes.NewReader(body)
		dec := gob.NewDecoder(r)
		var req parameters.PlotRequest
		if err := dec.Decode(&req); err != nil {
			return err
		}
		log.Printf("Processing plot request: %v", req)

		// Plot the image.
		img := buddhabrot.Plot(req.Plot)

		// Encode a PNG.
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return err
		}

		// Write the image to the local filesystem.
		filename, err := files.Write(buf.Bytes())
		if err != nil {
			return err
		}

		// Update the DB record with the filename.
		if err := database.Update(req.Id, filename); err != nil {
			return err
		}

		return nil
	})
}
