package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	f, err := Sqrt(-2)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
