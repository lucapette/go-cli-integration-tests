package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var shout = flag.Bool("shout", false, "says it out loud")

	flag.Parse()

	f := func(input string) string { return input }

	if *shout {
		f = strings.ToUpper
	}

	for _, arg := range flag.Args() {
		fmt.Println(f(arg))
	}
}
