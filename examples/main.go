package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bombsimon/ff"
)

func main() {
	recursive := flag.Bool("recursive", false, "parse file recursive")
	ignore := flag.String("ignore", "", "pattern to ignore")

	flag.Parse()

	var f []ff.Match

	switch flag.NArg() {
	case 0:
		f, _ = ff.FilesFromPattern(".", "*", *ignore, *recursive)
	default:
		f, _ = ff.FilesFromPattern(".", flag.Args()[0], *ignore, *recursive)
	}

	format := "%-50s | %s\n"

	fmt.Printf(format, "Path", "Filename")
	fmt.Println(strings.Repeat("-", 70))

	for _, x := range f {
		fmt.Printf(format, x.Path, x.Filename)
	}
}
