package cmd

import (
	_ "embed"
	"flag"
	"github.com/timo-reymann/yal/pkg/buildinfo"
	"github.com/timo-reymann/yal/pkg/compress"
	"github.com/timo-reymann/yal/pkg/config"
	"github.com/timo-reymann/yal/pkg/templating"
	"log"
	"net/http"
	"os"
)

func renderTemplate(templateFile string) ([]byte, error) {
	c, err := config.Load()
	if err != nil {
		return nil, err
	}

	return templating.RenderTemplate(templateFile, c)
}

func Run() {
	var templateFile = flag.String("template-file", "builtin", "Template file to render, builtin uses the bundled one with yal")
	var render = flag.Bool("render", false, "Render to output and exit")
	var serve = flag.Bool("serve", false, "Render and serve on HTTP")
	var renderOutput = flag.String("output", "templated.html", "File to render to if -render is specified, use - to render to stdout")
	flag.Parse()

	println("---------------------------------------------")
	buildinfo.PrintVersionInfo()
	println("---------------------------------------------")

	if *render && *serve {
		log.Fatal("-render and -serve are mutually exclusive")
	} else if !*render && !*serve {
		*serve = true
	}

	rendered, err := renderTemplate(*templateFile)
	if err != nil {
		log.Fatalf("Failed to render template: %s", err.Error())
	}

	if *render {
		log.Println("Rendering " + *renderOutput)

		if *renderOutput == "-" {
			_, _ = os.Stdout.Write(rendered)
			return
		}

		templatedFile, err := os.Create(*renderOutput)
		if err != nil {
			log.Fatalf("Failed to write templated file: %s", err.Error())
		}
		defer templatedFile.Close()

		_, err = templatedFile.Write(rendered)
		if err != nil {
			log.Fatalf("Failed to write templated content to file: %s", err.Error())
		}
	}

	if *serve {
		renderedGz, err := compress.GzBytes(rendered)
		if err != nil {
			log.Fatalf("Failed to compress HTML page for HTTP server: %s", err.Error())
		}

		log.Println("Starting server on 0.0.0.0:" + config.Port())
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Content-Encoding", "gzip")
			_, _ = w.Write(renderedGz)
		})

		log.Fatal(http.ListenAndServe(":"+config.Port(), nil))
	}
}
