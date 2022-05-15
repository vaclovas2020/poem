package adminfrontend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/net/xsrftoken"
	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/poems"
	"webimizer.dev/poem/runtime"
	"webimizer.dev/webimizer"
)

type categoriesTemplateParams struct {
	CategoriesTitle string                    // categories page link title
	PoemsTitle      string                    // poems page link title
	PageTitle       string                    // page title
	HomeTitle       string                    // home page title
	LogoutTitle     string                    // logout title
	CopyrightText   string                    // footer copyright text
	UserEmail       string                    // current user email
	Message         string                    // form error message
	Categories      map[int32]*poems.Category // Categories map
	SubmitButton    string                    // category create form submit button text
	ActionUrl       string                    // category create form action url
	XsrfToken       string                    // category create form xsrf token hidden field
}

func (p *adminFrontendCmd) renderCategoriesPage(session *sessions.Session, rw http.ResponseWriter, r *http.Request) {
	var (
		hashKey   = []byte(p.hashKey)
		cryptoKey = []byte(p.cryptoKey)
	)
	xsrf := xsrftoken.Generate(p.hashKey, session.ID, "add_category")
	session.Values["xsrf_token"] = xsrf
	secureXsrf, err := securecookie.New(*reverseBytes(hashKey), *reverseBytes(cryptoKey)).Encode("xsrf_token", xsrf)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Cache-Control", "no-store, must-revalidate")
	rw.Header().Set("Pragma", "no-cache")
	obj := &categoriesTemplateParams{
		HomeTitle:       "Dashboard",
		LogoutTitle:     "Logout",
		CategoriesTitle: "Categories",
		PoemsTitle:      "Poems",
		PageTitle:       "Categories | Poem CMS",
		CopyrightText:   "Copyright © 2022 Vaclovas Lapinskis",
		UserEmail:       session.Values["email"].(string),
		SubmitButton:    "Add new category",
		ActionUrl:       "/categories",
		XsrfToken:       secureXsrf,
	}
	for _, v := range session.Flashes() {
		obj.Message = v.(string)
	}
	categories, err := p.grpcGetCategories(&poems.CategoriesRequest{Status: poems.CategoriesRequest_PUBLISHED, UserId: session.Values["userId"].(int64)})
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	obj.Categories = categories.Categories
	err = session.Save(r, rw)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	output, err := runtime.TemplateParse(templates, "template/categories.html", obj)
	if err != nil {
		errorMsg(rw, err, http.StatusInternalServerError)
		return
	}
	fmt.Fprint(rw, output)
}

func (p *adminFrontendCmd) addCategoriesPageHandler() {
	http.Handle("/categories", webimizer.HttpHandlerStruct{
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
					p.renderCategoriesPage(session, rw, r)
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
					name := r.FormValue("name")
					category_id := r.FormValue("category_id")
					if (token == "" || name == "") && action == "create" {
						session.AddFlash("Please enter category name")
						p.renderCategoriesPage(session, rw, r)
						return
					}
					if (token == "" || name == "" || category_id == "") && action == "update" {
						session.AddFlash("Please enter category name")
						p.renderCategoriesPage(session, rw, r)
						return
					}
					if (token == "" || category_id == "") && action == "delete" {
						session.AddFlash("Please enter category Id")
						p.renderCategoriesPage(session, rw, r)
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
					valid := xsrftoken.Valid(realToken, p.hashKey, session.ID, "add_category")
					if valid {
						switch action {
						case "create":
							response, err := p.grpcAddCategory(&admin.AdminCategory{Name: name, Slug: fmt.Sprintf("%s-%s", name, uuid.NewString()), Status: admin.AdminCategory_PUBLISHED, UserId: session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot add new category")
							}
							p.renderCategoriesPage(session, rw, r)
							break
						case "delete":
							categoryId, err := strconv.ParseInt(category_id, 10, 32)
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							response, err := p.grpcDeleteCategory(&admin.DeleteCategoryRequest{CategoryId: int32(categoryId), UserId: session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot delete category")
							}
							p.renderCategoriesPage(session, rw, r)
							break
						case "update":
							categoryId, err := strconv.ParseInt(category_id, 10, 32)
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							response, err := p.grpcEditCategory(&admin.AdminCategoryEdit{Name: name, Slug: runtime.GenerateSlug(fmt.Sprintf("%s-%s", name, uuid.NewString())), CategoryId: int32(categoryId), UserId: session.Values["userId"].(int64)})
							if err != nil {
								errorMsg(rw, err, http.StatusInternalServerError)
								return
							}
							if !response.Success {
								session.AddFlash("Cannot update category")
							}
							p.renderCategoriesPage(session, rw, r)
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
