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
 * Complete the 'boardCutting' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost_y
 *  2. INTEGER_ARRAY cost_x
 */

func boardCutting(cost_y []int32, cost_x []int32) int32 {
	sort.Slice(cost_x, func(i, j int) bool {
		return cost_x[j] < cost_x[i]
	})
	sort.Slice(cost_y, func(i, j int) bool {
		return cost_y[j] < cost_y[i]
	})

	var MOD int = 1000000007
	var m, n int = len(cost_x), len(cost_y)
	var i, j int = 0, 0
	var h, v int = 1, 1
	var cost int = 0

	for i < m && j < n {
		if cost_x[i] > cost_y[j] {
			cost += int(cost_x[i]) * h
			v++
			i++
		} else {
			cost += int(cost_y[j]) * v
			h++
			j++
		}
		cost %= MOD
	}

	for i < m {
		cost += int(cost_x[i]) * h
		v++
		i++
		cost %= MOD
	}

	for j < n {
		cost += int(cost_y[j]) * v
		h++
		j++
		cost %= MOD
	}

	return int32(cost)
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
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		m := int32(mTemp)

		nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		n := int32(nTemp)

		cost_yTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost_y []int32

		for i := 0; i < int(m)-1; i++ {
			cost_yItemTemp, err := strconv.ParseInt(cost_yTemp[i], 10, 64)
			checkError(err)
			cost_yItem := int32(cost_yItemTemp)
			cost_y = append(cost_y, cost_yItem)
		}

		cost_xTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost_x []int32

		for i := 0; i < int(n)-1; i++ {
			cost_xItemTemp, err := strconv.ParseInt(cost_xTemp[i], 10, 64)
			checkError(err)
			cost_xItem := int32(cost_xItemTemp)
			cost_x = append(cost_x, cost_xItem)
		}

		result := boardCutting(cost_y, cost_x)

		fmt.Fprintf(writer, "%d\n", result)
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
