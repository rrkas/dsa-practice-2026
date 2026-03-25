package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
 * Complete the 'organizingContainers' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts 2D_INTEGER_ARRAY container as parameter.
 */

func organizingContainers(container [][]int64) string {
	n := len(container)
	rowSums, colSums := make([]int64, n), make([]int64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rowSums[i] += container[i][j]
			colSums[j] += container[i][j]
		}
	}

	slices.Sort(rowSums)
	slices.Sort(colSums)

	for i := 0; i < n; i++ {
		if rowSums[i] < colSums[i] {
			return "Impossible"
		}
	}

	return "Possible"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var container [][]int64
		for i := 0; i < int(n); i++ {
			containerRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var containerRow []int64
			for _, containerRowItem := range containerRowTemp {
				containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
				checkError(err)
				containerItem := int64(containerItemTemp)
				containerRow = append(containerRow, containerItem)
			}

			if len(containerRow) != int(n) {
				panic("Bad input")
			}

			container = append(container, containerRow)
		}

		result := organizingContainers(container)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
