package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ebeeton/buddhabrot-go/shared/database"
	"github.com/ebeeton/buddhabrot-go/shared/parameters"
	"github.com/ebeeton/buddhabrot-go/shared/queue"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

var validate *validator.Validate

func main() {
	log.Println("API starting.")

	// Register a validator for plot parameters.
	validate = validator.New()
	if err := validate.RegisterValidation("validateStops", ValidateStops); err != nil {
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
	id, err := database.Insert(string(b))
	if err != nil {
		log.Fatal(err)
	}

	// Enqueue the plot request.
	req := parameters.PlotRequest{
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

	// Set the response location header with the ID.
	ids := strconv.FormatInt(id, 10)
	if l, err := url.JoinPath(r.URL.String(), ids); err != nil {
		log.Fatal(err)
	} else {
		w.Header().Add("Location", l)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func getImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	i := p.ByName("id")
	if id, err := strconv.ParseInt(i, 10, 64); err != nil {
		log.Fatal(err)
	} else if filename, err := database.GetFilename(id); err != nil {
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

// ValidateStops validates that the state of a slice of Stops.
func ValidateStops(fl validator.FieldLevel) bool {
	// TODO:: How do you add specific error messages?
	stops := fl.Field().Interface().([]parameters.Stop)
	if len(stops) < 2 {
		return false
	} else if stops[0].Position != 0 {
		return false
	} else if stops[len(stops)-1].Position != 1 {
		return false
	}

	return true
}
