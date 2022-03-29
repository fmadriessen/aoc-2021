package helper

import (
	"bufio"
	"log"
	"os"
)

// Read lines from AoC input and return the lines as a slice
func ReadInput(input string) (result []string) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	// defer defers the execution of the below function until the surrounding function returns
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
