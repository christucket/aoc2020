package main

import (
	"os"
)

var debug bool

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		debug = true
	}

	day22()
}
