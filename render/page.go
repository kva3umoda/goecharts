package render

import (
	"bytes"
	"io"
	"os"
	"regexp"

	"github.com/kva3umoda/goecharts/model"
	"github.com/kva3umoda/goecharts/templates"
)

type Layout string

const (
	PageNoneLayout   Layout = "none"
	PageCenterLayout Layout = "center"
	PageFlexLayout   Layout = "flex"
)

// Charter represents a chart value which provides its type, assets and can be validated.
type Charter interface {
	Build() *model.Option
}

// Page represents a page chart.
type Page struct {
	Initialization
	Assets

	charts []Chart
	Layout Layout
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.Layout = PageCenterLayout
	return page
}

// SetLayout sets the layout of the Page.
func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
	return page
}

// AddCharts adds new charts to the page.
func (page *Page) AddCharts(charts ...Chart) *Page {
	page.charts = append(page.charts, charts...)

	return page
}

// Validate validates the given configuration.
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}

// Render renders the page into the given io.Writer.
func (page *Page) Render(w io.Writer) error {
	page.Validate()
	page.Initialization.Validate()
	contents := []string{
		templates.LoadUnsafe(templates.TemplateHeader),
		templates.LoadUnsafe(templates.TemplateBase),
		templates.LoadUnsafe(templates.TemplatePage),
	}

	tpl := MustTemplate(ModPage, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, ModPage, page); err != nil {
		return err
	}

	pat := regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)
	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}

func (page *Page) Charts() []ChartRender {
	var charts []ChartRender
	for _, chart := range page.charts {
		chartRender := ChartRender{Chart: chart, Initialization: Initialization{}}
		chartRender.Validate()
		charts = append(charts, chartRender)
	}

	return charts
}

func (page *Page) RenderHTML(fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	err = page.Render(io.MultiWriter(f))
	if err != nil {
		return err
	}

	return nil
}
