package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {

	circleSquare := flag.Float64("square", 0, "circle square")
	flag.Parse()

	if *circleSquare == 0 {
		fmt.Print("Insert circle square: ")
		fmt.Fscan(os.Stdin, circleSquare)
	}

	diameter, circleLength := diameterAndLengthOfCircleBySquare(*circleSquare)

	fmt.Println("Diameter is", diameter)
	fmt.Println("Circle length is", circleLength)

}

func diameterAndLengthOfCircleBySquare(square float64) (diameter, circleLength float64) {
	circleLength = math.Sqrt(square * 4 * math.Pi)
	diameter = circleLength / math.Pi

	return diameter, circleLength
}
