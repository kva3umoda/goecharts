package render

import (
	"html/template"

	"github.com/kva3umoda/goecharts/model"
)

type Chart interface {
	//ChartID() string
	Build() *model.Option
	JSONNotEscaped() template.HTML
}

type ChartRender struct {
	Chart
	Initialization
	JSFunctions
}

