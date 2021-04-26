package main

import "fmt"

func Sqrt(x float64) float64 {
	z := float64(1)
	for index := 0; index < 10; index++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return float64(z)
}

func main() {
	Sqrt(2)
}
