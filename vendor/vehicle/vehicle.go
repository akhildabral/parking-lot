package vehicle

//IVehicle - Interface for Vehicle
type IVehicle interface {
	GetColor() string              //Get Color of Vehicle
	SetColor(string)               // Set Color of Vehicle
	GetRegistrationNumber() string // Get Registration Number Of Vehicle
	SetRegistrationNumber(string)  // Set Registration Number Of Vehicle
}

//Vehicle - Vehicle Struct, Implements IVehicle Interface
type Vehicle struct {
	registrationNumber string //Registration Number
	color              string //Color
}

//GetColor - Get Color Of Vehicle
// Returns - Color
func (v *Vehicle) GetColor() string {
	return v.color
}

//SetColor - Sets color of Vehicle
//color : color
func (v *Vehicle) SetColor(color string) {
	v.color = color
}

//GetRegistrationNumber - Get Vehicle Registration Number
//Returns - Registration Number
func (v *Vehicle) GetRegistrationNumber() string {
	return v.registrationNumber
}

//SetRegistrationNumber - Sets GRegistration Number Of Vehicle
// rn : Registration Number
func (v *Vehicle) SetRegistrationNumber(rn string) {
	v.registrationNumber = rn
}
