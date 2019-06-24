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

	var f []string

	switch flag.NArg() {
	case 0:
		f, _ = ff.FilesFromPattern("*")
	default:
		f, _ = ff.FilesFromPattern(flag.Args()[0])
	}

	format := "%-20s\n"

	fmt.Printf(format, "File or dir")
	fmt.Println(strings.Repeat("-", 70))

	for _, x := range f {
		fmt.Printf(format, x)
	}
}
