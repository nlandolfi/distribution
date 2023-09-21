package version

import (
	"fmt"
	"io"
	"os"
)

// FprintVersion outputs the version string to the writer, in the following
// format, followed by a newline:
//
//	<cmd> <project> <version>
//
// For example, a binary "registry" built from github.com/nlandolfi/distribution
// with version "v2.0" would print the following:
//
//	registry github.com/nlandolfi/distribution v2.0
func FprintVersion(w io.Writer) {
	fmt.Fprintln(w, os.Args[0], Package, Version)
}

// PrintVersion outputs the version information, from Fprint, to stdout.
func PrintVersion() {
	FprintVersion(os.Stdout)
}
