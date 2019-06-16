package cli

import (
	"fmt"
	"strings"
)

type ICommand interface {
	Init()
	Run([]string)
	ProcessInput()
	Validate([]string) bool
	Help()
}

type Command struct {
}

type commands map[string]ICommand

type commandHandler struct {
	commands commands
}

func CreateCmdHandler() *commandHandler {
	cmh := &commandHandler{make(commands)}
	cmh.RegisterCommand()
	return cmh
}

// func (c *cli) readCommandsConfig() {
// 	//Get config file's absolute path
// 	configPath, _ := filepath.Abs("./vendor/cli/commands.json")

// 	//Read Commands config
// 	file, err := os.Open(configPath)
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer file.Close()

// 	byteValue, _ := ioutil.ReadAll(file)

// 	// we unmarshal our byteArray which contains our
// 	// jsonFile's content into 'users' which we defined above
// 	json.Unmarshal(byteValue, &c.commands)

// 	//fmt.Println(c.commands["create_parking_lot"].Handler.(func()))
// 	MyCmd["create_parking_lot"]()
// }

func (c *commandHandler) RegisterCommand() {
	for cmd, h := range CmdRegistry {
		c.commands[cmd] = h()
	}
}

//Init ... Exported
func (c *commandHandler) Init() {
	//Read supported commands from config file
	//	c.readCommandsConfig()
	//	c.processCommand()
}

func (c *commandHandler) processInput(cmd string) (string, []string) {
	s := strings.Split(cmd, " ")
	return s[0], s[1:]
}

func (c *commandHandler) processOutput(out string, err error) {

}

func (c *commandHandler) processCommand() {
	//c.commands["create_parking_lot"].Run()
}

func (c *commandHandler) Run(cmd string) (string, error) {
	command, args := c.processInput(cmd)
	if c.commands[command] != nil {
		c.commands[command].Run(args)
	} else {
		fmt.Println("Invalid COmmand")
	}

	return "da", nil
}
