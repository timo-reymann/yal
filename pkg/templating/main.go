package templating

import (
	"bytes"
	_ "embed"
	"github.com/timo-reymann/yal/pkg/config"
	"os"
	"text/template"
)

//go:embed index.gohtml
var DefaultTemplate string

func RenderTemplate(path string, c *config.Configuration) (string, error) {
	var htmlFile []byte
	var err error
	if path == "builtin" {
		htmlFile = []byte(DefaultTemplate)
	} else {
		htmlFile, err = os.ReadFile(path)
		if err != nil {
			return "", nil
		}
	}

	tmpl := template.New("index")
	tmpl, err = tmpl.Parse(string(htmlFile))
	if err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer
	err = tmpl.Execute(&outputBuffer, c)
	if err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}
