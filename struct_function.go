package main

import "fmt"

const value float64 = 65534
const value2 float64 = 1.60934

type car struct {
	gasPedal      uint16
	breakPedal    uint16
	steeringWheel uint16
	topSpeedKm    float64
}

// Value receivers
func (c car) kmh() float64 {
	return float64(c.gasPedal) * float64(c.topSpeedKm)
}

func main() {
	fmt.Println("GO is working!")

	aCar := &car{
		gasPedal:      22341,
		breakPedal:    0,
		steeringWheel: 12561,
		topSpeedKm:    225.0,
	}
	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.breakPedal)
	fmt.Println(aCar.steeringWheel)
	fmt.Println(aCar.topSpeedKm)
	fmt.Println("Function")
	fmt.Println(aCar.kmh())
}
