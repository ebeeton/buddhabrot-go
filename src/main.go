package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ebeeton/buddhalbrot-go/parameters"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			d := json.NewDecoder(r.Body)
			d.DisallowUnknownFields()
			var plot parameters.RgbPlot
			if err := d.Decode(&plot); err != nil {
				log.Fatal(err)
			}

			b, err := json.Marshal(plot)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(b)
		}

	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
