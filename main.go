package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var shout = flag.Bool("shout", false, "says it out loud")
	var whisper = flag.Bool("whisper", false, "says it quietly")

	flag.Parse()

	f := func(input string) string { return input }

	if *shout {
		f = strings.ToUpper
	}

	if *whisper {
		f = strings.ToLower
	}

	for _, arg := range flag.Args() {
		fmt.Println(f(arg))
	}
}
