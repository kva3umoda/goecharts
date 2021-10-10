package render

import (
	"fmt"
	"html/template"
	"io"
)

const (
	ModChart = "chart"
	ModPage  = "page"
)

// Renderer
// Any kinds of charts have their render implementation and
// you can define your own render logic easily.
type Renderer interface {
	Render(w io.Writer) error
}

// MustTemplate creates a new template with the given name and parsed contents.
func MustTemplate(name string, contents []string) *template.Template {
	tpl := template.Must(template.New(name).Parse(contents[0])).
		Funcs(
			template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			})

	for _, cont := range contents[1:] {
		tpl = template.Must(tpl.Parse(cont))
	}
	return tpl
}
