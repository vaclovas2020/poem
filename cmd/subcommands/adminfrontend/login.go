package adminfrontend

import (
	"fmt"
	"net/http"

	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type LoginTemplateParams struct {
	PageTitle      string // page title
	LoginActionUrl string // login form action url
	CopyrightText  string // footer copyright text
}

func httpNotAllowFunc(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, "Bad Request")
}

func (p *adminFrontendCmd) addLoginPageHandler() error {
	obj := &LoginTemplateParams{
		PageTitle:      "Login | Poem CMS",
		LoginActionUrl: "/login",
		CopyrightText:  "Copyright &copy; 2022 Vaclovas Lapinskis",
	}
	output, err := runtime.TemplateParse(templates, "template/login.html", obj)
	if err != nil {
		return err
	}
	http.Handle("/login", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				fmt.Fprint(rw, output)
			})
			webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				fmt.Fprint(rw, "HTTP POST method was used. Not implemented yet")
			})
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET", "POST"},                         // define allowed methods
	}.Build())
	return nil
}
