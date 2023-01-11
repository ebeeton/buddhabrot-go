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
			d := json.NewDecoder(r.Body)
			d.DisallowUnknownFields()
			var plot parameters.RgbPlot
			if err := d.Decode(&plot); err != nil {
				log.Fatal(err)
			}

			img := buddhabrot.Plot(plot)
			if err := WriteImage(w, img); err != nil {
				log.Fatal(err)
			}
		}

	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

// WriteImage writes an image.RGBA to an http.ResponseWriter.
func WriteImage(w http.ResponseWriter, img *image.RGBA) error {
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
