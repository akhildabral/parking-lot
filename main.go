package main

import (
	"cli"
	"fileinput"
	"os"
)

//main entry point of the program
func main() {
	//Get command line arguments
	argsWithoutProgName := os.Args[1:]

	//Check if a file name is passed
	if len(argsWithoutProgName) == 0 {
		//Call CLI here...
		cli.CreateCli().Execute()
	} else {
		//Read input from file://...
		fileinput.CreateFileInput(os.Args[1]).Execute()
	}
}
