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
	validate := validator.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			log.Println("Processing request.")
			d := json.NewDecoder(r.Body)
			d.DisallowUnknownFields()
			var plot parameters.Plot
			if err := d.Decode(&plot); err != nil {
				log.Println("Decode failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := validate.RegisterValidation("validateStops", parameters.ValidateStops); err != nil {
				log.Println("RegisterValidation failed: ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			} else if err := validate.Struct(plot); err != nil {
				log.Println("Plot parameter failed validation: ", err.Error())
				w.WriteHeader(http.StatusBadRequest)
			}

			img := buddhabrot.Plot(plot)

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
