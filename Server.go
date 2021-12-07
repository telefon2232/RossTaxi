package main

import (
	"net/http"
)

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"html/template"
)

type Hello struct {
	Hi  string
	Rol string
}

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}

	return userName
}
func getRole(request *http.Request) (role string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			role = cookieValue["role"]
		}
	}
	return role
}

func setSession(userName string, password string, role string, response http.ResponseWriter) {
	value := map[string]string{
		"name":     userName,
		"password": password,
		"role":     role,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler

func loginHandler(response http.ResponseWriter, request *http.Request) {

	name := request.FormValue("name")
	pass := request.FormValue("password")
	role := request.FormValue("who")

	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, pass, role, response)
		redirectTarget = "/internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

// index page

func indexPageHandler(response http.ResponseWriter, request *http.Request) {

	var indexPage, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	indexPage.Execute(response, "")
}

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	role := getRole(request)
	page := ""
	if userName != "" {
		if role == "I'm client" {
			page = "StatusForUser.html"
		} else if role == "I'm taxi driver" {
			page = "StatusForDriver.html"
		}
		if userName != "" {
			var indexPage, err = template.ParseFiles(page)
			if err != nil {
				fmt.Println(err)
				return
			}
			indexPage.Execute(response, Hello{userName, role})
		} else {
			http.Redirect(response, request, "/", 302)
		}

	}

}

var router = mux.NewRouter()



func main() {
	// Connect to DataBase

	//
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}
