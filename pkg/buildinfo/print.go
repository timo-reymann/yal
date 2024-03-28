package buildinfo

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
)

// ErrAbort states that the regular execution should be aborted
var ErrAbort = errors.New("abort")

func addLine(w *tabwriter.Writer, heading string, val string) {
	_, _ = fmt.Fprintf(w, heading+"\t%s\n", val)
}

// PrintVersionInfo prints a tabular list with build info
func PrintVersionInfo() {
	fmt.Printf("yal %s (%s) by Timo Reymann\n", Version, BuildTime)
	println()
	println("Build information")
	w := tabwriter.NewWriter(os.Stderr, 10, 1, 10, byte(' '), tabwriter.TabIndent)
	addLine(w, "GitSha", GitSha)
	addLine(w, "Version", Version)
	addLine(w, "BuildTime", BuildTime)
	addLine(w, "Go-Version", runtime.Version())
	addLine(w, "OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
	_ = w.Flush()
}
