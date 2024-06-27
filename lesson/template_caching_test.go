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

//go:embed templates/*.html
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.html", "Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateCaching(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
