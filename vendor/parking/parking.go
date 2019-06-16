package parking

import (
	"fmt"
	"slot"
	"strconv"
	"strings"
	"vehicle"
)

//Holds ParkingLot instance for later use
var instance *Parking = nil

//Parking - Structure of Parking lot
//	capacity : Size of parking lot
// slots : Slice to hold slots instances
type Parking struct {
	capacity int
	slots    []*slot.Slot
}

//CreateParking - Creates Parking Object
// cap : Size of parking lot
func CreateParking(cap int) *Parking {
	instance = &Parking{capacity: cap, slots: nil}
	instance.init()

	return instance
}

//GetInstance - Get Singlton instance of Parking Lot
func GetInstance() *Parking {
	//Check if instance is nil
	if instance == nil {
		fmt.Println("Parking not found")
	}

	return instance
}

//init - Initialize Paeking lot
func (p *Parking) init() {
	//Initialize slots for specied size
	p.slots = make([]*slot.Slot, uint64(p.capacity))

	//Loop over slots and create instance of Slots
	for i := range p.slots {
		p.slots[i] = slot.CreateSlot(i + 1)
	}

	fmt.Println("Created a parking lot with", p.capacity, "slots")
}

//ParkVehicle - Add Vehicle to parking lot
func (p *Parking) ParkVehicle(number string, color string) {
	//Find a empty slot
	n, s := p.FindEmptySlot()

	// Check if empty slot found. If not show message
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

//RemoveVehicleAtSlot - Removes vehicle at specific slot number
// slotNumber : Slot number where vehicle need to be removed
func (p *Parking) RemoveVehicleAtSlot(slotNumber int) {
	//Check if slot number is corrent
	if slotNumber > p.capacity {
		fmt.Println("Invalid Slot Number")
		return
	}

	//Remove vehicle at the spot
	p.slots[slotNumber-1].RemoveVehicle()

	fmt.Println("Slot number", slotNumber, "is free")
}

//ParkingStatus - Prints current parking status
func (p *Parking) ParkingStatus() {
	output := []string{fmt.Sprintf("%-12s%-19s%-10s",
		"Slot No.",
		"Registration No",
		"Colour",
	)}
	for _, s := range p.slots {
		if !s.IsEmpty() {
			output = append(output, fmt.Sprintf("%-12v%-19s%-10s",
				s.GetSlotNumber(),
				s.GetRegistrationNumber(),
				s.GetColor(),
			))
		}
	}

	//Add new line at the end of each item in list
	s := strings.Join(output, "\n")
	fmt.Println(s)
}

//FindEmptySlot - Search for empty slot in parking lot
// Returns : Slot number, Slot Object
func (p *Parking) FindEmptySlot() (int, *slot.Slot) {
	//Loop through all slots to find is the is any empty slot available
	for _, slot := range p.slots {
		if slot.IsEmpty() == true {
			return slot.GetSlotNumber(), slot
		}
	}
	return 0, nil
}

//GetRegNoOfCarsByColor - Find all cars by Color, Return registration numbers
// color : color of th car
// Returns - List of all the registration numbers
func (p *Parking) GetRegNoOfCarsByColor(color string) []string {
	regNos := []string{}
	for _, slot := range p.slots {
		if (!slot.IsEmpty()) && (slot.GetColor() == color) {
			regNos = append(regNos, slot.GetRegistrationNumber())
		}
	}

	//Show message if nothing found
	if len(regNos) == 0 {
		regNos = append(regNos, "Not found")
	}

	fmt.Println(strings.Join(regNos, ", "))
	return regNos
}

//GetSlotNosOfCarsByColor - Get Slot numbers of cars with same color
// color : color of the car
// Returns - List of all the slot numbers
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

//GetSlotNosOfCarsByRegNo - Find slot number by Registration number
// regNo : registration nunmber of the car
// Returns - List of cars
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
