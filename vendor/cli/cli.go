package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cli struct {
	cmdHandler *commandHandler
}

func CreateCli() *Cli {
	c := &Cli{}
	c.Init()
	return c
}

func (c *Cli) Init() {
	c.cmdHandler = CreateCmdHandler()
}

func (c *Cli) ShowPrompt() {
	fmt.Print(">> ")
}

func (c *Cli) Execute() {
	reader := bufio.NewReader(os.Stdin)
	c.ShowPrompt()
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, "\n")
		if "" != cmdInput {
			out, err := c.cmdHandler.Run(cmdInput)
			c.cmdHandler.processOutput(out, err)
		} else {
			fmt.Println("HELPPPPP")
			c.cmdHandler.commands["help"].Run(nil)
		}
		c.ShowPrompt()
	}
}
