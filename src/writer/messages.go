package writer

//Commands
const (
	ParkingNotFound          = "PARKING_NOT_FOUND"
	ParkingCreated           = "PARKING_CREATED"
	ParkingIsFull            = "PARKING_IS_FULL"
	ParkingSlotAllocated     = "PARKING_SLOT_ALLOCATED"
	ParkingInvalidSlotNumber = "PARKING_INVALID_SLOT_NUMBER"
	ParkingSlotNumberIsFree  = "PARKING_SLOT_NUMBER_IS_FREE"
	NotFound                 = "NOT_FOUND"
)

//Messages ...
var Messages = map[string]string{
	ParkingNotFound:          "Parking not found",
	ParkingCreated:           "Created a parking lot with %v slots",
	ParkingIsFull:            "Sorry, parking lot is full",
	ParkingSlotAllocated:     "Allocated slot number: %v",
	ParkingInvalidSlotNumber: "Invalid Slot Number",
	ParkingSlotNumberIsFree:  "Slot number %v is free",
	NotFound:                 "Not found",
}
