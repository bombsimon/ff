package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bombsimon/ff"
)

func main() {
	var (
		ignore    string
		recursive bool
	)

	flag.BoolVar(&recursive, "recursive", false, "parse file recursive")
	flag.BoolVar(&recursive, "r", false, "parse file recursive (short)")
	flag.StringVar(&ignore, "ignore", "", "pattern to ignore")
	flag.StringVar(&ignore, "v", "", "pattern to ignore (short)")

	flag.Parse()

	var f []ff.Match

	switch flag.NArg() {
	case 0:
		f, _ = ff.FilesFromPattern(".", "*", ignore, recursive)
	case 1:
		f, _ = ff.FilesFromPattern(".", flag.Args()[0], ignore, recursive)
	default:
		f, _ = ff.FilesFromPattern(flag.Args()[0], flag.Args()[1], ignore, recursive)
	}

	format := "%-20s | %-20s | %s\n"

	fmt.Printf(format, "Path", "Filename", "Full name")
	fmt.Println(strings.Repeat("-", 70))

	for _, x := range f {
		fmt.Printf(format, x.Path, x.Filename, x.FullName())
	}
}
