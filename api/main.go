package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ebeeton/buddhabrot-go/parameters"
	"github.com/ebeeton/buddhabrot-go/shared/queue"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

var validate *validator.Validate

func main() {
	log.Println("API starting.")

	// Register a validator for plot parameters.
	validate = validator.New()
	if err := validate.RegisterValidation("validateStops", parameters.ValidateStops); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.POST("/api/plots", plotRequest)
	router.GET("/api/plots/:id", getImage)

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

func plotRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	// Create a record in the database so we have an ID to return to the
	// caller.
	b, err := json.Marshal(plot)
	if err != nil {
		log.Fatal(err)
	}
	id, err := insert(string(b))
	if err != nil {
		log.Fatal(err)
	}

	// Enqueue the plot request.
	req := PlotRequest{
		Id:   id,
		Plot: plot,
	}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(req); err != nil {
		log.Fatal(err)
	}
	queue.Enqueue(buf.Bytes())
	log.Println("Request queued.")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func getImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	i := p.ByName("id")
	if id, err := strconv.ParseInt(i, 10, 64); err != nil {
		log.Fatal(err)
	} else if filename, err := getFilename(id); err != nil {
		log.Fatal(err)
	} else if filename == "" {
		// This is not an error condition. A plot has been requested but hasn't
		// completed yet.
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Image ID %d hasn't completed yet.", id)
	} else if b, err := readPng(filename); err != nil {
		log.Fatal(err)
	} else {
		w.Header().Set("Content-type", "image/png")
		w.Header().Set("Content-length", strconv.Itoa(len(b)))
		if _, err := w.Write(b); err != nil {
			log.Fatal(err)
		}
		log.Printf("Returned image ID %d successfully.", id)
	}
}
