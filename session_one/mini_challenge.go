package sessionone

import (
	"fmt"
	"math"
)

func square(number int) bool  {
	if number <= 0 {
		return false
	}

	findSquare := int(math.Sqrt(float64(number)))
	return findSquare * findSquare == number
}

func cube(number int) bool  {
	if number <= 0 {
		return false
	}

	findCube := int(math.Round(math.Pow(float64(number), 1.0 / 3.0)))
  return findCube*findCube*findCube == number
}

func MiniChallengeOne()  {
	var userInput int

	fmt.Print("Input the number:")
	fmt.Scan(&userInput)


	for i := 1; i <= userInput; i++ {
		if square(i) && cube(i) {
			fmt.Println("SquareCube")
		}else if square(i) {
			fmt.Println("Square")
		}else if cube(i) {
			fmt.Println("Cube")
		}else{
			fmt.Println(i)
		}
	}
}