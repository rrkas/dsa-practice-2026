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
 * Complete the 'gridSearch' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING_ARRAY G
 *  2. STRING_ARRAY P
 */

// func gridSearch(G []string, P []string) string {
// 	M, N := len(G), len(G[0])
// 	m, n := len(P), len(P[0])

// 	for i := 0; i <= M-m; i++ {
// 		for j := 0; j <= N-n; j++ {
// 			if G[i][j] == P[0][0] {
// 				found := true

// 				for ii := 0; found && ii < m; ii++ {
// 					for jj := 0; found && jj < n; jj++ {
// 						if G[i+ii][j+jj] != P[ii][jj] {
// 							found = false
// 							break
// 						}
// 					}
// 				}

// 				if found {
// 					return "YES"
// 				}
// 			}
// 		}
// 	}

// 	return "NO"
// }

func gridSearch(G []string, P []string) string {
	M, m := len(G), len(P)

	for i := 0; i <= M-m; i++ {
		start := 0

		for {
			j := strings.Index(G[i][start:], P[0])
			if j == -1 {
				break
			}

			j += start
			found := true

			for k := 1; k < m; k++ {
				if G[i+k][j:j+len(P[0])] != P[k] {
					found = false
					break
				}
			}

			if found {
				return "YES"
			}

			start = j + 1
		}
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		RTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		R := int32(RTemp)

		// CTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		// C := int32(CTemp)

		var G []string

		for i := 0; i < int(R); i++ {
			GItem := readLine(reader)
			G = append(G, GItem)
		}

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		rTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		r := int32(rTemp)

		// cTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)
		// c := int32(cTemp)

		var P []string

		for i := 0; i < int(r); i++ {
			PItem := readLine(reader)
			P = append(P, PItem)
		}

		result := gridSearch(G, P)

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
