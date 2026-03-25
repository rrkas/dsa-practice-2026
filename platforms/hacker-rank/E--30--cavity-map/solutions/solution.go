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
 * Complete the 'cavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

// func cavityMap(grid []string) []string {
// 	m, n := len(grid), len(grid[0])
// 	res := make([]string, m)

// 	for i := 0; i < m; i++ {
// 		s := strings.Builder{}
// 		for j := 0; j < n; j++ {
// 			if i == 0 || j == 0 || i == m-1 || j == n-1 {
// 				s.WriteByte(grid[i][j])
// 			} else {
// 				isCavity := true

// 				isCavity = isCavity && (grid[i][j] > grid[i-1][j])
// 				isCavity = isCavity && (grid[i][j] > grid[i+1][j])
// 				isCavity = isCavity && (grid[i][j] > grid[i][j-1])
// 				isCavity = isCavity && (grid[i][j] > grid[i][j+1])

// 				if isCavity {
// 					s.WriteByte('X')
// 				} else {
// 					s.WriteByte(grid[i][j])
// 				}
// 			}
// 		}
// 		res[i] = s.String()
// 	}

// 	return res
// }

func cavityMap(grid []string) []string {
    m, n := len(grid), len(grid[0])
    res := make([]string, m)

    for i := 0; i < m; i++ {
        row := []byte(grid[i]) // faster than builder

        for j := 0; j < n; j++ {
            if i == 0 || j == 0 || i == m-1 || j == n-1 {
                continue
            }

            curr := grid[i][j]

            if curr > grid[i-1][j] &&
               curr > grid[i+1][j] &&
               curr > grid[i][j-1] &&
               curr > grid[i][j+1] {
                row[j] = 'X'
            }
        }

        res[i] = string(row)
    }

    return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

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
