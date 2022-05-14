package adminfrontend

import (
	"net/http"

	"webimizer.dev/webimizer"
)

func (p *adminFrontendCmd) addLogoutHandler() error {
	http.Handle("/logout", webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "sid")
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			if v, found := session.Values["userLoggedIn"].(bool); found && v {
				session, err = session.Store().New(r, "sid")
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
				session.Save(r, rw)
				http.Redirect(rw, r, "/login", http.StatusMovedPermanently)
				return
			} else {
				http.Redirect(rw, r, "/login", http.StatusMovedPermanently)
				return
			}
		}), // webimizer.HttpHandler call only if method is allowed
		NotAllowHandler: webimizer.HttpNotAllowHandler(httpNotAllowFunc), // webimizer.HtttpNotAllowHandler call if method is not allowed
		AllowedMethods:  []string{"GET"},                                 // define allowed methods
	}.Build())
	return nil
}
