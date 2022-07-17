package adminfrontend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/net/xsrftoken"
	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type domainTemplateParams struct {
	CategoriesTitle string // categories page link title
	PoemsTitle      string // poems page link title
	PageTitle       string // page title
	HomeTitle       string // home page title
	DomainTitle     string // domain title
	LogoutTitle     string // logout title
	CopyrightText   string // footer copyright text
	UserEmail       string // current user email
	Message         string // form error message
	Domain          string // domain value
	SubmitButton    string // category create form submit button text
	ActionUrl       string // category create form action url
	XsrfToken       string // category create form xsrf token hidden field
}

func (p *adminFrontendCmd) renderDomainPage(session *sessions.Session, rw http.ResponseWriter, r *http.Request) {
	var (
		hashKey   = []byte(p.hashKey)
		cryptoKey = []byte(p.cryptoKey)
	)
	xsrf := xsrftoken.Generate(p.hashKey, session.ID, "add_domain")
	session.Values["xsrf_token"] = xsrf
	secureXsrf, err := securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Encode("xsrf_token", xsrf)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Cache-Control", "no-store, must-revalidate")
	rw.Header().Set("Pragma", "no-cache")
	obj := &domainTemplateParams{
		HomeTitle:       "Dashboard",
		LogoutTitle:     "Logout",
		CategoriesTitle: "Categories",
		PoemsTitle:      "Poems",
		DomainTitle:     "Domain",
		PageTitle:       "Domain | Poem CMS",
		CopyrightText:   "Copyright © 2022 Vaclovas Lapinskis",
		UserEmail:       session.Values["email"].(string),
		SubmitButton:    "Add/update domain",
		ActionUrl:       "/domain",
		XsrfToken:       secureXsrf,
	}
	for _, v := range session.Flashes() {
		obj.Message = v.(string)
	}
	res, err := p.grpcGetDomain(&admin.GetAdminDomain{UserId: session.Values["userId"].(int64)})
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	if res.Success {
		obj.Domain = res.Domain.Domain
	}
	err = session.Save(r, rw)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	output, err := runtime.TemplateParse(templates, "template/domain.html", obj)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	fmt.Fprint(rw, output)
}

func (p *adminFrontendCmd) addDomainPageHandler() {
	http.Handle("/domain", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			var (
				hashKey   = []byte(p.hashKey)
				cryptoKey = []byte(p.cryptoKey)
			)
			session, err := store.Get(r, "sid")
			if err != nil {
				errorMsg(rw, err, http.StatusInternalServerError)
				return
			}
			if v, found := session.Values["userLoggedIn"].(bool); found && v {
				webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
					p.renderDomainPage(session, rw, r)
				})
				webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
					err = r.ParseForm()
					if err != nil {
						errorMsg(rw, err, http.StatusBadRequest)
						return
					}
					action := r.FormValue("action")
					if action == "" {
						errorMsg(rw, fmt.Errorf("action parameter is required"), http.StatusBadRequest)
						return
					}
					token := r.FormValue("xsrf_token")
					domain := r.FormValue("domain")
					if (token == "" || domain == "") && action == "create" {
						session.AddFlash("Please enter all form fields")
						p.renderDomainPage(session, rw, r)
						return
					}
					if (token == "" || domain == "") && action == "delete" {
						session.AddFlash("Please enter all form fields")
						p.renderDomainPage(session, rw, r)
						return
					}
					var realToken string
					securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Decode("xsrf_token", token, &realToken)
					if realToken == "" {
						errorMsg(rw, fmt.Errorf("cannot decode xsrf_token value"), http.StatusBadRequest)
						return
					}
					if _, exists := session.Values["xsrf_token"]; !exists {
						errorMsg(rw, fmt.Errorf("no xsrf_token value stored in current session"), http.StatusBadRequest)
						return
					}
					sessToken := session.Values["xsrf_token"].(string)
					if sessToken != realToken {
						errorMsg(rw, fmt.Errorf("xsrf_token value is not valid for this session"), http.StatusBadRequest)
						return
					}
					valid := xsrftoken.Valid(realToken, p.hashKey, session.ID, "add_domain")
					if valid {
						switch action {
						case "create":
							response, err := p.grpcAddDomain(&admin.AdminDomain{Domain: domain, UserId: session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot add domain")
							}
							p.renderDomainPage(session, rw, r)
							break
						case "delete":
							response, err := p.grpcDeleteDomain(&admin.AdminDomain{Domain: domain, UserId: session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot delete domain")
							}
							p.renderDomainPage(session, rw, r)
							break
						default:
							break
						}
					} else {
						errorMsg(rw, fmt.Errorf("xsrf_token expired"), http.StatusBadRequest)
						return
					}
				})
				return
			} else {
				session.Save(r, rw)
				rw.Header().Set("Cache-Control", "no-store, must-revalidate")
				rw.Header().Set("Pragma", "no-cache")
				http.Redirect(rw, r, "/login", http.StatusFound)
				return
			}
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET", "POST"},                         // define allowed methods
	}.Build())
}
