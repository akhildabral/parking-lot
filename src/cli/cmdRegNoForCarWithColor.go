package cli

import (
	"fmt"
	"parking"
)

type cmdRegNoForCarWithColor struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdRegNoForCarWithColor() ICommand {
	c := &cmdRegNoForCarWithColor{}
	c.Init()
	return c
}

func (cmd *cmdRegNoForCarWithColor) Init() {
	cmd.noOfArgs = 1
	cmd.name = "registration_numbers_for_cars_with_colour"
}

func (cmd *cmdRegNoForCarWithColor) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	parking.GetInstance().GetRegNoOfCarsByColor(args[0])

}

func (cmd *cmdRegNoForCarWithColor) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdRegNoForCarWithColor) Help() {
	fmt.Println(`Help --> 
registration_numbers_for_cars_with_colour :  <color>  =>  (Get Registration Number of all cars with specific Color)`)
}
