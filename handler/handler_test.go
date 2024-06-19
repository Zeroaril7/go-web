package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	mainURL          = "http://localhost:8000"
	pathStart        = "/start"
	pathQuery        = "/test"
	singleQuery      = "?name=aril"
	doubleQuery      = "?name=aril&age=10"
	multiValuesQuery = "?name=Fakhril&name=Ainur"
)

func StartServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Start Server")
}

func TestStart(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathStart, nil)
	rec := httptest.NewRecorder()

	StartServer(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func SayHi(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" {
		fmt.Fprint(w, "Hi Man")
	} else {
		fmt.Fprintf(w, "Hi %s, age %s", name, age)
	}
}

func TestQuerySayHi(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathQuery+doubleQuery, nil)
	rec := httptest.NewRecorder()

	SayHi(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleParamValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultiParamValues(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathQuery+multiValuesQuery, nil)
	rec := httptest.NewRecorder()

	MultipleParamValues(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
