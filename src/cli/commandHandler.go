package cli

import (
	"strings"
)

//ICommand - Command Interface
type ICommand interface {
	Init()                  // Initialize
	Run([]string)           // Run Command
	Validate([]string) bool //Validate Command
	Help()                  // SHow Help
}

type commands map[string]ICommand

//commandHandle - Command Handler Struct
type commandHandler struct {
	commands commands
}

//CreateCmdHandler - Instanciate CommandHandler
func CreateCmdHandler() *commandHandler {
	cmh := &commandHandler{make(commands)}
	//Register all the commands defined in commands file
	cmh.RegisterCommand()
	return cmh
}

//RegisteeCommand - Register Command defined in CmdRegistry map
func (c *commandHandler) RegisterCommand() {
	//Loop through all the commands and instanciate
	for cmd, h := range CmdRegistry {
		c.commands[cmd] = h()
	}
}

//processInput - Process Input Command
func (c *commandHandler) processInput(cmd string) (string, []string) {
	//Split command attributes
	s := strings.Split(cmd, " ")
	//Return Command string and arguments list
	return s[0], s[1:]
}

//processOutput - Process Output Data
func (c *commandHandler) processOutput(out string, err error) {

}

func (c *commandHandler) processCommand() {
	//c.commands["create_parking_lot"].Run()
}

//Run - Exucute Command
func (c *commandHandler) Run(cmd string) (string, error) {
	command, args := c.processInput(cmd)
	if c.commands[command] != nil {
		c.commands[command].Run(args)
	} else {
		c.commands[command].Help()
	}

	return "", nil
}
