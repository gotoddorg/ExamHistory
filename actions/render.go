package actions

import (
	"html/template"
	"io"

	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var templates = packr.NewBox("../templates")

type Template struct {
	Names []string
}

func (t Template) Render(w io.Writer, data map[string]interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}
	var body string
	var err error
	for _, name := range t.Names {
		s, err := templates.MustString(name)
		if err != nil {
			return errors.WithStack(err)
		}

		body, err = plush.Render(s, plush.NewContextWith(data))
		if err != nil {
			return errors.WithStack(err)
		}
		data["yield"] = template.HTML(body)
	}
	_, err = w.Write([]byte(body))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
