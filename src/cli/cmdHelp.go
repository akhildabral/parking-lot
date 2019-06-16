package cli

import (
	"fmt"
)

type cmdHelp struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdHelp() ICommand {
	c := &cmdHelp{}
	c.Init()
	return c
}

func (cmd *cmdHelp) Init() {
	cmd.noOfArgs = 0
	cmd.name = "help"
}

func (cmd *cmdHelp) Run(args []string) {
	cmd.Help()
}

func (cmd *cmdHelp) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdHelp) ProcessInput() {

}

func (cmd *cmdHelp) Help() {

	fmt.Println(`Help --> 
Parking Lot
-------------
Use following commands :
Command					Params       Description
------------------------------------------------
create_parking_lot :	<number of slots>  => (Create a parking lot of specified size)
park :		<registration no> <color> =>  (Park Car with Registration number and Color)
leave  :	<slot number>  =>   (Remove Car from Parking lot at specific Slot number)
status :	=>   (Check Parking Status)
registration_numbers_for_cars_with_colour :  <color>  =>  (Get Registration Number of all cars with specific Color)
slot_numbers_for_cars_with_colour : <color>  =>  (Get Slot Number of all the cars with specific Color)
slot_number_for_registration_number :  <registration number>  =>  (GEt Slot number of a specific Car by Registration Number)
exit :  =>  (Exit Program)
	`)
}
