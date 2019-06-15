package parking

import (
	"fmt"
	"slot"
	"vehicle"
)

type Parking struct {
	capacity int
	slots    []*slot.Slot
}

func CreateParking(cap int) *Parking {
	p := &Parking{capacity: cap, slots: nil}
	p.init()
	return p
}

func (p *Parking) init() {

	p.slots = make([]*slot.Slot, uint64(p.capacity))

	for i := range p.slots {
		p.slots[i] = slot.CreateSlot(i)
	}
}

func (p *Parking) AddVehicle(number string, color string) {
	s := p.FindEmptySlot()

	if s == nil {
		fmt.Println("All Slots Are Full")
		return
	}

	//Create a New Vehicle
	v := vehicle.Create(number, color)

	//Assign vehicle to the slot
	s.AssignVehicle(v)
}

func (p *Parking) RemoveVehicle(number string, color string) {
	s := p.FindEmptySlot()

	if s == nil {
		fmt.Println("All Slots Are Full")
		return
	}

	//Create a New Vehicle
	v := vehicle.Create(number, color)

	//Assign vehicle to the slot
	s.AssignVehicle(v)
}

func (p *Parking) RemoveVehicleAtSlot(slotNumber int) {
	if slotNumber > p.capacity {
		fmt.Println("Invalid Slot Number")
	}

	p.slots[slotNumber].RemoveVehicle()
}

func (p *Parking) ParkingStatus() {
	fmt.Println("Slot No.         Registration           Color")
	for _, s := range p.slots {
		if !s.IsEmpty() {
			fmt.Println(s.GetSlotNumber(), "       ", s.GetRegistrationNumber(), "      ", s.GetColor())
		}
	}
}

func (p *Parking) FindEmptySlot() *slot.Slot {
	for _, slot := range p.slots {
		if slot.IsEmpty() == true {
			return slot
		}
	}
	return nil
}
