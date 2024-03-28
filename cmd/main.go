package cmd

import (
	_ "embed"
	"flag"
	"github.com/timo-reymann/yal/pkg/config"
	"github.com/timo-reymann/yal/pkg/templating"
	"log"
	"net/http"
	"os"
)

func renderTemplate(templateFile string) (string, error) {
	c, err := config.Load()
	if err != nil {
		return "", nil
	}

	return templating.RenderTemplate(templateFile, c)
}

func Run() {
	var templateFile = flag.String("template-file", "builtin", "Template file to render, builtin uses the bundled one with yal")
	var render = flag.Bool("render", false, "Render to output and exit")
	var serve = flag.Bool("serve", false, "Render and serve on HTTP")
	var renderOutput = flag.String("output", "templated.html", "File to render to if -render is specified, use - to render to stdout")

	flag.Parse()

	if *render && *serve {
		log.Fatal("-render and -serve are mutually exclusive")
	} else if !*render && !*serve {
		*serve = true
	}

	rendered, err := renderTemplate(*templateFile)
	if err != nil {
		log.Fatal(err)
	}

	if *render {
		log.Println("Rendering " + *renderOutput)

		if *renderOutput == "-" {
			_, _ = os.Stdout.WriteString(rendered)
			return
		}

		templatedFile, err := os.Create(*renderOutput)
		if err != nil {
			log.Fatal(err)
		}
		defer templatedFile.Close()

		_, err = templatedFile.WriteString(rendered)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serve {
		log.Println("Starting server on 0.0.0.0:" + config.Port())
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(rendered))
		})

		log.Fatal(http.ListenAndServe(":"+config.Port(), nil))
	}
}
