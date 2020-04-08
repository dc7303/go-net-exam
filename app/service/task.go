package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

type FooHandler struct{}

func (f *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

// get parameter exam
func Bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.URL.Query().Get("name"))
}
