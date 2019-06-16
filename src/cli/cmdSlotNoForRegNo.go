package cli

import (
	"fmt"
	"parking"
)

type cmdSlotNoForRegNo struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdSlotNoForRegNo() ICommand {
	c := &cmdSlotNoForRegNo{}
	c.Init()
	return c
}

func (cmd *cmdSlotNoForRegNo) Init() {
	cmd.noOfArgs = 1
	cmd.name = "slot_number_for_registration_number"
}

func (cmd *cmdSlotNoForRegNo) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	parking.GetInstance().GetSlotNosOfCarsByRegNo(args[0])

}

func (cmd *cmdSlotNoForRegNo) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdSlotNoForRegNo) Help() {
	fmt.Println(`Help --> 
slot_number_for_registration_number :  <registration number>  =>  (GEt Slot number of a specific Car by Registration Number)`)
}
