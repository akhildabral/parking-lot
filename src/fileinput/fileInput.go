package fileinput

import (
	"bufio"
	"cli"
	"os"
)

type FileInput struct {
	fileName string
}

func CreateFileInput(fileName string) *FileInput {
	fi := &FileInput{}
	fi.Init(fileName)
	return fi
}

func (fi *FileInput) Init(fileName string) {
	fi.fileName = fileName
}

func (fi *FileInput) Execute() {
	cmdFile, err := os.Open(fi.fileName)
	if err != nil {
		return
	}
	defer cmdFile.Close()

	cmdScanner := bufio.NewScanner(cmdFile)
	cmdMgr := cli.CreateCmdHandler()
	var cmdString string
	for cmdScanner.Scan() {
		cmdString = cmdScanner.Text()
		cmdMgr.Run(cmdString)
		//cmdMgr.processOutput(out, err)
	}

	if err := cmdScanner.Err(); err != nil {
		return
	}
	return
}
