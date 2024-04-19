package templating

import (
	"github.com/timo-reymann/yal/pkg/config"
	"os"
	"testing"
)

var c = config.Get()

func init() {
	_ = c.Load()
}

func TestRenderTemplate_Default(t *testing.T) {
	os.Chdir("../..")

	templated, err := RenderTemplate("builtin", c)
	if err != nil {
		t.Fatal(err)
	}
	if templated == nil {
		t.Fatal("Expected template to be rendered for default")
	}
}
