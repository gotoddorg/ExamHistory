package ae

import (
	"net/http"

	"github.com/gotoddorg/examhistory/actions"
)

func init() {
	http.Handle("/", actions.App())
}
