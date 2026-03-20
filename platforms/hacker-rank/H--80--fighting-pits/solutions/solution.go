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
 * Complete the 'fightingPits' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. 2D_INTEGER_ARRAY fighters
 *  3. 2D_INTEGER_ARRAY queries
 */

import (
	"sort"
	// "strconv"
)

func fightingPits(k int32, fighters [][]int32, queries [][]int32) []int32 {
	teams := make([][]int32, k+1)
	prefix := make([][]int64, k+1)

	// Build teams
	for _, f := range fighters {
		p, t := f[0], f[1]
		teams[t] = append(teams[t], p)
	}

	// Sort descending + prefix sum
	for i := int32(1); i <= k; i++ {
		sort.Slice(teams[i], func(a, b int) bool {
			return teams[i][a] > teams[i][b]
		})

		prefix[i] = make([]int64, len(teams[i])+1)
		for j := 0; j < len(teams[i]); j++ {
			prefix[i][j+1] = prefix[i][j] + int64(teams[i][j])
		}
	}

	memo := make(map[string]int32)
	var res []int32

	for _, q := range queries {
		if q[0] == 1 {
			p, t := q[1], q[2]

			teams[t] = append(teams[t], p)
			sort.Slice(teams[t], func(a, b int) bool {
				return teams[t][a] > teams[t][b]
			})

			prefix[t] = make([]int64, len(teams[t])+1)
			for j := 0; j < len(teams[t]); j++ {
				prefix[t][j+1] = prefix[t][j] + int64(teams[t][j])
			}

		} else {
			x, y := q[1], q[2]

			key := strconv.Itoa(int(x)) + "_" + strconv.Itoa(int(y)) +
				"_" + strconv.Itoa(len(teams[x])) + "_" + strconv.Itoa(len(teams[y]))

			if v, ok := memo[key]; ok {
				res = append(res, v)
				continue
			}

			A, B := teams[x], teams[y]
			preA, preB := prefix[x], prefix[y]

			i, j := 0, 0
			turnA := true

			for i < len(A) && j < len(B) {
				if turnA {
					power := int64(A[i])

					// binary search on B
					lo, hi := j+1, len(B)
					for lo <= hi {
						mid := (lo + hi) / 2
						if preB[mid]-preB[j] <= power {
							lo = mid + 1
						} else {
							hi = mid - 1
						}
					}

					j = lo - 1

					if j >= len(B) {
						memo[key] = x
						res = append(res, x)
						break
					}

				} else {
					power := int64(B[j])

					lo, hi := i+1, len(A)
					for lo <= hi {
						mid := (lo + hi) / 2
						if preA[mid]-preA[i] <= power {
							lo = mid + 1
						} else {
							hi = mid - 1
						}
					}

					i = lo - 1

					if i >= len(A) {
						memo[key] = y
						res = append(res, y)
						break
					}
				}

				turnA = !turnA
			}
		}
	}

	return res
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

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var fighters [][]int32
	for i := 0; i < int(n); i++ {
		fightersRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var fightersRow []int32
		for _, fightersRowItem := range fightersRowTemp {
			fightersItemTemp, err := strconv.ParseInt(fightersRowItem, 10, 64)
			checkError(err)
			fightersItem := int32(fightersItemTemp)
			fightersRow = append(fightersRow, fightersItem)
		}

		if len(fightersRow) != 2 {
			panic("Bad input")
		}

		fighters = append(fighters, fightersRow)
	}

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := fightingPits(k, fighters, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
