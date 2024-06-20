package lesson

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	name := r.PostForm.Get("name")
	age := r.PostForm.Get("age")
	fmt.Fprintf(w, "Name %s and age %s", name, age)
}

func TestFormPost(t *testing.T) {
	reqBody := strings.NewReader("name=aril&age=20")

	req := httptest.NewRequest(http.MethodPost, mainURL+pathTest, reqBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	FormPost(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
