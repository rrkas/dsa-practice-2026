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
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func detonate(grid [][]bool, m, n int) [][]bool {
	newGrid := make([][]bool, m)
	for i := 0; i < m; i++ {
		newGrid[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] {
				continue
			}
			valid := true
			if i > 0 && grid[i-1][j] {
				valid = false
			}
			if valid && i+1 < m && grid[i+1][j] {
				valid = false
			}
			if valid && j > 0 && grid[i][j-1] {
				valid = false
			}
			if valid && j+1 < n && grid[i][j+1] {
				valid = false
			}
			newGrid[i][j] = valid
		}
	}

	return newGrid
}

func fullgrid(m, n int) []string {
	res := make([]string, m)
	for i := 0; i < m; i++ {
		s := ""
		for j := 0; j < n; j++ {
			s += "O"
		}
		res[i] = s
	}
	return res
}

func gridString(grid [][]bool, m, n int) []string {
	res := make([]string, m)
	for i := 0; i < m; i++ {
		s := ""
		for j := 0; j < n; j++ {
			if grid[i][j] {
				s += "O"
			} else {
				s += "."
			}
		}
		res[i] = s
	}
	return res
}

func bomberMan(t int32, grid []string) []string {
	if t == 1 {
		return grid
	}

	m, n := len(grid), len(grid[0])

	if t == 2 || t == 4 {
		return fullgrid(m, n)
	}

	igrid := make([][]bool, m)
	for ri, e := range grid {
		row := make([]bool, n)
		for ci, c := range e {
			if c != '.' {
				row[ci] = true
			}
		}
		igrid[ri] = row
	}

	gridA := detonate(igrid, m, n)

	if t == 3 {
		return gridString(gridA, m, n)
	}

	gridB := detonate(gridA, m, n)

	if t == 5 {
		return gridString(gridB, m, n)
	}

	t -= 6
	d := t % 4
	if d == 0 || d == 2 {
		return fullgrid(m, n)
	} else if d == 1 {
		return gridString(gridA, m, n)
	} else {
		return gridString(gridB, m, n)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	// cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	// c := int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
