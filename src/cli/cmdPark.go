package cli

import (
	"fmt"
	"parking"
)

type cmdPark struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdPark() ICommand {
	c := &cmdPark{}
	c.Init()
	return c
}

func (cmd *cmdPark) Init() {
	cmd.noOfArgs = 2
	cmd.name = "park"
}

func (cmd *cmdPark) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	parking.GetInstance().ParkVehicle(args[0], args[1])

}

func (cmd *cmdPark) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}


func (cmd *cmdPark) Help() {
	fmt.Println(`Help --> 
park :		<registration no> <color> =>  (Park Car with Registration number and Color)`)
}
