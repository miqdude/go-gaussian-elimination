package main

import (
	"fmt"
)

var idxVar = []string{"X", "Y", "Z"}

// SolveLinearEquation uses Gaussian Elimination to solve Ax = b
func SolveLinearEquation(A [][]float64, b []float64) ([]float64, error) {

	// Gaussian Elimination
	rowA := len(A)
	for i := 0; i < rowA-1; i++ {
		for j := i + 1; j < rowA; j++ {
			if A[i][i] == 0 {
				continue
			}
			scalingFactor := A[j][i] / A[i][i]

			// reduce the current row with the row below
			colA := rowA
			for k := 0; k < colA; k++ {
				temp := A[j][k] - scalingFactor*A[i][k]
				A[j][k] = temp
			}

			// do the same thing for matrix B
			temp_ := b[j] - scalingFactor*b[i]
			b[j] = temp_
		}

		// display per step
		fmt.Printf("Step : %d\n", i)
		showFormatted(A, b)
	}

	// backward substitution
	res := make([]float64, rowA)
	for i := rowA - 1; i >= 0; i-- {
		if A[i][i] == 0 {
			continue
		}

		quantifier := 0.0
		for j := i + 1; j < rowA; j++ {
			quantifier += A[i][j] * res[j]
		}

		res[i] = (b[i] - quantifier) / A[i][i]
	}

	return res, nil
}

func readInput() ([][]float64, []float64) {
	var length int
	fmt.Scanf("%d\n", &length)

	input := make([][]float64, length)
	for i := 0; i < length; i++ {
		input[i] = make([]float64, length)
		for j := 0; j < length; j++ {
			var tmp float64
			fmt.Scan(&tmp)
			input[i][j] = tmp
		}
	}

	res := make([]float64, length)
	for i := 0; i < length; i++ {
		var tmp float64
		fmt.Scan(&tmp)
		res[i] = tmp
	}

	return input, res

}

func showFormatted(A [][]float64, b []float64) {
	row := len(A)
	for i := 0; i < row; i++ {
		isLeading := true
		for j := 0; j < row; j++ {
			el := A[i][j]
			if el == 0 {
				continue
			} else {
				isLeading = false
			}

			if !isLeading {
				if el > 1 {
					fmt.Printf("+%.0f", el)
				} else if el < 1 {
					fmt.Printf("%.0f", el)
				}
			}

			fmt.Print(idxVar[j], " ")
		}

		fmt.Printf("= %.0f\n", b[i])
	}
}

func printSolution(b []float64) {
	fmt.Print("Solution: ")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%s= %.0f ", idxVar[i], b[i])
	}
	fmt.Println()
}

func main() {
	A, b := readInput()

	// Solve Ax = b
	solution, err := SolveLinearEquation(A, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printSolution(solution)
}
