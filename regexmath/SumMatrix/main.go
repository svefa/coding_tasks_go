package main

import "fmt"

type NumMatrix struct {
	n int
	m int
	S [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{n: len(matrix) + 1, m: len(matrix[0]) + 1, S: make([][]int, len(matrix)+1)}

	numMatrix.S[0] = make([]int, len(matrix[0])+1)

	for i := 0; i < numMatrix.n-1; i++ {
		numMatrix.S[i+1] = make([]int, len(matrix[0])+1)

		t := 0
		for j := 0; j < numMatrix.m-1; j++ {
			t += matrix[i][j]
			numMatrix.S[i+1][j+1] = numMatrix.S[i][j+1] + t
		}
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	fmt.Println()
	return this.S[row2+1][col2+1] + this.S[row1][col1] - this.S[row1][col2+1] - this.S[row2+1][col1]
}

func main() {
	row1, col1, row2, col2 := 1, 1, 2, 2
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	obj := Constructor(matrix)
	fmt.Println(obj)
	s := obj.SumRegion(row1, col1, row2, col2)
	fmt.Println(s)

}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
