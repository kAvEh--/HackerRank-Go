package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("input.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	c := make([]string, int(t))
	start, end := 0, 0

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		tmp := strings.Fields(s)
		if tmp[0] == "1" {
			c[end] = tmp[1]
			end++
		} else if tmp[0] == "2" {
			start++
		} else {
			fmt.Fprintf(writer, "%s\n", c[start])
		}
	}

	writer.Flush()
}

func readLine2(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError2(err error) {
	if err != nil {
		panic(err)
	}
}
