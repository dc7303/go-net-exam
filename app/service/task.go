package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
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
	w.WriteHeader(http.StatusOK)
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

// file server handler
func FileServer() http.Handler {
	return http.FileServer(http.Dir("public"))
}

// upload action
func UploadsHanlder(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()
	curPath, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	dirname := path.Join(curPath, "../../uploads")
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	defer file.Close()
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}
