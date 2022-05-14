package adminfrontend

import (
	"fmt"
	"net/http"

	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type homeTemplateParams struct {
	CategoriesTitle string // categories page link title
	PoemsTitle      string // poems page link title
	PageTitle       string // page title
	CopyrightText   string // footer copyright text
}

func (p *adminFrontendCmd) addHomePageHandler() error {
	obj := &homeTemplateParams{
		CategoriesTitle: "Categories",
		PoemsTitle:      "Poems",
		PageTitle:       "Login | Poem CMS",
		CopyrightText:   "Copyright © 2022 Vaclovas Lapinskis",
	}
	output, err := runtime.TemplateParse(templates, "template/home.html", obj)
	if err != nil {
		return err
	}
	http.Handle("/", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "sid")
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			if v, found := session.Values["userLoggedIn"]; found {
				valid := v.(bool)
				if valid {
					fmt.Fprint(rw, output)
				}
			} else {
				http.Redirect(rw, r, "/login", http.StatusMovedPermanently)
			}
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET"},                                 // define allowed methods
	}.Build())
	return nil
}
