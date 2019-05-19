package main

import (
	"gojek_takehome/lib"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		// take commands from file
		fileName := argsWithoutProg[0]
		lib.ReadAndProcessFromFile(fileName)
	} else {
		//We need to make it interactive session now
		lib.ReadAndProcessStdIn()
	}

}
