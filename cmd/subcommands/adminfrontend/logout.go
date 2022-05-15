package adminfrontend

import (
	"net/http"

	"webimizer.dev/webimizer"
)

func (p *adminFrontendCmd) addLogoutHandler() {
	http.Handle("/logout", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "sid")
			if err != nil {
				errorMsg(rw, err, http.StatusInternalServerError)
				return
			}
			if v, found := session.Values["userLoggedIn"].(bool); found && v {
				session.Values = make(map[interface{}]interface{})
				session.Save(r, rw)
				rw.Header().Set("Cache-Control", "no-store, must-revalidate")
				rw.Header().Set("Pragma", "no-cache")
				http.Redirect(rw, r, "/login", http.StatusFound)
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
}
