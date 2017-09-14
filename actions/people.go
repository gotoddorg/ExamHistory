package actions

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/datastore"
	"google.golang.org/appengine/log"

	"github.com/gotoddorg/examhistory/models"
	"github.com/monoculum/formam"
	"google.golang.org/appengine"
)

func PeopleList(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	people := []*models.Person{}

	query := datastore.NewQuery("Person").Order("created")
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		handleError(res, err)
		return
	}

	keys, err := client.GetAll(ctx, query, &people)
	if err != nil {
		handleError(res, err)
		return
	}

	t := Template{Names: []string{"people/index.html", "application.html"}}

	err = t.Render(res, map[string]interface{}{
		"people": people,
		"keys":   keys,
	})

	if err != nil {
		handleError(res, err)
		return
	}
}

func PersonCreate(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	person := &models.Person{
		Address: models.Address{},
	}

	err := req.ParseForm()
	if err != nil {
		fmt.Fprint(res, err)
		return
	}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "form"})
	if err := dec.Decode(req.Form, person); err != nil {
		fmt.Fprint(res, err)
		return
	}

	if err := person.Create(ctx); err != nil {
		fmt.Fprint(res, err)
		return
	}

	t := Template{Names: []string{"people/show.html", "application.html"}}

	err = t.Render(res, map[string]interface{}{
		"person": person,
	})

	if err != nil {
		log.Errorf(ctx, "could not render template: %v", err)
		fmt.Fprint(res, err)
		return
	}
}
