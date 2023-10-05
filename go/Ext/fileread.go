package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "tidb.log"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Extract data from the line and store it in a variable.
		variable := line

		// You can do further processing with the data as needed.
		// fmt.Println(variable)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
