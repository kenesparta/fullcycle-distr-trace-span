package entity

type Temperature struct {
	fahrenheit float64
	celsius    float64
	kelvin     float64
}

func NewTemperature(celsius float64) *Temperature {
	t := &Temperature{}
	t.celsius = celsius
	t.convert()
	return t
}

func (t *Temperature) convertFahrenheit() {
	t.fahrenheit = t.celsius*1.8 + 32.0
}

func (t *Temperature) convertKelvin() {
	t.kelvin = t.celsius + 273.0
}

func (t *Temperature) Fahrenheit() float64 {
	return t.fahrenheit
}

func (t *Temperature) Kelvin() float64 {
	return t.kelvin
}

func (t *Temperature) Celsius() float64 {
	return t.celsius
}

func (t *Temperature) convert() {
	t.convertFahrenheit()
	t.convertKelvin()
}
