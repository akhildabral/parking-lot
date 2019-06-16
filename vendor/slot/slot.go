package slot

import (
	"vehicle"
)

//Slot - Slot struct
type Slot struct {
	number  int              //Slot Number
	vehicle vehicle.IVehicle // Pointer to Vehicle interface
	isEmpty bool             //Is slot empty
}

//CreateSlot - Create new Slot
// n : Slot number
// Returns - Pointer to New SLot
func CreateSlot(n int) *Slot {
	return &Slot{number: n, vehicle: nil, isEmpty: true}
}

//AssignVehicle - Assign Vehicle to he slot
// v - vehicle instance
func (s *Slot) AssignVehicle(v vehicle.IVehicle) {
	s.vehicle = v
	s.isEmpty = false
}

//RemoveVehicle - Remove Vehicle from the slot
func (s *Slot) RemoveVehicle() {
	s.vehicle = nil
	s.isEmpty = true
}

//GetVehicle - Get Vehicle at the spot
// Returns - Vehicle Instance
func (s *Slot) GetVehicle() vehicle.IVehicle {
	return s.vehicle
}

//GetSlotNumber - Get Current Slot number
//Return - Slot number
func (s *Slot) GetSlotNumber() int {
	return s.number
}

//GetRegistrationNumber - Get Registration number of vehicle at the spot
// Return - Registration Number
func (s *Slot) GetRegistrationNumber() string {
	return s.vehicle.GetRegistrationNumber()
}

//GetColor - Get color of the vehicle at the spot
// Returns : color of vehicle
func (s *Slot) GetColor() string {
	return s.vehicle.GetColor()
}

//AssignSlotNumber - Assign slot number to the slot
// n : slot number
func (s *Slot) AssignSlotNumber(n int) {
	s.number = n
}

//IsEmpty : Check is slot is empty
// Returns : Slot status
func (s *Slot) IsEmpty() bool {
	return s.isEmpty
}
