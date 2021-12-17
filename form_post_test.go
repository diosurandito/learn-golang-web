package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	// Bisa langsung ambil menggunakan ini
	// firstName := request.PostFormValue("first_name")
	// lastName := request.PostFormValue("last_name")

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)

}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Dio&last_name=Surandito")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(t, "Hello Dio Surandito", string(body), "They Shoul be equal")

}
