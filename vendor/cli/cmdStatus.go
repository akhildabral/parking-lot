package cli

import (
	"fmt"
	"parking"
)

type cmdStatus struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdStatus() ICommand {
	c := &cmdStatus{}
	c.Init()
	return c
}

func (cmd *cmdStatus) Init() {
	cmd.noOfArgs = 0
	cmd.name = "status"
}

func (cmd *cmdStatus) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	parking.GetInstance().ParkingStatus()

}

func (cmd *cmdStatus) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdStatus) Help() {
	fmt.Println(`Help --> 
status :	=>   (Check Parking Status)`)
}
