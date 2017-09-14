package actions

import (
	"fmt"
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
		app.HandleFunc("/people", PeopleList).Methods("GET")

		app.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(assets)))
	}
	return app
}

func handleError(res http.ResponseWriter, err error) {
	res.WriteHeader(500)
	fmt.Fprint(res, err)
}
