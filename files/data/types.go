package data

import "fmt"

// import "fmt"

type integer = int
type json = map[string]string
type distance float64 // miles more semantics to the type
type distanceKm float64

type Location string

func (loc Location) DistanceTo(destination Location) distance {
	// TO DO
	distance := distance(10.1)
	fmt.Printf(" Origin %v, Destination %v", loc, destination)
 return distance
}

func locationTest() {
	nyc := Location("32432, 94837")
	nyc.DistanceTo(Location("-23, -2348"))
	fmt.Println(nyc)
}

// METHODS
// adding methods to types
// we are injecting methods into the type distance
func (miles distance) ToKm() distanceKm { // (miles distance) method takes in distance, so we can do d.ToKm()
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distance { // (miles distance) method
	return distance(km / 1.60934)
}
func test() {
	d := distance(4.5)
	km := d.ToKm()
	miles := km.ToMiles()
	fmt.Println(d, km, miles)
}

var x integer
