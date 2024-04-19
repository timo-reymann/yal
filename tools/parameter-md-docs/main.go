package main

import (
	"flag"
	"fmt"
	"github.com/timo-reymann/yal/pkg/config"
	"strings"
)

func main() {
	c := config.Get()
	_ = c.Load()

	println("| Environment variable | Flag | Default | Description |")
	println("| :------------------- | :--- | :------ | :---------- |")

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf(
			"| %-25s | %-20s | %s | %s |\n",
			"YAL_"+strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_")),
			"--"+f.Name,
			fmt.Sprintf("`%s`", f.DefValue),
			strings.Split(f.Usage, "(env:")[0],
		)
	})
}
