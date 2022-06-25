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

func CookieSetting(w http.ResponseWriter) (sessionid string) {
	session_id := generateCode(20)
	c := &http.Cookie{
		Name:   "sessionid",
		Value:  session_id,
		Path:   "/",
		MaxAge: 3600,
	}

	http.SetCookie(w, c)

	return session_id
}

func DbSetting(r *http.Request, sessionid string) {
	user := &db.User{
		SessionId: sessionid,
		Name:      r.Form["username"][0],
		Password:  r.Form["password"][0],
		Life:      1,
	}
	err := user.Set()
	if err != nil {
		log.Fatal(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register method:", r.Method)
	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("pages/register.gtpl")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		sessionid := CookieSetting(w)
		DbSetting(r, sessionid)

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
