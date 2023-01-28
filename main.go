package main

import (
	"go-session-1/session_tes"
	"log"
	"net/http"
)

//const SESSION_ID = "id"

//func newCookieStore() *sessions.CookieStore {
//	authKey := []byte("my-auth-key-very-secret")
//	encryptionKey := []byte("my-encryption-key-very-secret123")
//
//	store := sessions.NewCookieStore(authKey, encryptionKey)
//	store.Options.Path = "/"
//	store.Options.MaxAge = 86400 * 7
//	store.Options.HttpOnly = true
//
//	return store
//}

func main() {
	http.HandleFunc("/set", session_tes.Set)
	http.HandleFunc("/get", session_tes.Get)
	//http.HandleFunc("/delete", session_tes.Delete)

	log.Print("Server started on: http://localhost:8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}

//func main() {
//	store := newCookieStore()
//
//	e := echo.New()
//
//	e.Use(echo.WrapMiddleware(context.ClearHandler))
//
//	e.GET("/get", func(c echo.Context) error {
//		session, _ := store.Get(c.Request(), SESSION_ID)
//
//		if len(session.Values) == 0 {
//			return c.String(http.StatusOK, "empty result")
//		}
//
//		return c.String(http.StatusOK, fmt.Sprintf(
//			"%s %s",
//			session.Values["message1"],
//			session.Values["message2"],
//		))
//	})
//
//	e.GET("/set", func(c echo.Context) error {
//		session, _ := store.Get(c.Request(), SESSION_ID)
//		session.Values["message1"] = "hello"
//		session.Values["message2"] = "world"
//		session.Save(c.Request(), c.Response())
//
//		return c.Redirect(http.StatusTemporaryRedirect, "/get")
//	})
//
//	e.GET("/delete", func(c echo.Context) error {
//		session, _ := store.Get(c.Request(), SESSION_ID)
//		session.Options.MaxAge = -1
//		session.Save(c.Request(), c.Response())
//
//		return c.Redirect(http.StatusTemporaryRedirect, "/get")
//	})
//}
