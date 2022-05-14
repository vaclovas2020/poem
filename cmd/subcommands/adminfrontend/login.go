package adminfrontend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/net/xsrftoken"
	"webimizer.dev/poem/oauth"
	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type loginTemplateParams struct {
	PageTitle      string // page title
	LoginActionUrl string // login form action url
	CopyrightText  string // footer copyright text
	XsrfToken      string // secure xsrf_token
	Message        string // form error message (optional)
}

func httpNotAllowFunc(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, "Bad Request")
}

func (p *adminFrontendCmd) generateNewTokenAndShowLogin(session *sessions.Session, rw http.ResponseWriter, r *http.Request) {
	var (
		hashKey   = []byte(p.hashKey)
		cryptoKey = []byte(p.cryptoKey)
	)
	xsrf := xsrftoken.Generate(p.hashKey, session.ID, "oauth")
	session.Values["xsrf_token"] = xsrf
	secureXsrf, err := securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Encode("xsrf_token", xsrf)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	obj := &loginTemplateParams{
		PageTitle:      "Login | Poem CMS",
		LoginActionUrl: "/login",
		CopyrightText:  "Copyright © 2022 Vaclovas Lapinskis",
		XsrfToken:      secureXsrf,
	}
	for _, v := range session.Flashes() {
		obj.Message = v.(string)
	}
	err = session.Save(r, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Cache-Control", "no-store, must-revalidate")
	rw.Header().Set("Pragma", "no-cache")
	output, err := runtime.TemplateParse(templates, "template/login.html", obj)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(rw, output)
}

func (p *adminFrontendCmd) addLoginPageHandler() error {
	http.Handle("/login", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			var (
				hashKey   = []byte(p.hashKey)
				cryptoKey = []byte(p.cryptoKey)
			)
			session, err := store.Get(r, "sid")
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			if v, found := session.Values["userLoggedIn"].(bool); found && v {
				http.Redirect(rw, r, "/", http.StatusFound)
				return
			}
			webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				p.generateNewTokenAndShowLogin(session, rw, r)
			})
			webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				err = r.ParseForm()
				if err != nil {
					http.Error(rw, err.Error(), http.StatusBadRequest)
					return
				}
				token := r.FormValue("xsrf_token")
				username := r.FormValue("username")
				password := r.FormValue("password")
				if token == "" || username == "" || password == "" {
					session.AddFlash("Please enter all form fields")
					p.generateNewTokenAndShowLogin(session, rw, r)
					return
				}
				var realToken string
				securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Decode("xsrf_token", token, &realToken)
				if realToken == "" {
					http.Error(rw, "cannot decode xsrf_token value", http.StatusBadRequest)
					return
				}
				if _, exists := session.Values["xsrf_token"]; !exists {
					http.Error(rw, "no xsrf_token value stored in current session", http.StatusBadRequest)
					return
				}
				sessToken := session.Values["xsrf_token"].(string)
				if sessToken != realToken {
					http.Error(rw, "xsrf_token value is not valid for this session: "+sessToken+" "+realToken, http.StatusBadRequest)
					return
				}
				valid := xsrftoken.Valid(realToken, p.hashKey, session.ID, "oauth")
				if valid {
					response, err := p.grpcAuthUser(&oauth.AuthRequest{Username: username, Password: password, Role: oauth.UserRole_admin})
					if err != nil {
						session.AddFlash(err.Error())
						p.generateNewTokenAndShowLogin(session, rw, r)
						return
					}
					if response.Success {
						session, err = session.Store().New(r, "sid")
						if err != nil {
							http.Error(rw, err.Error(), http.StatusInternalServerError)
							return
						}
						session.Values["userLoggedIn"] = true
						session.Values["username"] = response.User.Name
						session.Values["role"] = response.User.Role.String()
						session.Save(r, rw)
						rw.Header().Set("Cache-Control", "no-store, must-revalidate")
						rw.Header().Set("Pragma", "no-cache")
						http.Redirect(rw, r, "/", http.StatusFound)
						return
					} else {
						session.AddFlash("Invalid username or password")
						p.generateNewTokenAndShowLogin(session, rw, r)
						return
					}
				} else {
					http.Error(rw, "xsrf_token expired", http.StatusBadRequest)
					return
				}
			})
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET", "POST"},                         // define allowed methods
	}.Build())
	return nil
}
