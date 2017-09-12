package actions

import (
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/soma/examhistory/models"
	"google.golang.org/appengine"
)

func Home(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	t := Template{Names: []string{"home.html", "application.html"}}

	data := map[string]interface{}{
		"person": models.Person{
			Address: models.Address{},
		},
	}

	err := t.Render(w, data)
	if err != nil {
		log.Errorf(ctx, "could not render template: %v", err)
	}
}
