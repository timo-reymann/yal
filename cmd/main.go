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

	printSeparator()
	buildinfo.PrintVersionInfo()
	printSeparator()

	if c.Options.IsRenderAndServe() {
		log.Fatal("Render and serve mode are mutually exclusive and can not be used together.")
	} else if !*c.Options.Render && !*c.Options.Serve {
		*c.Options.Serve = true
	}

	rendered, err := templating.RenderTemplate(*c.Options.TemplateFile, c)
	if err != nil {
		log.Fatalf("Failed to render template: %s.", err.Error())
	}

	if *c.Options.Render {
		render(c, rendered)
	}

	if *c.Options.Serve {
		serve(c, rendered)
	}
}

func printSeparator() {
	println("---------------------------------------------")
}

func serve(c *config.Configuration, rendered []byte) {
	renderedGz, err := compress.GzBytes(rendered)
	if err != nil {
		log.Fatalf("Failed to compress HTML page for HTTP server: %s.", err.Error())
	}

	log.Println("Starting server on 0.0.0.0:" + *c.Port + ".")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Encoding", "gzip")
		_, _ = w.Write(renderedGz)
	})

	log.Fatal(http.ListenAndServe(":"+*c.Port, nil))
}

func render(c *config.Configuration, rendered []byte) {
	log.Println("Rendering " + *c.Options.RenderOutput)

	if c.Options.IsStdoutOutput() {
		_, _ = os.Stdout.Write(rendered)
		return
	}

	templatedFile, err := os.Create(*c.Options.RenderOutput)
	if err != nil {
		log.Fatalf("Failed to write templated file: %s", err.Error())
	}
	defer templatedFile.Close()

	_, err = templatedFile.Write(rendered)
	if err != nil {
		log.Fatalf("Failed to write templated content to file: %s", err.Error())
	}
	return
}
