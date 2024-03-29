package templating

import (
	"github.com/timo-reymann/yal/pkg/config"
	"github.com/timo-reymann/yal/pkg/model"
	"testing"
)

var testConfig = &config.Configuration{
	Sections:      []model.Section{},
	SearchEngines: []model.SearchEngine{},
	Assets: config.Assets{
		Logo:       "",
		Mascot:     "",
		Background: "",
		Favicon:    "",
	},
}

func TestRenderTemplate_Default(t *testing.T) {
	templated, err := RenderTemplate("builtin", testConfig)
	if err != nil {
		t.Fatal(err)
	}
	if templated == nil {
		t.Fatal("Expected template to be rendered for default")
	}
}

func TestRenderTemplate_DefaultFile(t *testing.T) {
	templated, err := RenderTemplate("index.gohtml", testConfig)
	if err != nil {
		t.Fatal(err)
	}
	if templated == nil {
		t.Fatal("Expected template to be rendered for default")
	}
}
