package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(path string) ([]byte, error) {
	// Open the file in read-only mode
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure file is closed even in case of errors

	// Allocate a buffer to hold the file data
	buffer := make([]byte, 1024) // Adjust buffer size based on expected file size
	var data []byte

	// Read the file contents in chunks
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// Reached end of file, break the loop
				break
			}
			return nil, fmt.Errorf("error reading file: %w", err)
		}

		// Append read data to the final data slice
		data = append(data, buffer[:n]...)
	}

	return data, nil
}
