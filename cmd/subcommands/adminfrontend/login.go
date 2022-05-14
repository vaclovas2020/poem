package adminfrontend

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/securecookie"
	"golang.org/x/net/xsrftoken"
	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type loginTemplateParams struct {
	PageTitle      string // page title
	LoginActionUrl string // login form action url
	CopyrightText  string // footer copyright text
	XsrfToken      string // secure xsrf_token
}

func httpNotAllowFunc(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, "Bad Request")
}

func (p *adminFrontendCmd) addLoginPageHandler() error {
	obj := &loginTemplateParams{
		PageTitle:      "Login | Poem CMS",
		LoginActionUrl: "/login",
		CopyrightText:  "Copyright © 2022 Vaclovas Lapinskis",
	}
	output, err := runtime.TemplateParse(templates, "template/login.html", obj)
	if err != nil {
		return err
	}
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
			webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				xsrf := xsrftoken.Generate(p.hashKey, session.ID, "oauth")
				session.Values["xsrf_token"] = xsrf
				secureXsrf, err := securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Encode("xsrf_token", xsrf)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
				err = session.Save(r, rw)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprint(rw, strings.Replace(output, "$xsrf_token", secureXsrf, 1))
			})
			webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				err = r.ParseForm()
				if err != nil {
					http.Error(rw, err.Error(), http.StatusBadRequest)
					return
				}
				token := r.FormValue("xsrf_token")
				if token == "" {
					http.Error(rw, "no xsrf_token provided", http.StatusBadRequest)
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
					fmt.Fprintln(rw, "xsrf_token is valid")
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
