package router

import (
	"github.com/mizuirorivi/session_study/session"
	"net/http"
	"text/template"
)

/**
show the login page
*/
func login(w http.ResponseWriter, r *http.Request) {
	sess := session.GetGlobalSesion().SessionStart(w, r)
	// parse form data in request
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))

	} else {
		// if not get method
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
