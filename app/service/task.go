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
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
	fmt.Fprintf(
		w,
		"%s\n%s\n%s\n%v",
		user.FirstName,
		user.LastName,
		user.Email,
		user.CreatedAt,
	)
}

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

// get parameter exam
func Bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Query())
}
