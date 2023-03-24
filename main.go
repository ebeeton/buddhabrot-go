package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/ebeeton/buddhabrot-go/buddhabrot"
	"github.com/ebeeton/buddhabrot-go/parameters"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Register a validator for plot parameters.
	validate := validator.New()
	if err := validate.RegisterValidation("validateStops", parameters.ValidateStops); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting.")
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

			img := buddhabrot.Plot(plot)

			// Persist the plot and parameters.
			if id, err := insert(plot, img); err != nil {
				log.Printf("Failed to connect to database: %v", err)
			} else {
				log.Printf("Insert returned %d.", id)
			}

			if err := writeImage(w, img); err != nil {
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

func writeImage(w http.ResponseWriter, img *image.RGBA) error {
	buf := new(bytes.Buffer)

	if err := png.Encode(buf, img); err != nil {
		return err
	}

	w.Header().Set("Content-type", "image/png")
	w.Header().Set("Content-length", strconv.Itoa(buf.Len()))
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}
