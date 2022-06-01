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
	fmt.Println("login method:", r.Method)
	sessionid, err := r.Cookie("sessionid")
	if r.Method == "GET" {
		if err != nil && sessionid == nil {
			ViewLogin(w, r)
			return
		}

		if sessionid != nil {
			_, err := db.Get(sessionid.Value)
			if err != nil {
				if err == sql.ErrNoRows {
					t, _ := template.ParseFiles("pages/login.gtpl")
					t.Execute(w, nil)
					return
				}
				log.Fatal(err)
			}
			http.Redirect(w, r, "/success", http.StatusTemporaryRedirect)
		}

	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		user, err := db.Get(sessionid.Value)
		if err != nil {
			log.Fatal(err)
		} else if user.Name == r.Form["username"][0] && user.Password == r.Form["password"][0] {
			http.Redirect(w, r, "/success", http.StatusFound)
		}
	}
}
