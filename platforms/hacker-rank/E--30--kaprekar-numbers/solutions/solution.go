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
 * Complete the 'kaprekarNumbers' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER q
 */

func checkKaprekar(n int32) bool {
	if n == 1 {
		return true
	}

	// count digits
	d := 0
	t := n
	for t > 0 {
		t /= 10
		d++
	}

	sq := int64(n) * int64(n)
	pow := int64(1)
	for i := 0; i < d; i++ {
		pow *= 10
	}

	r := sq % pow
	l := sq / pow

	return int32(l+r) == n
}

func kaprekarNumbers(p int32, q int32) {
	res := make([]int32, 0)
	for i := p; i <= q; i++ {
		if checkKaprekar(i) {
			res = append(res, i)
		}
	}
	if len(res) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		for _, e := range res {
			fmt.Printf("%v ", e)
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
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
