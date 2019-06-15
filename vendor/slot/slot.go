package slot

import (
	"vehicle"
)

type Slot struct {
	number  int
	vehicle vehicle.IVehicle
	isEmpty bool
}

func CreateSlot(n int) *Slot {
	return &Slot{number: n, vehicle: nil, isEmpty: true}
}

func (s *Slot) AssignVehicle(v vehicle.IVehicle) {
	s.vehicle = v
	s.isEmpty = false
}

func (s *Slot) RemoveVehicle() {
	s.vehicle = nil
	s.isEmpty = true
}

func (s *Slot) GetVehicle() vehicle.IVehicle {
	return s.vehicle
}

func (s *Slot) GetSlotNumber() int {
	return s.number
}

func (s *Slot) GetRegistrationNumber() string {
	return s.vehicle.GetRegistrationNumber()
}

func (s *Slot) GetColor() string {
	return s.vehicle.GetColor()
}

func (s *Slot) AssignSlotNumber(n int) {
	s.number = n
}

func (s *Slot) IsEmpty() bool {
	return s.isEmpty
}
