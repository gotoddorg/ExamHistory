package actions

import (
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/monoculum/formam"
	"github.com/soma/examhistory/models"
	"google.golang.org/appengine"
)

func PersonCreate(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	person := models.Person{
		Address: models.Address{},
	}

	err := req.ParseForm()
	if err != nil {
		log.Errorf(ctx, "could not parse form: %v", err)
	}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "form"})
	if err := dec.Decode(req.Form, &person); err != nil {
		log.Errorf(ctx, "could not decode form: %v", err)
	}

	t := Template{Names: []string{"people/show.html", "application.html"}}

	err = t.Render(res, map[string]interface{}{
		"person": person,
	})
	if err != nil {
		log.Errorf(ctx, "could not render template: %v", err)
	}
}
