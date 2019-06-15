package vehicle

type IVehicle interface {
	GetColor() string
	SetColor(string)
	GetRegistrationNumber() string
	SetRegistrationNumber(string)
}

type Vehicle struct {
	registrationNumber string
	color              string
}

func (v *Vehicle) GetColor() string {
	return v.color
}

func (v *Vehicle) SetColor(color string) {
	v.color = color
}

func (v *Vehicle) GetRegistrationNumber() string {
	return v.registrationNumber
}

func (v *Vehicle) SetRegistrationNumber(rn string) {
	v.registrationNumber = rn
}
