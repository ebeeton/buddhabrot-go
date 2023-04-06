package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ebeeton/buddhabrot-go/parameters"
	"github.com/ebeeton/buddhabrot-go/queue"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

var validate *validator.Validate

func main() {
	log.Println("Starting.")

	// Register a validator for plot parameters.
	validate = validator.New()
	if err := validate.RegisterValidation("validateStops", parameters.ValidateStops); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.POST("/", plotRequest)

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
