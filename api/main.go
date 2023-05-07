package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"

	"github.com/ebeeton/buddhabrot-go/shared/database"
	"github.com/ebeeton/buddhabrot-go/shared/files"
	"github.com/ebeeton/buddhabrot-go/shared/models"
	"github.com/ebeeton/buddhabrot-go/shared/parameters"
	"github.com/ebeeton/buddhabrot-go/shared/queue"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"

	"gorm.io/gorm"
)

var validate *validator.Validate
var db *gorm.DB

func main() {
	log.Println("API starting.")

	// Connect to MySQL and migrate the plots table.
	var err error
	if db, err = database.Connect(); err != nil {
		log.Fatal(err)
	} else if err = db.AutoMigrate(&models.Plot{}); err != nil {
		log.Fatal(err)
	}

	// Register a validator for plot parameters.
	validate = validator.New()
	if err := validate.RegisterValidation("validateStops", ValidateStops); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.POST("/api/plots", plotRequest)
	router.GET("/api/plots/:id", getImage)
	router.GET("/api/healthcheck", healthcheck)
	router.PanicHandler = handlePanic

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

func handlePanic(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
}

func plotRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Processing request.")
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	var plot parameters.Plot
	if err := d.Decode(&plot); err != nil {
		log.Panic(err)
	} else if err := validate.Struct(plot); err != nil {
		log.Printf("Plot parameter failed validation: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a record in the database so we have an ID to return to the
	// caller.
	b, err := json.Marshal(plot)
	if err != nil {
		log.Panic(err)
	}
	p := models.Plot{
		Params: string(b),
	}
	if r := db.Create(&p); r.Error != nil {
		log.Panic(r.Error)
	}

	// Enqueue the plot request.
	req := parameters.PlotRequest{
		Id:   p.ID,
		Plot: plot,
	}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(req); err != nil {
		log.Panic(err)
	}
	queue.Enqueue(buf.Bytes())
	log.Println("Request queued.")

	// Set the response location header with the ID.
	ids := strconv.FormatUint(uint64(p.ID), 10)
	if l, err := url.JoinPath(r.URL.String(), ids); err != nil {
		log.Panic(err)
	} else {
		w.Header().Add("Location", l)
	}
	w.WriteHeader(http.StatusCreated)

	// Write the ID to the response.
	resp := struct {
		Id uint
	}{
		Id: p.ID,
	}
	json.NewEncoder(w).Encode(resp)
}

func getImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id uint
	if i, err := strconv.ParseUint(p.ByName("id"), 10, 32); err != nil {
		log.Panic(err)
	} else {
		id = uint(i)
	}

	var plot models.Plot
	if r := db.First(&plot, id); r.Error != nil {
		log.Panic(r.Error)
	} else if plot.Filename == "" {
		// This is not an error condition. A plot has been requested but hasn't
		// completed yet.
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Image ID %d hasn't completed yet.", id)
	} else if b, err := files.Read(plot.Filename); err != nil {
		log.Panic(err)
	} else {
		w.Header().Set("Content-type", "image/png")
		w.Header().Set("Content-length", strconv.Itoa(len(b)))
		if _, err := w.Write(b); err != nil {
			log.Panic(err)
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

func healthcheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	resp := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(resp)
}
