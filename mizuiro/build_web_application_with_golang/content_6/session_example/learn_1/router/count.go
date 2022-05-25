package router

import (
	"github.com/mizuirorivi/session_study/session"
	"html/template"
	"net/http"
	"time"
)

/**
show the  value of "countnum" in session
*/
func count(w http.ResponseWriter, r *http.Request) {
	// session start
	sess := session.GetGlobalSesion().SessionStart(w, r)
	// get session value of "createtime"
	createtime := sess.Get("createtime")
	if createtime == nil {
		//set create time to session
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		// expire after 360 seconds
		session.GetGlobalSesion().SessionDestroy(w, r)
		// create new session
		sess = session.GetGlobalSesion().SessionStart(w, r)
	}
	// get session value of "countnum"
	ct := sess.Get("countnum")
	if ct == nil {
		// set value of "countnum "
		sess.Set("countnum", 1)
	} else {
		// increase countnum
		sess.Set("countnum", ct.(int)+1)
	}
	t, _ := template.ParseFiles("count.gtpl")
	// set http response header
	w.Header().Set("Content-Type", "text/html")
	// execute template
	t.Execute(w, sess.Get("countnum"))
}
