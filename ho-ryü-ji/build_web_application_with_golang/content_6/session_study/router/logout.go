package router

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ho-ryue-ji/session_study/db"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout method:", r.Method)
	if r.Method == "POST" {
		sessionid, err := r.Cookie("sessionid")
		if err != nil {
			log.Fatal(err)
		}
		sessionid.MaxAge = -1
		http.SetCookie(w, sessionid)

		db.Delete(sessionid.Value)

		t, _ := template.ParseFiles("pages/login.gtpl")
		t.Execute(w, nil)
	}
}
