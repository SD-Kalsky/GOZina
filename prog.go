package main

import "fmt"

func main() {
	var ar1 [3]float32 = [3]float32{0.25, 0.5, 0.25}
	var ar2 [3]float32 = [3]float32{4, 2, 4}
	var drv = DRV{
		size:          3,
		probabilities: ar1,
		values:        ar2,
	}
	fmt.Println(ExpectedValue((drv)))

}

// type Fraction struct {
// 	numinator   int
// 	denominator int
// }

type DRV struct {
	size int
	// probabilities []Fraction
	probabilities []float32
	values        []float32
}

func ExpectedValue(drv DRV) float32 {
	var ev float32 = 0
	for i := 0; i < drv.size; i++ {
		ev = ev + (drv.probabilities[i] * drv.values[i])
	}
	return ev
}
