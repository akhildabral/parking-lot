package cli

import (
	"fmt"
	"parking"
	"strconv"
)

type cmdLeave struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdLeave() ICommand {
	c := &cmdLeave{}
	c.Init()
	return c
}

func (cmd *cmdLeave) Init() {
	cmd.noOfArgs = 1
	cmd.name = "leave"
}

func (cmd *cmdLeave) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	slot, _ := strconv.Atoi(args[0])
	parking.GetInstance().RemoveVehicleAtSlot(slot)

}

func (cmd *cmdLeave) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdLeave) Help() {
	fmt.Println(`Help --> 
leave  :	<slot number>  =>   (Remove Car from Parking lot at specific Slot number)`)
}
