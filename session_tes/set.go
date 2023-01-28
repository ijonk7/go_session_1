package session_tes

import (
	"net/http"
)

func Set(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}
