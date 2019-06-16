package vehicle

//Car - Car Struct, Implements IVehicle
type Car struct {
	registrationNumber string
	color              string
}

//GetColor - Get Color Of Vehicle
// Returns - Color
func (v *Car) GetColor() string {
	return v.color
}

//SetColor - Sets color of Vehicle
//color : color
func (v *Car) SetColor(color string) {
	v.color = color
}

//GetRegistrationNumber - Get Vehicle Registration Number
//Returns - Registration Number
func (v *Car) GetRegistrationNumber() string {
	return v.registrationNumber
}

//SetRegistrationNumber - Sets GRegistration Number Of Vehicle
// rn : Registration Number
func (v *Car) SetRegistrationNumber(rn string) {
	v.registrationNumber = rn
}

//Create - Creates New Car Object
// number : registration number
// color : color of car
func Create(number string, color string) *Car {
	return &Car{registrationNumber: number, color: color}
}
