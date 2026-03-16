package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'reverseShuffleMerge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func reverseShuffleMerge(s string) string {
	runes := []rune(s)

	freq := make(map[rune]int)
	remain := make(map[rune]int)
	for _, c := range runes {
		freq[c]++
		remain[c]++
	}
	need := make(map[rune]int, len(freq))
	for k, v := range freq {
		need[k] = v / 2
	}
	used := make(map[rune]int, len(freq))
	stack := make([]rune, 0)

	for i := len(runes) - 1; i >= 0; i-- {
		c := runes[i]
		remain[c]--
		if used[c] == need[c] {
			continue
		}

		for len(stack) > 0 {
			top := stack[len(stack)-1]

			if top <= c {
				break
			}

			if used[top]-1+remain[top] < need[top] {
				break
			}

			stack = stack[:len(stack)-1]
			used[top]--
		}

		stack = append(stack, c)
		used[c]++
	}

	sb := strings.Builder{}
	for _, c := range stack {
		sb.WriteRune(c)
	}
	return sb.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := reverseShuffleMerge(s)

	fmt.Fprintf(writer, "%s\n", result)

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
