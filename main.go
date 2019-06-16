package main

import (
	"cli"
	"fileinput"
	"os"
)

func main() {
	//Get command line arguments
	argsWithoutProg := os.Args[1:]

	//Check if a file name is passed
	if len(argsWithoutProg) == 0 {
		//Call Intractive shell here
		cli.CreateCli().Execute()
	} else {
		//Read input from file://
		fileinput.CreateFileInput(os.Args[1]).Execute()

	}
}
