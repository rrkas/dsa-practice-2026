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
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func totDays(year int32) int32 {
	var days int32 = 365

	if year == 1918 {
		days -= 13
	} else if year < 1918 {
		if year%4 == 0 {
			days++
		}
	} else {
		if year%400 == 0 {
			days++
		} else if year%4 == 0 && year%100 != 0 {
			days++
		}
	}

	return days
}

func dayOfProgrammer(year int32) string {
	days := totDays(year)
	if days == 352 {
		return "26.09.1918"
	} else if days == 366 {
		return fmt.Sprintf("12.09.%v", year)
	} else {
		return fmt.Sprintf("13.09.%v", year)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
