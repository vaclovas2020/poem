package adminfrontend

import (
	"fmt"
	"net/http"

	"webimizer.dev/poem/runtime"
)

func errorMsg(rw http.ResponseWriter, err error, status int) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	msgHtml, err := runtime.TemplateParse(templates, "template/error.html", err.Error())
	if err != nil {
		http.Error(rw, err.Error(), status)
		return
	}
	rw.WriteHeader(status)
	fmt.Fprint(rw, msgHtml)
}
