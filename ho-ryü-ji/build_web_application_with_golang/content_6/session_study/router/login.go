package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ho-ryue-ji/session_study/db"
)

func LoginBySessionid(w http.ResponseWriter, r *http.Request) (string, error) {
	sessionid, err := r.Cookie("sessionid")
	if err != nil || sessionid == nil {
		return "login", nil
	}

	if sessionid != nil {
		_, err := db.Get("SessionId", sessionid.Value)
		if err != nil {
			if err == sql.ErrNoRows {
				return "", err
			}
			log.Fatal(err)
		}
		return "success", nil
	} else {
		return "", err
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login method:", r.Method)
	switch r.Method {
	case "GET":
		pageid, err := LoginBySessionid(w, r)
		switch pageid {
		case "login":
			t, _ := template.ParseFiles("pages/login.gtpl")
			t.Execute(w, nil)
		case "success":
			http.Redirect(w, r, "/success", http.StatusTemporaryRedirect)
		default:
			log.Fatal(err)
		}
	case "POST":
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		user, err := db.Get("name", r.Form["username"][0])
		if err != nil {
			if err == sql.ErrNoRows {
				t, _ := template.ParseFiles("pages/login.gtpl")
				t.Execute(w, nil)
				return
			}
			log.Fatal(err)
		} else if user.Life == 0 {
			fmt.Println("This accoount has been deleted.")
			t, _ := template.ParseFiles("pages/login.gtpl")
			t.Execute(w, nil)
		} else if user.Password == r.Form["password"][0] {
			sessionid := CookieSetting(w)
			fmt.Println("session_id:", sessionid)
			user.SessionId = sessionid
			db.Update(sessionid, user)
			http.Redirect(w, r, "/success", http.StatusFound)
		}
	}
}
