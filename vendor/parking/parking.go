package parking

import (
	"fmt"
	"slot"
	"strconv"
	"strings"
	"vehicle"
)

var instance *Parking = nil

type Parking struct {
	capacity int
	slots    []*slot.Slot
}

func CreateParking(cap int) *Parking {
	instance = &Parking{capacity: cap, slots: nil}
	instance.init()

	return instance
}

func GetInstance() *Parking {
	if instance == nil {
		fmt.Println("Parking not found")
	}

	return instance
}

func (p *Parking) init() {

	p.slots = make([]*slot.Slot, uint64(p.capacity))

	for i := range p.slots {
		p.slots[i] = slot.CreateSlot(i + 1)
	}

	fmt.Println("Created a parking lot with", p.capacity, "slots")
}

func (p *Parking) ParkVehicle(number string, color string) {
	n, s := p.FindEmptySlot()

	if s == nil {
		fmt.Println("Sorry, parking lot is full")
		return
	}

	//Create a New Vehicle
	v := vehicle.Create(number, color)

	//Assign vehicle to the slot
	s.AssignVehicle(v)
	fmt.Println("Allocated slot number:", n)
}

func (p *Parking) RemoveVehicleAtSlot(slotNumber int) {
	if slotNumber > p.capacity {
		fmt.Println("Invalid Slot Number")
	}

	p.slots[slotNumber-1].RemoveVehicle()

	fmt.Println("Slot number", slotNumber, "is free")
}

func (p *Parking) ParkingStatus() {
	output := []string{fmt.Sprintf("%-10s%-20s%-10s",
		"Slot No.",
		"Registration No",
		"Colour",
	)}
	for _, s := range p.slots {
		if !s.IsEmpty() {
			output = append(output, fmt.Sprintf("%-10v%-20s%-10s",
				s.GetSlotNumber(),
				s.GetRegistrationNumber(),
				s.GetColor(),
			))
		}
	}
	s := strings.Join(output, "\n")
	fmt.Println(s)
}

func (p *Parking) FindEmptySlot() (int, *slot.Slot) {
	for _, slot := range p.slots {
		if slot.IsEmpty() == true {
			return slot.GetSlotNumber(), slot
		}
	}
	return 0, nil
}

func (p *Parking) GetRegNoOfCarsByColor(color string) []string {
	regNos := []string{}
	for _, slot := range p.slots {
		if (!slot.IsEmpty()) && (slot.GetColor() == color) {
			regNos = append(regNos, slot.GetRegistrationNumber())
		}
	}

	if len(regNos) == 0 {
		regNos = append(regNos, "Not found")
	}

	fmt.Println(strings.Join(regNos, ", "))
	return regNos
}

func (p *Parking) GetSlotNosOfCarsByColor(color string) []string {
	slotNos := []string{}
	for _, slot := range p.slots {
		if (!slot.IsEmpty()) && (slot.GetColor() == color) {
			slotNos = append(slotNos, strconv.Itoa(slot.GetSlotNumber()))
		}
	}

	if len(slotNos) == 0 {
		slotNos = append(slotNos, "Not found")
	}

	fmt.Println(strings.Join(slotNos, ", "))
	return slotNos
}

func (p *Parking) GetSlotNosOfCarsByRegNo(regNo string) []string {
	slotNos := []string{}
	for _, slot := range p.slots {
		if (!slot.IsEmpty()) && (slot.GetRegistrationNumber() == regNo) {
			slotNos = append(slotNos, strconv.Itoa(slot.GetSlotNumber()))
		}
	}

	if len(slotNos) == 0 {
		slotNos = append(slotNos, "Not found")
	}

	fmt.Println(strings.Join(slotNos, ", "))
	return slotNos
}
