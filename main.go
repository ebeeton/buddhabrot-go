package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ebeeton/buddhalbrot-go/buddhabrot"
	"github.com/ebeeton/buddhalbrot-go/parameters"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
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
			} else if img, err := buddhabrot.Plot(plot); err != nil {
				log.Println("Plot failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else if err := writeImage(w, img); err != nil {
				log.Println("WriteImage failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
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
