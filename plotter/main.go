package main

import (
	"bytes"
	"encoding/json"
	"image/png"
	"log"

	"github.com/ebeeton/buddhabrot-go/plotter/buddhabrot"
	"github.com/ebeeton/buddhabrot-go/plotter/parameters"
	"github.com/ebeeton/buddhabrot-go/plotter/queue"
)

func main() {
	log.Println("Starting.")
	queue.Dequeue(func(plot parameters.Plot) {
		// Plot the image.
		img := buddhabrot.Plot(plot)

		// Encode a PNG.
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			log.Fatal(err)
		}

		// Write the image to the local filesystem.
		filename, err := writePng(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		// Get a JSON representation of the plot to persist in the database.
		b, err := json.Marshal(plot)
		if err != nil {
			log.Fatal(err)
		}
		json := string(b)

		// Persist the plot and parameters.
		if id, err := insert(json, filename); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Insert returned %d.", id)
		}
	})
}
