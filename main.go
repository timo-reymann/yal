package main

import (
	_ "embed"

	"github.com/timo-reymann/yal/cmd"
)

//go:embed NOTICE
var noticeContent []byte

func main() {
	cmd.Run(noticeContent)
}
