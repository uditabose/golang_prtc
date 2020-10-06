package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {

	var temp Fahrenheit
	var tt float32
	tt = (float32(t) * 1.8) + float32(32)
	temp = Fahrenheit(tt)
	return temp

}

func main() {
	fmt.Println(toFahrenheit(10.3))
	fmt.Println(toFahrenheit(100))
	fmt.Println(toFahrenheit(32))
}
