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
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func multiply(a []int, b int) []int {
	res := make([]int, 0)
	carry := 0
	for i := 0; i < len(a); i++ {
		v := (a[i] * b) + carry
		carry = v / 10
		v = v % 10
		res = append(res, v)
	}
	for carry > 0 {
		res = append(res, carry%10)
		carry /= 10
	}
	return res
}

func narray(a int) []int {
	res := make([]int, 0)
	for a > 0 {
		d := a % 10
		res = append(res, d)
		a /= 10
	}
	return res
}

func extraLongFactorials(n int32) {
	f := narray(1)
	for i := 2; i <= int(n); i++ {
		f = multiply(f, i)
	}
	for i := len(f) - 1; i >= 0; i-- {
		fmt.Print(f[i])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
