package main

import (
	"fmt"
	"math"
)

func main() {
	var array [11]float64
	var x float64

	//ask the user for 11 numbers
	fmt.Print("Enter 11 numbers: ")

	for i := 0; i < 11; i++ {
		n, _ := fmt.Scan(&x)
		//checks is n is a number
		if !math.IsNaN(float64(n)) {
			//adds the float number
			array[i] = x
		}
	}
	//reverse the array
	var reverseArray [11]float64

	for index, item := range array {
		reverseArray[10-index] = item
	}

	for _, item := range reverseArray {
		operatedNum := performOperation(item)
		if operatedNum > 100 {
			fmt.Println("Too Large")
		} else {
			fmt.Println(item)
		}
	}

}

func performOperation(x float64) float64 {
	result := math.Abs(x) * 2
	return result
}
