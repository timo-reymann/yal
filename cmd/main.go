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

func Run() {
	c := config.Get()
	flag.Parse()
	if err := c.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}

	println("---------------------------------------------")
	buildinfo.PrintVersionInfo()
	println("---------------------------------------------")

	if c.Options.IsRenderAndServe() {
		log.Fatal("-render and -serve are mutually exclusive")
	} else if !*c.Options.Render && !*c.Options.Serve {
		*c.Options.Serve = true
	}

	rendered, err := templating.RenderTemplate(*c.Options.TemplateFile, c)
	if err != nil {
		log.Fatalf("Failed to render template: %s", err.Error())
	}

	if *c.Options.Render {
		log.Println("Rendering " + *c.Options.RenderOuput)

		if c.Options.IsStdoutOutput() {
			_, _ = os.Stdout.Write(rendered)
			return
		}

		templatedFile, err := os.Create(*c.Options.RenderOuput)
		if err != nil {
			log.Fatalf("Failed to write templated file: %s", err.Error())
		}
		defer templatedFile.Close()

		_, err = templatedFile.Write(rendered)
		if err != nil {
			log.Fatalf("Failed to write templated content to file: %s", err.Error())
		}
	}

	if *c.Options.Serve {
		renderedGz, err := compress.GzBytes(rendered)
		if err != nil {
			log.Fatalf("Failed to compress HTML page for HTTP server: %s", err.Error())
		}

		log.Println("Starting server on 0.0.0.0:" + *c.Port)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Content-Encoding", "gzip")
			_, _ = w.Write(renderedGz)
		})

		log.Fatal(http.ListenAndServe(":"+*c.Port, nil))
	}
}
