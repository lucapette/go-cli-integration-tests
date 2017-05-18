package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		shout   = flag.Bool("shout", false, "say it out loud")
		version = flag.Bool("version", false, "shows version information")
	)

	flag.Parse()

	if *version {
		fmt.Println("echo v0.0.1")
		os.Exit(0)
	}

	f := func(input string) string { return input }

	if *shout {
		f = strings.ToUpper
	}

	for _, arg := range flag.Args() {
		fmt.Println(f(arg))
	}
}
