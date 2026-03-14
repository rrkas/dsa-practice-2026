package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'maximumPeople' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. LONG_INTEGER_ARRAY p
 *  2. LONG_INTEGER_ARRAY x
 *  3. LONG_INTEGER_ARRAY y
 *  4. LONG_INTEGER_ARRAY r
 */

func maximumPeople(p []int64, x []int64, y []int64, r []int64) int64 {
	n := len(p)
	m := len(y)

	type event struct {
		pos int64
		typ int // 0=start,1=town,2=end
		id  int
		pop int64
	}

	events := make([]event, 0, n+2*m)

	for i := 0; i < m; i++ {
		events = append(events, event{y[i] - r[i], 0, i, 0})
		events = append(events, event{y[i] + r[i], 2, i, 0})
	}

	for i := 0; i < n; i++ {
		events = append(events, event{x[i], 1, i, p[i]})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].pos == events[j].pos {
			return events[i].typ < events[j].typ
		}
		return events[i].pos < events[j].pos
	})

	active := map[int]struct{}{}
	unique := make([]int64, m)

	var sunny int64 = 0

	for _, e := range events {

		if e.typ == 0 {
			active[e.id] = struct{}{}
		} else if e.typ == 2 {
			delete(active, e.id)
		} else {

			if len(active) == 0 {
				sunny += e.pop
			} else if len(active) == 1 {
				for id := range active {
					unique[id] += e.pop
				}
			}
		}
	}

	var best int64 = 0
	for _, v := range unique {
		if v > best {
			best = v
		}
	}

	return sunny + best
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int64

	for i := 0; i < int(n); i++ {
		pItem, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		p = append(p, pItem)
	}

	xTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var x []int64

	for i := 0; i < int(n); i++ {
		xItem, err := strconv.ParseInt(xTemp[i], 10, 64)
		checkError(err)
		x = append(x, xItem)
	}

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	yTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var y []int64

	for i := 0; i < int(m); i++ {
		yItem, err := strconv.ParseInt(yTemp[i], 10, 64)
		checkError(err)
		y = append(y, yItem)
	}

	rTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var r []int64

	for i := 0; i < int(m); i++ {
		rItem, err := strconv.ParseInt(rTemp[i], 10, 64)
		checkError(err)
		r = append(r, rItem)
	}

	result := maximumPeople(p, x, y, r)

	fmt.Fprintf(writer, "%d\n", result)

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
