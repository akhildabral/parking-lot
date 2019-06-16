package cli

import (
	"fmt"
	"parking"
)

type cmdSlotNoForCarsWithColor struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdSlotNoForCarsWithColor() ICommand {
	c := &cmdSlotNoForCarsWithColor{}
	c.Init()
	return c
}

func (cmd *cmdSlotNoForCarsWithColor) Init() {
	cmd.noOfArgs = 1
	cmd.name = "slot_numbers_for_cars_with_colour"
}

func (cmd *cmdSlotNoForCarsWithColor) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	parking.GetInstance().GetSlotNosOfCarsByColor(args[0])

}

func (cmd *cmdSlotNoForCarsWithColor) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdSlotNoForCarsWithColor) Help() {
	fmt.Println(`Help --> 
slot_numbers_for_cars_with_colour : <color>  =>  (Get Slot Number of all the cars with specific Color)`)
}
