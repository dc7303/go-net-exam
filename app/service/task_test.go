package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	Index(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello, world!", string(data))
}

func TestBarWithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	Bar(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello, ", string(data))
}

func TestBar(t *testing.T) {
	assert := assert.New(t)

	name := "margurt"
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name="+name, nil)

	Bar(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello, "+name, string(data))
}

func TestFooWithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", nil)

	foo := FooHandler{}
	foo.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFoo(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(
		"POST", "/foo",
		strings.NewReader(`{"first_name":"margurt", "last_name":"Choe", "email":"dc7303@gmail.com"}`),
	)

	foo := FooHandler{}
	foo.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("margurt", user.FirstName)
	assert.Equal("Choe", user.LastName)
}

func TestUploads(t *testing.T) {
	assert := assert.New(t)
	curPath, _ := os.Getwd()
	mockFile := path.Join(curPath, "../../mock/test.png")
	file, err := os.Open(mockFile)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(mockFile))
	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	UploadsHanlder(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := path.Join(curPath, "../../uploads", filepath.Base(mockFile))
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(mockFile)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}

// func TestUsers(t *testing.T) {
// 	assert := assert.New(t)

// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/users", nil)

// 	UsersHandler(res, req)
// 	assert.Equal(http.StatusOK, res.Code)
// 	data, _ := ioutil.ReadAll(res.Body)
// 	assert.Contains(string(data), "Get UserInfo")
// }

func TestUsersGetInfo(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users/margurt", nil)

	UsersHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	assert.Contains(string(data), "User ID:margurt")
}
