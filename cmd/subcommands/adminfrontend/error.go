package adminfrontend

import (
	"net/http"

	"webimizer.dev/poem/runtime"
)

func errorMsg(rw http.ResponseWriter, err error, status int) {
	msgHtml, err := runtime.TemplateParse(templates, "template/error.html", err.Error())
	if err != nil {
		http.Error(rw, err.Error(), status)
		return
	}
	http.Error(rw, msgHtml, status)
}
