package lesson

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is empty")
	} else {
		fmt.Fprint(w, "OK")
	}
}
func TestResCode(t *testing.T) {

	for i := 0; i < 2; i++ {
		theURL := mainURL + pathTest
		if i > 0 {
			theURL += singleQuery
		}

		req := httptest.NewRequest(http.MethodGet, theURL, nil)
		rec := httptest.NewRecorder()

		ResCode(rec, req)

		res := rec.Result()

		body, _ := io.ReadAll(res.Body)
		fmt.Println(res.StatusCode)
		fmt.Println(res.Status)
		fmt.Println(string(body))
		fmt.Println()
	}
}
