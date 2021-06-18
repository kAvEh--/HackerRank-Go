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

// Arr ...
type Arr struct {
	value int32
	index int
}

// ByValue ...
type ByValue []Arr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func minimumSwaps(input []int32) int32 {
	arr := make([]Arr, 0, len(input))

	// step 1. Generate an array of sorted indexes
	for i, v := range input {
		arr = append(arr, Arr{v, i})
	}

	sort.Sort(ByValue(arr))

	idx := make([]int, 0, len(input))
	for _, ar := range arr {
		idx = append(idx, ar.index)
	}

	// step 2. Sort the array by sorted indexes
	var result int32
	for i := 0; i < len(input); i++ {
		if i == idx[i] {
			continue
		}

		input[i], input[idx[i]] = input[idx[i]], input[i]
		idx[i], idx[input[idx[i]]-1] = i, idx[i]
		result++
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
