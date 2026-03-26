package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'twoPluses' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY grid as parameter.
 */

type plus struct {
	area  int32
	cells map[int]interface{}
}

func minarr(arr []int) int {
	v := arr[0]
	for _, e := range arr[1:] {
		v = min(v, e)
	}
	return v
}

func plusoverlap(p1, p2 plus) bool {
	for k := range p1.cells {
		if _, ok := p2.cells[k]; ok {
			return true
		}
	}
	return false
}

func twoPluses(grid []string) int32 {
	m, n := len(grid), len(grid[0])

	u, d, l, r := make([][]int, m), make([][]int, m), make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		u[i] = make([]int, n)
		d[i] = make([]int, n)
		l[i] = make([]int, n)
		r[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 'G' {
				continue
			}

			u[i][j] = 1
			l[i][j] = 1

			if i > 0 {
				u[i][j] += u[i-1][j]
			}

			if j > 0 {
				l[i][j] += l[i][j-1]
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != 'G' {
				continue
			}

			d[i][j] = 1
			r[i][j] = 1

			if i+1 < m {
				d[i][j] += d[i+1][j]
			}

			if j+1 < n {
				r[i][j] += r[i][j+1]
			}
		}
	}

	maskidx := func(i, j int) int {
		return i*n + j
	}

	pluses := make([]plus, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 'G' {
				continue
			}

			maxarm := minarr([]int{
				u[i][j],
				d[i][j],
				l[i][j],
				r[i][j],
			}) - 1

			for k := 0; k <= maxarm; k++ {
				mask := make(map[int]interface{})

				mask[maskidx(i, j)] = nil
				for d_ := 0; d_ <= k; d_++ {
					mask[maskidx(i-d_, j)] = nil
					mask[maskidx(i+d_, j)] = nil
					mask[maskidx(i, j-d_)] = nil
					mask[maskidx(i, j+d_)] = nil
				}

				area := int32(1 + 4*k)
				pluses = append(pluses, plus{cells: mask, area: area})
			}
		}
	}

	sort.Slice(pluses, func(i, j int) bool {
		return pluses[i].area > pluses[j].area
	})

	maxProd := int32(0)

	for i := 0; i < len(pluses)-1; i++ {
		if pluses[i].area*pluses[i+1].area <= maxProd {
			break
		}
		for j := i + 1; j < len(pluses); j++ {
			prod := pluses[i].area * pluses[j].area
			if prod <= maxProd {
				break
			}

			if !plusoverlap(pluses[i], pluses[j]) {
				maxProd = prod
			}
		}
	}

	return maxProd
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	// mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	// m := int32(mTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := twoPluses(grid)

	fmt.Fprintf(writer, "%d\n", result)

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
