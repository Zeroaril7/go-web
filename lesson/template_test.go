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

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	text := `<html><body>{{.}}</body></html>`

	// Lebih enak menggunakan must karena error sudah di handling
	// t, err := template.New("SIMPLE").Parse(text)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(text))

	t.ExecuteTemplate(w, "SIMPLE", "Hai HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.html"))

	t.ExecuteTemplate(w, "simple.html", "Halo")
}

func TestSimpleHTMLFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateDir(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.html"))

	t.ExecuteTemplate(w, "simple.html", "Hai simple HTML Template")
}

func TestTemplateDir(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateDir(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

//go:embed templates/*.html
var templateFS embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templateFS, "templates/*.html"))

	t.ExecuteTemplate(w, "simple.html", "Hai Embed Template")
}

func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
