package lesson

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ReqHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestReqHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()

	ReqHeader(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func ResHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-KEY", "admin-123")
	fmt.Fprint(w, "x-api-key")
}

func TestResHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	ResHeader(rec, req)

	header := rec.Header().Get("X-API-KEY")

	fmt.Println(header)
}
