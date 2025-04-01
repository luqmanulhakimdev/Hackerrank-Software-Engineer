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
 * Complete the 'efficientJanitor' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts FLOAT_ARRAY weight as parameter.
 */
func ShortFloat32(weight []float32) []float32 {
	// Convert the slice of float32 to a slice of float64 for sorting
	weightsFloat64 := make([]float64, len(weight))
	for i, v := range weight {
		weightsFloat64[i] = float64(v)
	}

	// Sort the float64 slice
	sort.Float64s(weightsFloat64)

	// Convert back to float32
	for i, v := range weightsFloat64 {
		weight[i] = float32(v)
	}

	return weight
}

func efficientJanitor(weight []float32) int32 {
	// Write your code here
	count := 0
	i, j := 0, len(weight)-1
	ShortFloat32(weight) // Sort the weights

	for i <= j {
		count++
		if weight[i]+weight[j] <= 3 {
			i++
		}
		j--
	}
	return int32(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	weightCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var weight []float32

	for i := 0; i < int(weightCount); i++ {
		weightItemTemp, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
		checkError(err)
		weightItem := float32(weightItemTemp)
		weight = append(weight, weightItem)
	}

	result := efficientJanitor(weight)

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
