package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'jimOrders' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts 2D_INTEGER_ARRAY orders as parameter.
 */

func jimOrders(orders [][]int32) []int32 {

    type Order struct {
        serveTime int32
        customer  int32
    }

    arr := make([]Order, len(orders))

    for i := 0; i < len(orders); i++ {
        arr[i] = Order{
            serveTime: orders[i][0] + orders[i][1],
            customer:  int32(i + 1),
        }
    }

    sort.Slice(arr, func(i, j int) bool {
        if arr[i].serveTime == arr[j].serveTime {
            return arr[i].customer < arr[j].customer
        }
        return arr[i].serveTime < arr[j].serveTime
    })

    result := make([]int32, len(arr))
    for i := range arr {
        result[i] = arr[i].customer
    }

    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var orders [][]int32
    for i := 0; i < int(n); i++ {
        ordersRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var ordersRow []int32
        for _, ordersRowItem := range ordersRowTemp {
            ordersItemTemp, err := strconv.ParseInt(ordersRowItem, 10, 64)
            checkError(err)
            ordersItem := int32(ordersItemTemp)
            ordersRow = append(ordersRow, ordersItem)
        }

        if len(ordersRow) != 2 {
            panic("Bad input")
        }

        orders = append(orders, ordersRow)
    }

    result := jimOrders(orders)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
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
