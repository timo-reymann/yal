package templating

import (
	"bytes"
	_ "embed"
	"github.com/timo-reymann/yal/pkg/buildinfo"
	"github.com/timo-reymann/yal/pkg/config"
	"os"
	"text/template"
)

//go:embed index.gohtml
var DefaultTemplate string

func RenderTemplate(path string, c *config.Configuration) ([]byte, error) {
	var htmlFile []byte
	var err error
	if path == "builtin" {
		htmlFile = []byte(DefaultTemplate)
	} else {
		htmlFile, err = os.ReadFile(path)
		if err != nil {
			return nil, err
		}
	}

	tmpl := template.New("index")
	tmpl.Funcs(template.FuncMap{
		"Version": func() string {
			return buildinfo.Version
		},
	})
	tmpl, err = tmpl.Parse(string(htmlFile))
	if err != nil {
		return nil, err
	}

	var outputBuffer bytes.Buffer
	err = tmpl.Execute(&outputBuffer, c)
	if err != nil {
		return nil, err
	}

	return outputBuffer.Bytes(), nil
}
