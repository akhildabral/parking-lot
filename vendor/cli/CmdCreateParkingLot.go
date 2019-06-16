package cli

import (
	"fmt"
	"parking"
	"strconv"
)

//cmdCreateParkingLot - CreateParkingLot command object, implementing ICommand interface
type cmdCreateParkingLot struct {
	ICommand        // ICommand Interface
	noOfArgs int    //Number of arguments required for the command
	name     string // Name of the command
}

//CmdCreateParkingLot - Create Instance of cmdCreateParkingLot Struct
func CmdCreateParkingLot() ICommand {
	c := &cmdCreateParkingLot{}
	c.Init()
	return c
}

//Init - Initialize command with No Of Arguments and Command Name
func (cmd *cmdCreateParkingLot) Init() {
	cmd.noOfArgs = 1
	cmd.name = "create_parking_lot"
}

//Run - Executes Command
// args : Arguments for the command
func (cmd *cmdCreateParkingLot) Run(args []string) {
	//Check is arguments are valid
	valid := cmd.Validate(args)

	//If not, the show help to the user and retur
	if !valid {
		cmd.Help()
		return
	}

	//Do What Command is intented to Do...
	cap, _ := strconv.Atoi(args[0])
	parking.CreateParking(cap)

}

//Validate - Validated the Arguments for the command
//args : list of arguments
func (cmd *cmdCreateParkingLot) Validate(args []string) bool {
	//Check if arguments length matched the required length
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

//Help - Show help to the User
func (cmd *cmdCreateParkingLot) Help() {
	fmt.Println(`Help --> 
create_parking_lot :	<number of slots>  => (Create a parking lot of specified size)`)
}
