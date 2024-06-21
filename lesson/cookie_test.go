package lesson

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Login-Info"
	cookie.Value = "Success"
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Set Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Login-Info")
	if err != nil {
		fmt.Fprint(w, "Cookie not setted yet")
	} else {
		fmt.Fprint(w, cookie.Value)
	}
}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	cookies := rec.Result().Cookies()

	for _, v := range cookies {
		fmt.Printf("Set cookie %s:%s", v.Name, v.Value)
	}
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Login-Info"
	cookie.Value = "Success"

	req.AddCookie(cookie)

	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
