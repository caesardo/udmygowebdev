package main

import (
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	First    string
	Last     string
	Password []byte
	Role     string
}

type session struct {
	un         string
	lastactive time.Time
}

var tpl *template.Template

const sessionLength int = 30

var dbSessionsCleaned time.Time

var dbUsers = make(map[string]user)       //user id, user
var dbSessions = make(map[string]session) //session id, session

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()
		username := req.FormValue("username")
		password := req.FormValue("password")
		first := req.FormValue("firstname")
		last := req.FormValue("lastname")
		role := req.FormValue("role")

		//username taken?
		if _, ok := dbUsers[username]; ok {
			http.Error(w, "username already taken", http.StatusForbidden)
			return
		}

		//store user
		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u := user{username, first, last, bs, role}
		dbUsers[username] = u

		//create session cookie
		sid := uuid.NewV4().String()
		c := &http.Cookie{
			Name:   "session",
			Value:  sid,
			Path:   "/",
			MaxAge: sessionLength,
		}
		http.SetCookie(w, c)
		session := session{username, time.Now()}
		dbSessions[sid] = session

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//form process
	if req.Method == http.MethodPost {
		req.ParseForm()
		un := req.FormValue("username")
		pass := req.FormValue("password")

		//check if un exist
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "user not registered", http.StatusForbidden)
			return
		}

		//compare stored pass & pass entered
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pass))
		if err != nil {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		//create session
		sid := uuid.NewV4()
		c := &http.Cookie{
			Name:   "session",
			Value:  sid.String(),
			MaxAge: sessionLength,
		}
		http.SetCookie(w, c)
		dbSessions[sid.String()] = session{un, time.Now()}

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//check session
	c, _ := req.Cookie("session") //no need to check err, already checked above

	//delete session
	delete(dbSessions, c.Value)
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
	return
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	u := getUser(w, req)
	if u.Role != "007" {
		http.Error(w, "only 007 agents allowed", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		return u
	}

	//if user exist, get user
	if s, ok := dbSessions[c.Value]; ok {
		un := s.un
		u = dbUsers[un]
	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		return false
	}

	s := dbSessions[c.Value] //session exists?
	_, ok := dbUsers[s.un]   //user exists given username?

	return ok

}

func cleanSessions() {
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastactive) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
}
