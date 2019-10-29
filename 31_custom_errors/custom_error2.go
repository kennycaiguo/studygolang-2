package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (a *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", a.radius, a.err)
}

func circleAera1(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleAera1(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}