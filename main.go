package main

import (
	"encoding/json"
	"net/http"
	"testing-demo/cars"

	"github.com/go-chi/chi"
)

func main() {

	carService := cars.Service{Repository: cars.NewDummyRepository()}

	r := chi.NewRouter()

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			result, err := carService.List()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			if err = json.NewEncoder(w).Encode(result); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			result, err := carService.Get(chi.URLParam(r, "id"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			if err = json.NewEncoder(w).Encode(result); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			var car cars.Car
			if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			result, err := carService.Create(&car)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			if err = json.NewEncoder(w).Encode(result); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
