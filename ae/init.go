package ae

import (
	"net/http"

	"github.com/soma/examhistory/actions"
)

func init() {
	http.Handle("/", actions.App())
}
