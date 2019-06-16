package fileinput

import (
	"bufio"
	"cli"
	"os"
)

//FileInput - File Input Struct
type FileInput struct {
	fileName string //Name of file
}

//CreateFileInput - Create Instance of File Input
// filename : name of file
// Returns - Point of FileInput Object
func CreateFileInput(fileName string) *FileInput {
	fi := &FileInput{}
	fi.Init(fileName)
	return fi
}

//Init - Initialize
// fileName : name of file
func (fi *FileInput) Init(fileName string) {
	fi.fileName = fileName
}

//Execute - Executes the File Input Reader
func (fi *FileInput) Execute() {
	//Open File
	cmdFile, err := os.Open(fi.fileName)
	if err != nil {
		return
	}
	defer cmdFile.Close()

	//Create File Scanner from buffer
	cmdScanner := bufio.NewScanner(cmdFile)
	//Create New Command Handler
	cmdHandler := cli.CreateCmdHandler()
	var cmdString string
	for cmdScanner.Scan() {
		cmdString = cmdScanner.Text()
		cmdHandler.Run(cmdString)
	}

	if err := cmdScanner.Err(); err != nil {
		return
	}
	return
}
