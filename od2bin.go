package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file containing the octal dump
	inputFile, err := os.Open("octal_dump.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Create the output binary file
	outputFile, err := os.Create("output.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into parts
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// Skip the first part if it looks like an offset (all digits and valid octal)
		startIndex := 0
		if isOctalOffset(parts[0]) {
			startIndex = 1
		}

		// Process each octal number starting from the valid data
		for _, part := range parts[startIndex:] {
			// Convert the octal string to an integer (16-bit unsigned value)
			val, err := strconv.ParseUint(part, 8, 16)
			if err != nil {
				log.Fatalf("Error converting octal to integer: %v", err)
			}

			// Write the two bytes in little-endian order: lower byte first, then higher byte.
			err = writeBinaryBytes(outputFile, byte(val&0xFF), byte(val>>8))
			if err != nil {
				log.Fatalf("Error writing bytes to file: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Binary file created successfully.")
}

// Helper function to write two bytes to the output file in little-endian order
func writeBinaryBytes(outputFile *os.File, byte1, byte2 byte) error {
	_, err := outputFile.Write([]byte{byte1, byte2})
	return err
}

// Helper function to check if a string is a valid octal offset
func isOctalOffset(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64) // Offsets are usually decimal numbers
	return err == nil
}
