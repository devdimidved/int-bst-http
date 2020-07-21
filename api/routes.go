package api

import (
	"net/http"
)

func (app *application) Routes() http.Handler {
	app.router.Use(app.setRequestID)
	app.router.Use(app.logRequest)
	app.router.HandleFunc("/search", app.searchHandler).Methods("GET")
	app.router.HandleFunc("/insert", app.insertHandler).Methods("POST")
	app.router.HandleFunc("/delete", app.deleteHandler).Methods("DELETE")
	return app.router
}
