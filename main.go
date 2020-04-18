package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"time"
)

var rd *render.Render

type User struct {
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	CreatedAt	time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Jihun", Email: "ghun2ee@gmail.com"}
	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
	//w.Header().Add("Content-type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(user)
	//fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Jihun", Email: "ghun2ee@gmail.com"}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New(render.Options{
		Directory: "template",
		Extensions: []string{".html", ".tmpl"},
		Layout: "hello",
	})
	mx := pat.New()

	mx.Get("/users", getUserInfoHandler)
	mx.Post("/users", addUserHandler)
	mx.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mx)

	log.Fatal(http.ListenAndServe(":3000", n))
}
