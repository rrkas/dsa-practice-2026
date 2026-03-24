package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'matrixRotation' function below.
 *
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY matrix
 *  2. INTEGER r
 */

func checkAllSame(matrix [][]int32, m, n int) bool {
	e := matrix[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if e != matrix[i][j] {
				return false
			}
		}
	}

	return true
}

func rotateLayer(matrix [][]int32, turns int32, t, b, l, r int) [][]int32 {
	peri := 2 * ((b - t) + (r - l))
	turns = turns % int32(peri)

	if turns == 0 {
		return matrix
	}

	arr := make([]int32, peri)
	ai := 0
	// top row
	for j := l; j < r; j++ {
		arr[ai] = matrix[t][j]
		ai++
	}

	// right col
	for i := t; i < b; i++ {
		arr[ai] = matrix[i][r]
		ai++
	}

	// bottom row
	for j := r; j > l; j-- {
		arr[ai] = matrix[b][j]
		ai++
	}

	// left col
	for i := b; i > t; i-- {
		arr[ai] = matrix[i][l]
		ai++
	}

	arr = append(arr[turns:], arr[:turns]...)

	ai = 0
	// top row
	for j := l; j < r; j++ {
		matrix[t][j] = arr[ai]
		ai++
	}

	// right col
	for i := t; i < b; i++ {
		matrix[i][r] = arr[ai]
		ai++
	}

	// bottom row
	for j := r; j > l; j-- {
		matrix[b][j] = arr[ai]
		ai++
	}

	// left col
	for i := b; i > t; i-- {
		matrix[i][l] = arr[ai]
		ai++
	}

	return matrix
}

func printMatrix(matrix [][]int32, m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}

func matrixRotation(matrix [][]int32, turns int32) {
	m, n := len(matrix), len(matrix[0])

	if checkAllSame(matrix, m, n) {
		printMatrix(matrix, m, n)
		return
	}

	t, b, l, r := 0, m-1, 0, n-1
	for t < b && l < r {
		rotateLayer(matrix, turns, t, b, l, r)
		t++
		b--
		l++
		r--
	}

	printMatrix(matrix, m, n)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	rTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	r := int32(rTemp)

	var matrix [][]int32
	for i := 0; i < int(m); i++ {
		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(n) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	matrixRotation(matrix, r)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
