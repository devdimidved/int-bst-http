package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) searchHandler(w http.ResponseWriter, r *http.Request) {
	val, err := strconv.Atoi(r.URL.Query().Get("val"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok := app.bst.Search(val)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) insertHandler(w http.ResponseWriter, r *http.Request) {
	insertReq := map[string]int{}
	err := json.NewDecoder(r.Body).Decode(&insertReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	val, ok := insertReq["val"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.bst.Insert(val)
	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	val, err := strconv.Atoi(r.URL.Query().Get("val"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.bst.Remove(val)
	w.WriteHeader(http.StatusNoContent)
}
