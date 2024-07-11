package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Define structs to represent elements in the ReqIF schema
type ReqIF struct {
	XMLName        xml.Name   `xml:"REQ-IF"`
	Header         Header     `xml:"THE-HEADER"`
	Content        Content    `xml:"CORE-CONTENT"`
	ToolExtensions []struct{} `xml:"TOOL-EXTENSIONS"`
}

type Header struct {
	ReqIFHeader ReqIFHeader `xml:"REQ-IF-HEADER"`
}

type ReqIFHeader struct {
	Identifier   string `xml:"IDENTIFIER,attr"`
	CreationTime string `xml:"CREATION-TIME"`
	RepositoryID string `xml:"REPOSITORY-ID"`
	ReqIFToolID  string `xml:"REQ-IF-TOOL-ID"`
	ReqIFVersion string `xml:"REQ-IF-VERSION"`
	SourceToolID string `xml:"SOURCE-TOOL-ID"`
	Title        string `xml:"TITLE"`
}

type Content struct {
	ReqIFContent ReqIFContent `xml:"REQ-IF-CONTENT"`
}

type ReqIFContent struct {
	DataTypes     []DataType    `xml:"DATATYPES"`
	SpecTypes     []SpecType    `xml:"SPEC-TYPES"`
	SpecObjects   []SpecObject  `xml:"SPEC-OBJECTS"`
	SpecRelations []interface{} `xml:"SPEC-RELATIONS"` // Interface to handle potential future elements
	// Specifications     []Specification `xml:"SPECIFICATIONS"`
	SpecRelationGroups []interface{} `xml:"SPEC-RELATION-GROUPS"` // Interface to handle potential future elements
}

type DataType struct {
	// ... Define nested structs for DataTypeDefinitionXHTML and DataTypeDefinitionEnumeration ...
}

type SpecType struct {
	// ... Define nested structs for SpecObjectType and SpecificationType ...
}

type SpecObject struct {
	Identifier string            `xml:"IDENTIFIER,attr"`
	LastChange string            `xml:"LAST-CHANGE,attr"`
	LongName   string            `xml:"LONG-NAME"`
	Values     []AttributeValue  `xml:"VALUES"`
	Type       SpecObjectTypeRef `xml:"TYPE"`
}

type AttributeValue struct {
	// ... Define nested structs for AttributeValueXHTML and AttributeValueEnumeration ...
}

type SpecObjectTypeRef struct {
	Ref string `xml:"SPEC-OBJECT-TYPE-REF"`
}

// ... Define similar structs for other nested elements ...

func main() {
	// Replace "path/to/file.reqif" with the actual path to your file
	data, err := readFile("../../../test/sample.reqif")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var reqif ReqIF
	err = xml.Unmarshal(data, &reqif)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	// Access and process parsed data
	fmt.Println("ReqIF Header:")
	fmt.Printf("  Identifier: %s\n", reqif.Header.ReqIFHeader.Identifier)
	fmt.Printf("  Creation Time: %s\n", reqif.Header.ReqIFHeader.CreationTime)
	fmt.Println("Content:")
	fmt.Printf("  Number of Datatypes: %d\n", len(reqif.Content.ReqIFContent.DataTypes))

	// Iterate and process other elements as needed
	for _, specObject := range reqif.Content.ReqIFContent.SpecObjects {
		fmt.Printf("  Spec Object: %s\n", specObject.LongName)
		for _, value := range specObject.Values {
			_ = value
			// // Handle different attribute value types
			// switch value.(type) {
			// case AttributeValueXHTML:
			// 	// Process AttributeValueXHTML data
			// case AttributeValueEnumeration:
			// 	// Process AttributeValueEnumeration data
			// }
		}
	}

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
