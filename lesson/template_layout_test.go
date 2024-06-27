package lesson

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Page struct {
	Title string
	Name  string
}

//go:embed templates/layout/*.html
var layout embed.FS

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(layout, "templates/layout/*.html"))
	t.ExecuteTemplate(w, "content", Page{
		Title: "Template Layout",
		Name:  "Wawan",
	})
}

func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL, nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, req)

	res := rec.Result()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
