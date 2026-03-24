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
 * Complete the 'happyLadybugs' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING b as parameter.
 */

func happyLadybugs(b string) string {
	vacancies := 0
	freq := make(map[rune]int)
	for _, c := range b {
		if c != '_' {
			freq[c]++
		} else {
			vacancies++
		}
	}
	for _, v := range freq {
		if v == 1 {
			return "NO"
		}
	}

	if vacancies > 0 {
		return "YES"
	}

	for i := 0; i < len(b); i++ {
		f := false
		if i > 0 && b[i] == b[i-1] {
			f = true
		}
		if i+1 < len(b) && b[i] == b[i+1] {
			f = true
		}
		if !f {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		_, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		// n := int32(nTemp)

		b := readLine(reader)

		result := happyLadybugs(b)

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
