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
type Arr struct {
	Test_array []int
}

var user_cache = make([]string , 30)

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

func addressHandler(response http.ResponseWriter, request *http.Request){
	Nickmame := request.FormValue("Nickname")
	AddressFrom := request.FormValue("AddressFrom")
	AddressTo := request.FormValue("AddressTo")
	user_cache := append(user_cache, "client " +Nickmame+ " "+ AddressFrom+" "+ AddressTo)
	fmt.Println(user_cache)


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
	test_array := Arr{}.Test_array
	test_array = append(test_array, 1,23,4,5,5,6,262,6)



	//
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")
	router.HandleFunc("/address", addressHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}