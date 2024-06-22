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
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Halo")
}

func TestSimpleHTMLFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	SimpleHTMLFile(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateDir(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hai simple HTML Template")
}

func TestTemplateDir(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateDir(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templateFS embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templateFS, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hai Embed Template")
}

func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
