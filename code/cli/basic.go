package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string

	flag.StringVar(&cmd, "cmd", cmd, `cmd can be either "hello" or "bye"`)
	flag.Parse()

	switch cmd {
	case "hello":
		fmt.Println("Hello!")
	case "bye":
		fmt.Println("Bye!")
	default:
		flag.Usage()
	}
}
