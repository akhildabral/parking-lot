package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Cli - Command Line Interface
type Cli struct {
	cmdHandler *commandHandler //Hold pointer to CommandHandlere
}

//CreateCli - Instanciate CLI
func CreateCli() *Cli {
	c := &Cli{}
	c.Init()
	return c
}

//Init - Initialize
func (c *Cli) Init() {
	//Create Command Handler object and assign to Cli object
	c.cmdHandler = CreateCmdHandler()
}

//ShowPrompt - Display Prompt
func (c *Cli) ShowPrompt() {
	fmt.Print("$ ")
}

//Execute - Execute CLI Process
func (c *Cli) Execute() {
	//Read Input Streame
	reader := bufio.NewReader(os.Stdin)
	c.ShowPrompt()
	//Keep reading input from Stdin untill 'exit'
	for {
		//Read input till new line(enter) is presses
		cmdInput, _ := reader.ReadString('\n')
		//Trim new Line character from the end
		cmdInput = strings.TrimRight(cmdInput, "\n")

		//Check if exit is types, Break the loop and close process
		if cmdInput == "exit" {
			return
		}

		//Is Cmd is not empty, Process Command, Else show help
		if "" != cmdInput {
			//Pass command to the handler
			out, err := c.cmdHandler.Run(cmdInput)
			c.cmdHandler.processOutput(out, err)
		} else {
			//Show Help Message
			c.cmdHandler.commands["help"].Run(nil)
		}
		c.ShowPrompt()
	}
}
