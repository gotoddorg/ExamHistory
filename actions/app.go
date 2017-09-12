package actions

import (
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

var app *mux.Router

func App() http.Handler {
	if app == nil {
		assets := packr.NewBox("../assets")
		app = mux.NewRouter()
		app.HandleFunc("/", Home).Methods("GET")
		app.HandleFunc("/people", PersonCreate).Methods("POST")

		app.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(assets)))
	}
	return app
}
