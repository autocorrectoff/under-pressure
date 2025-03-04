package main

import (
	"fmt"
)

// Pressure calculates the pressure at position (row, col)
func Pressure(row, col int) float64 {
	if row == 0 {
		return 0 // The top brick has no pressure
	}

	if col < 0 || col > row {
		return 0 // Out of bounds
	}

	leftParent := 0.0
	if col > 0 {
		leftParent = (1 + Pressure(row-1, col-1)) / 2
	}

	rightParent := 0.0
	if col < row {
		rightParent = (1 + Pressure(row-1, col)) / 2
	}

	return leftParent + rightParent
}

func main() {
	// rows := 10
	// for r := 0; r < rows; r++ {
	// 	for c := 0; c <= r; c++ {
	// 		fmt.Printf("pressure(%d, %d) = %.6f\n", r, c, Pressure(r, c))
	// 	}
	// }
	fmt.Printf("pressure(%d, %d) = %.6f\n", 322, 156, Pressure(322, 156))
}
