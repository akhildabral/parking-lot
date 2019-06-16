package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type command struct {
	Handler string `json:"handler`
	Args    int    `json:"args"`
}

type commands map[string]command

type cli struct {
	commands commands
}

func CreateInstance() *cli {
	return &cli{}
}

func (c *cli) readCommandsConfig() {
	//Get config file's absolute path
	configPath, _ := filepath.Abs("./vendor/cli/commands.json")

	//Read Commands config
	file, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &c.commands)

	fmt.Println(c.commands)
}

func (c *cli) RegisterCommand() {

}

//Init ... Exported
func (c *cli) Init() {
	//Read supported commands from config file
	c.readCommandsConfig()
	c.processInput()
}

func (c *cli) processInput() {

}

func (c *cli) processCommand() {

}

func (c *cli) Run(){
	
}
