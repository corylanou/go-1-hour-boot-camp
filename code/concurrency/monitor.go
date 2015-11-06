package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	// flag.Args contains all non-flag arguments
	fmt.Println(flag.Args())
}
