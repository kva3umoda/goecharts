package templates

import (
	"embed"
	"io"

	"github.com/pkg/errors"
)

//go:embed *.tmpl
var templateFS embed.FS

const (
	TemplateBase   = "base.tmpl"
	TemplateChart  = "chart.tmpl"
	TemplateHeader = "header.tmpl"
	TemplatePage   = "page.tmpl"
)

func Load(name string) (string, error) {
	f, err := templateFS.Open(name)
	if err != nil {
		return "", errors.Wrapf(err, "open file :%s", name)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return "", errors.Wrap(err, "read file")
	}

	return string(buf), nil
}

func LoadUnsafe(name string) string {
	template, err := Load(name)
	if err != nil {
		panic(err.Error())
	}
	return template
}
