package main

import (
	"fmt"
	"math"
)

var hourglasses = [6][6]int{
	{1, 1, 1, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 0},
	{0, 0, 2, 4, 4, 0},
	{0, 0, 0, 2, 0, 0},
	{0, 0, 1, 2, 4, 0},
}

func main() {
	fmt.Println(getHigherHourglass(hourglasses))
}

func getHigherHourglass(matrix [6][6]int) int {
	higher := math.MinInt32

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := sumHourglass(matrix, i, j)
			if sum > higher {
				higher = sum
			}
		}
	}

	return higher
}

func sumHourglass(matrix [6][6]int, row, col int) int {
	sum := matrix[row][col] + matrix[row][col+1] + matrix[row][col+2] +
		matrix[row+1][col+1] +
		matrix[row+2][col] + matrix[row+2][col+1] + matrix[row+2][col+2]
	return sum
}
