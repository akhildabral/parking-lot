package cli

import (
	"fmt"
	"parking"
	"strconv"
)

type cmdCreateParkingLot struct {
	ICommand
	noOfArgs int
	name     string
}

func CmdCreateParkingLot() ICommand {
	c := &cmdCreateParkingLot{}
	c.Init()
	return c
}

func (cmd *cmdCreateParkingLot) Init() {
	cmd.noOfArgs = 1
	cmd.name = "create_parking_lot"
}

func (cmd *cmdCreateParkingLot) Run(args []string) {
	valid := cmd.Validate(args)

	if !valid {
		cmd.Help()
		return
	}

	cap, _ := strconv.Atoi(args[0])
	parking.CreateParking(cap)

}

func (cmd *cmdCreateParkingLot) Validate(args []string) bool {
	if len(args) == cmd.noOfArgs {
		return true
	}
	return false
}

func (cmd *cmdCreateParkingLot) ProcessInput() {

}

func (cmd *cmdCreateParkingLot) Help() {
	fmt.Println("You Need SOme help Creating Paking!")
}
