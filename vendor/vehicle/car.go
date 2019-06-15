package vehicle

import "fmt"

type Car struct {
	registrationNumber string
	color              string
}

func (v *Car) GetColor() string {
	return v.color
}

func (v *Car) SetColor(color string) {
	v.color = color
}

func (v *Car) GetRegistrationNumber() string {
	return v.registrationNumber
}

func (v *Car) SetRegistrationNumber(rn string) {
	v.registrationNumber = rn
}

func Create(number string, color string) *Car {
	return &Car{registrationNumber: number, color: color}
}

func Do(v IVehicle) {
	v.SetColor("Red")
	v.SetRegistrationNumber("525252")

	fmt.Println(v.GetColor())
	fmt.Println(v.GetRegistrationNumber())
}
