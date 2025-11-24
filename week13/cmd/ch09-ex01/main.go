package main

import "fmt"

type Kilometers float64
type Meters float64
type Miles float64

func (km Kilometers) ToMiles() Miles {
	return Miles(km / 1.609)
}

func (m Meters) ToMiles() Miles {
	return Miles(m / 1609)
}

func main() {
	km := Kilometers(160.1)
	fmt.Printf("%0.3f km/h is equal to %0.3f mph\n", km, km.ToMiles())
	m := Meters(10000)
	fmt.Printf("%0.3f m/h is equal to %0.3f mph\n", m, m.ToMiles())
}
