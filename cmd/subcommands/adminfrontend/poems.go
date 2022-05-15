package adminfrontend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/net/xsrftoken"
	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/poems"
	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type poemsTemplateParams struct {
	CategoriesTitle string                    // categories page link title
	PoemsTitle      string                    // poems page link title
	PageTitle       string                    // page title
	HomeTitle       string                    // home page title
	LogoutTitle     string                    // logout title
	CopyrightText   string                    // footer copyright text
	UserEmail       string                    // current user email
	Message         string                    // form error message
	Categories      map[int32]*poems.Category // Categories map
	Poems           map[int32]*poems.Poem     // Poems map
	SubmitButton    string                    // poem create form submit button text
	ActionUrl       string                    // poem create form action url
	XsrfToken       string                    // poem create form xsrf token hidden field
}

func (p *adminFrontendCmd) renderPoemsPage(session *sessions.Session, rw http.ResponseWriter, r *http.Request) {
	var (
		hashKey   = []byte(p.hashKey)
		cryptoKey = []byte(p.cryptoKey)
	)
	xsrf := xsrftoken.Generate(p.hashKey, session.ID, "add_poem")
	session.Values["xsrf_token"] = xsrf
	secureXsrf, err := securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Encode("xsrf_token", xsrf)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
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
		SubmitButton:    "Add new poem",
		ActionUrl:       "/poems",
		XsrfToken:       secureXsrf,
	}
	categories, err := p.grpcGetCategories(&poems.CategoriesRequest{
		Status: poems.CategoriesRequest_PUBLISHED,
		UserId: session.Values["userId"].(int64)})
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	obj.Categories = categories.Categories
	poems, err := p.grpcGetPoems(&poems.PoemsRequest{
		UserId: session.Values["userId"].(int64)})
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	obj.Poems = poems.Poems
	for _, v := range session.Flashes() {
		obj.Message = v.(string)
	}
	err = session.Save(r, rw)
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

func (p *adminFrontendCmd) addPoemsPageHandler() {
	http.Handle("/poems", webimizer.HttpHandlerStruct{
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
					p.renderPoemsPage(session, rw, r)
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
					title := r.FormValue("title")
					text := r.FormValue("text")
					category_id := r.FormValue("category_id")
					poem_id := r.FormValue("poem_id")
					if (token == "" || title == "" || category_id == "" || text == "") && action == "create" {
						session.AddFlash("Please enter all form fields")
						p.renderPoemsPage(session, rw, r)
						return
					}
					if (token == "" || title == "" || category_id == "" || text == "" || poem_id == "") && action == "update" {
						session.AddFlash("Please enter all form fields")
						p.renderPoemsPage(session, rw, r)
						return
					}
					if (token == "" || poem_id == "") && action == "delete" {
						session.AddFlash("Please enter all form fields")
						p.renderPoemsPage(session, rw, r)
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
					valid := xsrftoken.Valid(realToken, p.hashKey, session.ID, "add_poem")
					if valid {
						switch action {
						case "create":
							categoryId, err := strconv.ParseInt(category_id, 10, 32)
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							response, err := p.grpcAddPoem(&admin.AdminPoem{
								Title:      title,
								Text:       text,
								CategoryId: int32(categoryId),
								UserId:     session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot add new poem")
							}
							p.renderPoemsPage(session, rw, r)
							break
						case "delete":
							poemId, err := strconv.ParseInt(poem_id, 10, 32)
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							response, err := p.grpcDeletePoem(&admin.DeletePoemRequest{
								PoemId: int32(poemId),
								UserId: session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot delete poem")
							}
							p.renderPoemsPage(session, rw, r)
							break
						case "update":
							categoryId, err := strconv.ParseInt(category_id, 10, 32)
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							poemId, err := strconv.ParseInt(poem_id, 10, 32)
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							response, err := p.grpcEditPoem(&admin.AdminPoemEdit{
								PoemId:     int32(poemId),
								Title:      title,
								Text:       text,
								CategoryId: int32(categoryId),
								UserId:     session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot update poem")
							}
							p.renderPoemsPage(session, rw, r)
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
