package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

			// Enqueue the plot.
			req := new(bytes.Buffer)
			enc := gob.NewEncoder(req)
			if err := enc.Encode(plot); err != nil {
				log.Fatal(err)
			}
			queue.Enqueue(req.Bytes())
			log.Println("Request queued.")

			w.WriteHeader(http.StatusCreated)
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
