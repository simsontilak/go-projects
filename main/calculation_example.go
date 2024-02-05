package main

import (
	"fmt"
	"math"
)

func main() {
	var celsius float64
	var fahrenheit float64
	var kelvin float64

	celsius = -40.00
	fahrenheit = (celsius * 9 / 5) + 32
	kelvin = celsius + 273.15

	fmt.Println("Celsius =", roundFloat(celsius, 2),
		"Fahrenheit =", roundFloat(fahrenheit, 2),
		"Kelvin =", roundFloat(kelvin, 2))

	pv := 1000.0
	r := 0.05
	n := 15

	fv := pv * math.Pow(1+r, float64(n))

	fmt.Println("PV = ", pv, "FV = ", roundFloat(fv, 2))

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
