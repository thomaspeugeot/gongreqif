package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type GenericElement struct {
	XMLName xml.Name
	Attrs   []xml.Attr       `xml:",any,attr"`
	Content []byte           `xml:",innerxml"`
	Nodes   []GenericElement `xml:",any"`
}

func main() {
	// Open the input XML file
	inputFile, err := os.Open("sample 2.reqif")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Parse the XML file
	var root GenericElement
	decoder := xml.NewDecoder(inputFile)
	err = decoder.Decode(&root)
	if err != nil && err != io.EOF {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Process the XML tree
	processElement(&root.Nodes[0])
	processElement(&root.Nodes[1])

	// Create the output XML file
	outputFile, err := os.Create("sample 2 modified.reqif")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Encode the modified XML tree to the output file
	encoder := xml.NewEncoder(outputFile)
	encoder.Indent("", "  ")
	err = encoder.Encode(root)
	if err != nil {
		fmt.Println("Error encoding XML:", err)
		return
	}

	fmt.Println("XML processing complete. Modified XML written to output.xml")
}

// processElement processes an XML element, prefixing all attributes with "reqif.xsd."
func processElement(element *GenericElement) {
	for i, attr := range element.Attrs {
		element.Attrs[i].Name.Local = "reqif.xsd:" + attr.Name.Local
	}
	for i := range element.Nodes {
		processElement(&element.Nodes[i])
	}
}
