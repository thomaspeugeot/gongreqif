package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

func main() {
	var filename string

	if len(os.Args) < 2 {
		filename = "../../../test/sample.reqif"
	} else {
		filename = os.Args[1]
	}

	// Replace "path/to/file.reqif" with the actual path to your file
	data, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var reqif schema.REQIF
	err = xml.Unmarshal(data, &reqif)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	// Access and process parsed data
	fmt.Println("ReqIF Header:")
	fmt.Printf("  Identifier: %s\n", reqif.HEADER.REQIFHEADER.IDENTIFIERAttr)
	// fmt.Printf("  Creation Time: %s\n", reqif.Header.ReqIFHeader.CreationTime)
	// fmt.Println("Content:")
	// fmt.Printf("  Number of Datatypes: %d\n", len(reqif.Content.ReqIFContent.DataTypes))

	// Iterate and process other elements as needed
	// for _, specObject := range reqif.Content.ReqIFContent.SpecObjects {
	// 	fmt.Printf("  Spec Object: %s\n", specObject.LongName)
	// 	for _, value := range specObject.Values {
	// 		_ = value
	// 		// // Handle different attribute value types
	// 		// switch value.(type) {
	// 		// case AttributeValueXHTML:
	// 		// 	// Process AttributeValueXHTML data
	// 		// case AttributeValueEnumeration:
	// 		// 	// Process AttributeValueEnumeration data
	// 		// }
	// 	}
	// }

	// ... Implement further processing logic based on your needs ...
}

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
