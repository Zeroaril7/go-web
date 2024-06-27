package lesson

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (p MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + p.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Mamat"}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Aril",
	})
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateFunction(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Aril",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateFunctionGlobal(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateAddFunctionGlobal(w http.ResponseWriter, r *http.Request) {

	t := template.New("FUNCTION")

	t = t.Funcs(map[string]interface{}{
		"lower": func(value string) string {
			return strings.ToLower(value)
		},
	})

	t = template.Must(t.Parse(`{{lower .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "HaI Bro",
	})
}

func TestTemplateAddFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateAddFunctionGlobal(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {

	t := template.New("FUNCTION")

	t = t.Funcs(map[string]interface{}{
		"lower": func(value string) string {
			return strings.ToLower(value)
		},
	})

	t = template.Must(t.Parse(`{{.SayHello "Mamat" | lower}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "HaI Bro",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, mainURL+pathTest, nil)
	rec := httptest.NewRecorder()

	TemplateFunctionPipeline(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}
