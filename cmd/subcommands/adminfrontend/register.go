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

type registerTemplateParams struct {
	LoginTitle        string // login page title
	RegisterTitle     string // register page title
	HomeTitle         string // home page title
	PageTitle         string // page title
	RegisterActionUrl string // register form action url
	CopyrightText     string // footer copyright text
	XsrfToken         string // secure xsrf_token
	Message           string // form error message (optional)
	EmailField        string // user email input label
	PasswordField     string // user password input label
	RePasswordField   string // user password input label
	SubmitButton      string // submit button text
}

func (p *adminFrontendCmd) generateNewTokenAndShowRegister(session *sessions.Session, rw http.ResponseWriter, r *http.Request) {
	var (
		hashKey   = []byte(p.hashKey)
		cryptoKey = []byte(p.cryptoKey)
	)
	xsrf := xsrftoken.Generate(p.hashKey, session.ID, "oauth")
	session.Values["xsrf_token"] = xsrf
	secureXsrf, err := securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Encode("xsrf_token", xsrf)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	obj := &registerTemplateParams{
		PageTitle:         "Register | Poem CMS",
		RegisterActionUrl: "/register",
		CopyrightText:     "Copyright © 2022 Vaclovas Lapinskis",
		XsrfToken:         secureXsrf,
		EmailField:        "Email",
		PasswordField:     "Password",
		RePasswordField:   "Re-enter password",
		HomeTitle:         "Home",
		LoginTitle:        "Login",
		RegisterTitle:     "Register",
		SubmitButton:      "Register",
	}
	for _, v := range session.Flashes() {
		obj.Message = v.(string)
	}
	err = session.Save(r, rw)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Cache-Control", "no-store, must-revalidate")
	rw.Header().Set("Pragma", "no-cache")
	output, err := runtime.TemplateParse(templates, "template/register.html", obj)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	fmt.Fprint(rw, output)
}

func (p *adminFrontendCmd) addRegisterPageHandler() error {
	http.Handle("/register", webimizer.HttpHandlerStruct{
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
				http.Redirect(rw, r, "/", http.StatusFound)
				return
			}
			webimizer.Get(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				p.generateNewTokenAndShowRegister(session, rw, r)
			})
			webimizer.Post(rw, r, func(rw http.ResponseWriter, r *http.Request) {
				err = r.ParseForm()
				if err != nil {
					errorMsg(rw, err, http.StatusBadRequest)
					return
				}
				token := r.FormValue("xsrf_token")
				email := r.FormValue("email")
				password := r.FormValue("password")
				repassword := r.FormValue("re-password")
				if token == "" || email == "" || password == "" || repassword == "" {
					session.AddFlash("Please enter valid email and password")
					p.generateNewTokenAndShowRegister(session, rw, r)
					return
				}
				if password != repassword {
					session.AddFlash("Password mismatch")
					p.generateNewTokenAndShowRegister(session, rw, r)
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
				valid := xsrftoken.Valid(realToken, p.hashKey, session.ID, "oauth")
				if valid {
					response, err := p.grpcNewUser(&oauth.AuthRequest{Email: email, Password: password, Role: oauth.UserRole_admin})
					if err != nil {
						session.AddFlash(err.Error())
						p.generateNewTokenAndShowRegister(session, rw, r)
						return
					}
					if response.Success {
						session, err = session.Store().New(r, "sid")
						if err != nil {
							errorMsg(rw, err, http.StatusInternalServerError)
							return
						}
						session.Values["userLoggedIn"] = true
						session.Values["email"] = response.User.Email
						session.Values["role"] = response.User.Role.String()
						session.Values["userId"] = response.User.UserId
						session.Save(r, rw)
						rw.Header().Set("Cache-Control", "no-store, must-revalidate")
						rw.Header().Set("Pragma", "no-cache")
						http.Redirect(rw, r, "/", http.StatusFound)
						return
					} else {
						session.AddFlash("User with this email already exists")
						p.generateNewTokenAndShowRegister(session, rw, r)
						return
					}
				} else {
					errorMsg(rw, fmt.Errorf("xsrf_token expired"), http.StatusBadRequest)
					return
				}
			})
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET", "POST"},                         // define allowed methods
	}.Build())
	return nil
}
