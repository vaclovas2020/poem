package adminfrontend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type poemsTemplateParams struct {
	CategoriesTitle string // categories page link title
	PoemsTitle      string // poems page link title
	PageTitle       string // page title
	HomeTitle       string // home page title
	LogoutTitle     string // logout title
	CopyrightText   string // footer copyright text
	UserEmail       string // current user email
	Message         string // form error message
}

func (p *adminFrontendCmd) renderPoemsPage(session *sessions.Session, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Cache-Control", "no-store, must-revalidate")
	rw.Header().Set("Pragma", "no-cache")
	obj := &poemsTemplateParams{
		HomeTitle:       "Dashboard",
		LogoutTitle:     "Logout",
		CategoriesTitle: "Categories",
		PoemsTitle:      "Poems",
		PageTitle:       "Poems | Poem CMS",
		CopyrightText:   "Copyright © 2022 Vaclovas Lapinskis",
		UserEmail:       session.Values["email"].(string),
	}
	for _, v := range session.Flashes() {
		obj.Message = v.(string)
	}
	err := session.Save(r, rw)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	output, err := runtime.TemplateParse(templates, "template/poems.html", obj)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	fmt.Fprint(rw, output)
}

func (p *adminFrontendCmd) addPoemsPageHandler() error {
	http.Handle("/poems", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "sid")
			if err != nil {
				errorMsg(rw, err, http.StatusInternalServerError)
				return
			}
			session.Save(r, rw)
			if v, found := session.Values["userLoggedIn"].(bool); found && v {
				p.renderPoemsPage(session, rw, r)
				return
			} else {
				rw.Header().Set("Cache-Control", "no-store, must-revalidate")
				rw.Header().Set("Pragma", "no-cache")
				http.Redirect(rw, r, "/login", http.StatusFound)
				return
			}
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET", "POST"},                         // define allowed methods
	}.Build())
	return nil
}
