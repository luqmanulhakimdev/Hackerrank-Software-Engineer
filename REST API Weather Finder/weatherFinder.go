package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
 * Complete the 'getTemperature' function below.
 *
 * URL for cut and paste
 * https://jsonmock.hackerrank.com/api/weather?name=<name>
 *
 * The function is expected to return an Integer.
 * The function accepts a singe parameter name.
 */

// Response struct to hold the weather data
type WeatherResponse struct {
	Data []struct {
		Weather string `json:"weather"`
	} `json:"data"`
}

func getTemperature(name string) int32 {
	// Create the URL with the provided city name
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/weather?name=%s", name)

	// Send a GET request to the weather API
	resp, err := http.Get(url)
	if err != nil {
		// Handle error while sending the request
		fmt.Printf("An error occurred: %v\n", err)
		return -1
	}
	defer resp.Body.Close()

	// Check if the response status is successful
	if resp.StatusCode != http.StatusOK {
		// If not, return an error message and status code
		fmt.Printf("Failed to fetch weather data for %s. Status code: %d\n", name, resp.StatusCode)
		return -1
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle error while reading the response body
		fmt.Printf("An error occurred while reading the response body: %v\n", err)
		return -1
	}

	// Declare a variable to hold the parsed weather data
	var weatherData WeatherResponse

	// Unmarshal the JSON data into the weatherData struct
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		// Handle error during unmarshalling
		fmt.Printf("An error occurred while parsing the JSON response: %v\n", err)
		return -1
	}

	// Extract the weather description
	if len(weatherData.Data) > 0 {
		// Split the weather string and take the first part (integer)
		weatherParts := strings.Split(weatherData.Data[0].Weather, " ")
		if len(weatherParts) > 0 {
			// Convert the first part of the weather string to integer
			var temperature int
			fmt.Sscanf(weatherParts[0], "%d", &temperature)
			return int32(temperature)
		}
	}

	// Return -1 if something goes wrong
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	name := readLine(reader)

	result := getTemperature(name)

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
