package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ho-ryue-ji/session_study/db"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("pages/login.gtpl")
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		sessionid, err := r.Cookie("sessionid")
		if err != nil && sessionid == nil {
			ViewLogin(w, r)
			return
		}

		if sessionid != nil {
			_, err := db.Get(sessionid.Value)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Redirect(w, r, "/register", http.StatusTemporaryRedirect)
					return
				}
				log.Fatal(err)
			}
		}

		http.Redirect(w, r, "/success", http.StatusTemporaryRedirect)

	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
