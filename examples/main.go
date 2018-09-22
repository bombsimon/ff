package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bombsimon/ff"
)

func main() {
	recursive := flag.Bool("recursive", false, "parse file recursive")

	flag.Parse()

	var f []ff.Match

	switch flag.NArg() {
	case 0:
		f, _ = ff.FilesFromPattern(".", "*", *recursive)
	default:
		f, _ = ff.FilesFromPattern(".", flag.Args()[0], *recursive)
	}

	format := "%-50s | %s\n"

	fmt.Printf(format, "Path", "Filename")
	fmt.Println(strings.Repeat("-", 70))

	for _, x := range f {
		fmt.Printf(format, x.Path, x.Filename)
	}
}
