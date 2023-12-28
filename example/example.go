package main

import (
	"fmt"
	"os"

	version "github.com/eric-tech01/simple-version"
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Println(version.BuildTime)
		fmt.Printf("%v", version.GetVersion())
	}

}
