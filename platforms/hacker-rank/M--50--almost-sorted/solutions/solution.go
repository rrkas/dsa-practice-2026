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
 * Complete the 'almostSorted' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func check(arr []int32) bool {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func reverse(arr []int32, l, r int) []int32 {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

func almostSorted(arr []int32) {
	n := len(arr)
	l, r := -1, -1
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			if l == -1 {
				l = i
			}
			r = i + 1
		}
	}

	if l == -1 {
		fmt.Println("yes")
		return
	}

	// swap
	arr[l], arr[r] = arr[r], arr[l]
	if check(arr) {
		fmt.Println("yes")
		fmt.Printf("swap %v %v", l+1, r+1)
		return
	}
	arr[l], arr[r] = arr[r], arr[l]

	// reverse
	arr1 := reverse(arr, l, r)
	if check(arr1) {
		fmt.Println("yes")
		fmt.Printf("reverse %v %v", l+1, r+1)
		return
	}

	fmt.Println("no")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	almostSorted(arr)
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
