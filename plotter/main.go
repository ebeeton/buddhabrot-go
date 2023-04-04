package main

import (
	"bytes"
	"encoding/gob"
	"image/png"
	"log"

	"github.com/ebeeton/buddhabrot-go/plotter/buddhabrot"
	"github.com/ebeeton/buddhabrot-go/plotter/queue"
)

func main() {
	log.Println("Starting.")
	queue.Dequeue(func(body []byte) {
		r := bytes.NewReader(body)
		dec := gob.NewDecoder(r)
		var req PlotRequest
		if err := dec.Decode(&req); err != nil {
			log.Fatal(err)
		}
		log.Printf("Received plot request: %v", req)

		// Plot the image.
		img := buddhabrot.Plot(req.Plot)

		// Encode a PNG.
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			log.Fatal(err)
		}

		// Write the image to the local filesystem.
		_, err := writePng(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		// TODO:: Update the database with the filename.
	})
}
