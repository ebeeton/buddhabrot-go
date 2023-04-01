package main

import (
	"bytes"
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/ebeeton/buddhabrot-go/buddhabrot"
	"github.com/ebeeton/buddhabrot-go/parameters"
	"github.com/ebeeton/buddhabrot-go/queue"
	"github.com/go-playground/validator/v10"
)

func main() {
	log.Println("Starting.")

	// Register a validator for plot parameters.
	validate := validator.New()
	if err := validate.RegisterValidation("validateStops", parameters.ValidateStops); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			log.Println("Processing request.")
			d := json.NewDecoder(r.Body)
			d.DisallowUnknownFields()
			var plot parameters.Plot
			if err := d.Decode(&plot); err != nil {
				log.Printf("Decode failed: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else if err := validate.Struct(plot); err != nil {
				log.Printf("Plot parameter failed validation: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Get a JSON representation of the plot. This is a bit redundant
			// given it came in that way, but validation requires a struct.
			b, err := json.Marshal(plot)
			if err != nil {
				log.Fatal(err)
			}
			json := string(b)

			// Enqueue the plot.
			queue.Enqueue(json)

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

			// Persist the plot and parameters.
			if id, err := insert(json, filename); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Insert returned %d.", id)
			}

			if err := writeResponse(w, buf); err != nil {
				log.Println("WriteImage failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			log.Println("Request processed.")
		}
	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func writeResponse(w http.ResponseWriter, buf *bytes.Buffer) error {

	w.Header().Set("Content-type", "image/png")
	w.Header().Set("Content-length", strconv.Itoa(buf.Len()))
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}
