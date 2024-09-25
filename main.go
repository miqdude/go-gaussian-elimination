package main

import (
	"fmt"
)

var idxVar = []string{"X", "Y", "Z"}

const PRECISION = 10

// SolveLinearEquation uses Gaussian Elimination to solve Ax = b
func SolveLinearEquation(A [][]float64, b []float64) ([]float64, error) {

	// Gaussian Elimination
	rowA := len(A)
	for i := 0; i < rowA-1; i++ {
		// fmt.Printf("Step : %d\n", i)
		for j := i + 1; j < rowA; j++ {
			if A[i][i] == 0 {
				continue
			}
			scalingFactor := A[j][i] / A[i][i]
			fmt.Printf("STEP : %d\n", i+j)
			fmt.Printf("Eliminating L%d with scaling factor from element %d,%d divided by element %d,%d\n", j+1, j+1, i+1, i+1, i+1)
			fmt.Printf("L%d = L%d - (%.1f*L%d)\n\n", j+1, j+1, scalingFactor, i+1)

			// reduce the current row with the row below
			colA := rowA
			for k := 0; k < colA; k++ {
				temp := A[j][k] - scalingFactor*A[i][k]
				A[j][k] = temp
			}

			// do the same thing for matrix B
			temp_ := b[j] - scalingFactor*b[i]
			b[j] = temp_

			formattedPrint(A, b)
			fmt.Print("\n")
		}
	}

	// backward substitution
	fmt.Print("Doing backward substitution\n\n")
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
		fmt.Printf("%s = (%.1f - %.1f)/%.1f = %.1f\n\n", idxVar[i], b[i], quantifier, A[i][i], res[i])
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

func isRound(val float64) bool {
	tmp := val * PRECISION
	mod := int(tmp) % PRECISION

	return mod == 0
}

func formattedPrint(A [][]float64, b []float64) {
	for i := 0; i < len(A); i++ {
		showFormatted(A[i][:], b[i])
	}
}

func showFormatted(A []float64, b float64) {
	res := ""
	isLeading := true
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			continue
		}

		if !isLeading && A[i] > 0 {
			res += "+"
		}

		if A[i] != 1 && A[i] != -1 {
			if isRound(A[i]) {
				res += fmt.Sprintf("%.0f", A[i])
			} else {
				res += fmt.Sprintf("%.1f", A[i])
			}
		}

		if A[i] == -1 {
			res += "-"
		}

		res += idxVar[i]
		isLeading = false
	}

	if isRound(b) {
		res += fmt.Sprintf("=%.0f", b)
	} else {
		res += fmt.Sprintf("=%.1f", b)
	}
	fmt.Print(res, "\n")
}

func main() {
	A, b := readInput()

	fmt.Print("Solving : \n\n")
	formattedPrint(A, b)
	fmt.Println()

	// Solve Ax = b
	_, err := SolveLinearEquation(A, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
