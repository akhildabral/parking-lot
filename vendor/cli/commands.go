package cli

type cmd map[string](func() ICommand)

var CmdRegistry cmd = cmd{
	"create_parking_lot": CmdCreateParkingLot,
	"leave":              CmdLeave,
	"status":             CmdStatus,
	"park":               CmdPark,
	"registration_numbers_for_cars_with_colour": CmdRegNoForCarWithColor,
	"slot_numbers_for_cars_with_colour":         CmdSlotNoForCarsWithColor,
	"slot_number_for_registration_number":       CmdSlotNoForRegNo,
}
