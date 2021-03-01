package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	aSizePtr := flag.Float64("a_size", 0, "size of A-side")
	bSizePtr := flag.Float64("b_size", 0, "size of B-side")

	flag.Parse()

	if *aSizePtr > 0 {
		fmt.Println("A size:", *aSizePtr)
	}
	if *bSizePtr > 0 {
		fmt.Println("B size:", *bSizePtr)
	}
	if *aSizePtr == 0 {
		fmt.Print("Insert A-side size: ")
		fmt.Fscan(os.Stdin, aSizePtr)
	}
	if *bSizePtr == 0 {
		fmt.Print("Insert B-side size: ")
		fmt.Fscan(os.Stdin, bSizePtr)
	}

	fmt.Println(calculateSquare(*aSizePtr, *bSizePtr))

}

func calculateSquare(a float64, b float64) float64 {
	return a * b
}
