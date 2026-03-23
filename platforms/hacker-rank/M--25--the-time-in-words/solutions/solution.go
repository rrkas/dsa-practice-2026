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
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func numToWords(n int32) string {
	if n >= 20 {
		tens := n / 10
		tensW := map[int32]string{
			2: "twenty",
			3: "thirty",
			4: "forty",
			5: "fifty",
		}[tens]
		return tensW + " " + numToWords(n%10)
	}

	return map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}[n]
}

func timeInWords(h int32, m int32) string {
	switch m {
	case 0:
		return numToWords(h) + " o' clock"
	case 15:
		return "quarter past " + numToWords(h)
	case 30:
		return "half past " + numToWords(h)
	case 45:
		return "quarter to " + numToWords(h+1)
	case 1:
		return "one minute past " + numToWords(h)
	case 59:
		return "one minute to " + numToWords(h+1)
	default:
		if m < 30 {
			return numToWords(m) + " minutes past " + numToWords(h)
		} else {
			return numToWords(60-m) + " minutes to " + numToWords(h+1)
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
