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
	HomeTitle       string // home page title
	LogoutTitle     string // logout title
	CopyrightText   string // footer copyright text
}

func (p *adminFrontendCmd) addHomePageHandler() error {
	http.Handle("/", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "sid")
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			session.Save(r, rw)
			if v, found := session.Values["userLoggedIn"].(bool); found && v {
				rw.Header().Set("Cache-Control", "no-store, must-revalidate")
				rw.Header().Set("Pragma", "no-cache")
				obj := &homeTemplateParams{
					HomeTitle:       "Dashboard",
					LogoutTitle:     "Logout",
					CategoriesTitle: "Categories",
					PoemsTitle:      "Poems",
					PageTitle:       "Admin dashboard | Poem CMS",
					CopyrightText:   "Copyright © 2022 Vaclovas Lapinskis",
				}
				output, err := runtime.TemplateParse(templates, "template/home.html", obj)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprint(rw, output)
				return
			} else {
				rw.Header().Set("Cache-Control", "no-store, must-revalidate")
				rw.Header().Set("Pragma", "no-cache")
				http.Redirect(rw, r, "/login", http.StatusFound)
				return
			}
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET"},                                 // define allowed methods
	}.Build())
	return nil
}
