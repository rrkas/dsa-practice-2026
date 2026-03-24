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
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */

func countSubs(s1 string, s2 string) int32 {
	c := int32(0)

	for i := 0; i < len(s1); i++ {
		if s1[i] == '1' || s2[i] == '1' {
			c++
		}
	}

	return c
}

func acmTeam(topic []string) []int32 {
	maxc, cmax := int32(0), int32(0)

	n := len(topic)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			c := countSubs(topic[i], topic[j])
			if c > maxc {
				maxc = c
				cmax = 1
			} else if c == maxc {
				cmax++
			}
		}
	}

	return []int32{maxc, cmax}
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

	// _, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	// m := int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

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
