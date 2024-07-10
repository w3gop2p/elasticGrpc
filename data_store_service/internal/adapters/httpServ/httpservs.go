package httpServ

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func writeResponseOK(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, response)
}
func writeResponse(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func (a *Adapter) seedItemHandler(w http.ResponseWriter, r *http.Request) {
	if err := a.api.SeedData(context.Background()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	response := "data successfully transferred into elastic database"
	writeResponseOK(w, response)
}

func (a *Adapter) createItemHandler(w http.ResponseWriter, r *http.Request) {
	if err := a.api.PlaceData(context.Background()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	response := "data successfully transferred into elastic database"
	writeResponseOK(w, response)
}

func (a *Adapter) getAllDocsHandler(w http.ResponseWriter, r *http.Request) {
	response, err := a.api.GetAllData(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	writeResponseOK(w, response)
}

func (a *Adapter) searchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	response, err := a.api.TextSearch(context.Background(), title)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	writeResponseOK(w, response)
}

func (a *Adapter) scrollSearch(w http.ResponseWriter, r *http.Request) {
	from, _ := strconv.Atoi(r.FormValue("from"))
	size, _ := strconv.Atoi(r.FormValue("size"))
	response, err := a.api.ScrollSearch(context.Background(), from, size)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	writeResponseOK(w, response)
}

func (a *Adapter) aggSubcategory(w http.ResponseWriter, r *http.Request) {
	response, err := a.api.AggregateSubcategory(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	writeResponseOK(w, response)
}
