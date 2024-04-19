package config

import (
	"encoding/json"
	"flag"
	"github.com/timo-reymann/yal/pkg/assets"
	"github.com/timo-reymann/yal/pkg/model"
	"os"
	path2 "path"
	"strings"
)

var (
	trueVal  = true
	falseVal = false
)

type Assets struct {
	LogoPath         *string
	MascotPath       *string
	BackgroundPath   *string
	BackgroundFilter *string
	FaviconPath      *string
}

type Options struct {
	TemplateFile *string
	RenderOuput  *string
	Render       *bool
	Serve        *bool
}

type Configuration struct {
	Sections      []model.Section
	SearchEngines []model.SearchEngine
	Assets        Assets
	Options       Options
	PageTitle     *string
	Port          *string
	ImageFolder   *string
	ConfigFolder  *string
}

func envVarName(name string) string {
	return "YAL_" + strings.ToUpper(strings.ReplaceAll(name, "-", "_"))
}

func stringConfigOpt(name string, fallback string, description string) *string {
	envVar := envVarName(name)
	envVal := os.Getenv(envVar)
	if envVal != "" {
		return &envVal
	}
	return flag.String(name, fallback, description+" (env: "+envVar+")")
}

func boolConfigOpt(name string, fallback bool, description string) *bool {
	envVar := envVarName(name)
	envVal := os.Getenv(envVar)
	if envVal != "" {
		if envVal == "1" || envVal == "true" {
			return &trueVal
		} else if envVal == "0" || envVal == "false" {
			return &falseVal
		}
	}
	return flag.Bool(name, fallback, description+" (env: "+envVar+")")
}

var config *Configuration

// Get config based on file system and environment variables
func Get() *Configuration {
	if config == nil {
		config = &Configuration{
			PageTitle:    stringConfigOpt("page-title", "LinkHub - The place where it just clicks.", "Title of the HTML page generated"),
			Port:         stringConfigOpt("port", "2024", "The HTTP port of the server when run with serve (default)"),
			ImageFolder:  stringConfigOpt("images-folder", "images", "Relative or absolute path where the images reside"),
			ConfigFolder: stringConfigOpt("config-folder", "config", "Relative or absolute path where the configuration files reside"),
			Assets: Assets{
				LogoPath:         stringConfigOpt("logo", "logo", "Basename of a file  without extension (searched in images-folder) or an HTTP url of the image to be used as a logo on the right"),
				MascotPath:       stringConfigOpt("mascot", "mascot", "Basename of a file without extension (searched in images-folder) or an HTTP url of the image to be used as a logo on the left"),
				BackgroundPath:   stringConfigOpt("background", "background", "Basename of a file without extension (searched in images-folder) or an HTTP url of the image to be used as a background image"),
				BackgroundFilter: stringConfigOpt("background-filter", "blur(5px) brightness(0.9)", "CSS Filter to apply to the background image. See [MDN docs](https://developer.mozilla.org/en-US/docs/Web/CSS/filter-function) for more information and examples for the filter CSS function for more information"),
				FaviconPath:      stringConfigOpt("favicon", "favicon", "Basename of a file without extension (searched in images-folder) or an HTTP url of the image to be used as favicon for the page"),
			},
			Options: Options{
				TemplateFile: stringConfigOpt("template-file", "builtin", "Template file to Render, builtin uses the bundled one with yal"),
				RenderOuput:  stringConfigOpt("output", "templated.html", "File to render to if -render is specified, use - to render to stdout"),
				Render:       boolConfigOpt("render", false, "Render to output and exit"),
				Serve:        boolConfigOpt("serve", false, "Render and Serve on HTTP"),
			},
		}
	}

	return config
}

func (c *Configuration) Load() error {
	searchEngines, err := c.loadSearchEngines()
	if err != nil {
		return err
	}
	c.SearchEngines = searchEngines

	sections, err := c.loadSections()
	if err != nil {
		return err
	}
	c.Sections = sections

	return nil
}

func (c *Configuration) loadConfigFile(baseName string, v interface{}) error {
	data, err := os.ReadFile(path2.Join(*c.ConfigFolder, baseName+".json"))
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (c *Configuration) loadSearchEngines() ([]model.SearchEngine, error) {
	var searchEngines []model.SearchEngine
	err := c.loadConfigFile("searchEngines", &searchEngines)
	return searchEngines, err
}

func (c *Configuration) loadSections() ([]model.Section, error) {
	var sections []model.Section
	err := c.loadConfigFile("items", &sections)
	return sections, err
}

func (c *Configuration) loadImage(path *string) (string, error) {
	resolvedPath, err := assets.LookupImgAnyExt(path2.Join(*c.ImageFolder, *path))
	// in case file can not be found locally try to load from URL
	// or search for full file name
	if err != nil {
		return assets.InlineIcon(*path)
	}

	return assets.InlineIcon(resolvedPath)
}

func (c *Configuration) Logo() (string, error) {
	return c.loadImage(c.Assets.LogoPath)
}

func (c *Configuration) Background() (string, error) {
	return c.loadImage(c.Assets.BackgroundPath)
}

func (c *Configuration) Favicon() (string, error) {
	return c.loadImage(c.Assets.FaviconPath)
}

func (c *Configuration) Mascot() (string, error) {
	return c.loadImage(c.Assets.MascotPath)
}

func (o *Options) IsRenderAndServe() bool {
	return *o.Render && *o.Serve
}

func (o *Options) IsStdoutOutput() bool {
	return *o.RenderOuput == "-"
}
