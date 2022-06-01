package router

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/ho-ryue-ji/session_study/db"
)

func generateCode(num int) string {
	rand.Seed(time.Now().Unix())
	seed := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r := make([]rune, num)
	for i := range r {
		r[i] = seed[rand.Intn(len(seed))]
	}
	return string(r)
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("pages/register.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		session_id := generateCode(20)
		c := &http.Cookie{
			Name:   "sessionid",
			Value:  session_id,
			Path:   "/",
			MaxAge: 3600,
		}

		http.SetCookie(w, c)

		user := &db.User{
			SessionId: session_id,
			Name:      r.Form["username"][0],
			Password:  r.Form["password"][0],
		}
		err := user.Set()
		if err != nil {
			log.Fatal(err)
		}

		t, _ := template.ParseFiles("pages/login.gtpl")
		t.Execute(w, nil)
	}
}
