package api

import (
	"github.com/devdimidved/int-bst-http/bst"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type application struct {
	router *mux.Router
	log    *logrus.Logger
	bst    bst.Service
}

func NewApplication(log *logrus.Logger, bst bst.Service) *application {
	app := &application{
		router: mux.NewRouter(),
		log:    log,
		bst:    bst,
	}
	app.Routes()
	return app
}

func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}
