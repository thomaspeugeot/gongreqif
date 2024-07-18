package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type element struct {
	XMLName xml.Name `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd emm"` // Replace with your actual namespace
	Attrs   []xml.Attr
	Content []interface{}
}

func (e *element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	e.Attrs = make([]xml.Attr, 0)
	err := d.DecodeToken(&e.XMLName)
	if err != nil {
		return err
	}
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case xml.StartElement:
			var v interface{}
			err := d.DecodeElement(&v, &t)
			if err != nil {
				return err
			}
			e.Content = append(e.Content, v)
		case xml.EndElement:
			return nil
		case xml.CharData:
			e.Content = append(e.Content, string(t))
		}
	}
}

func prefixAttributes(e *element) {
	for i := range e.Attrs {
		e.Attrs[i].Name.Space = "reqif.xsd"
	}
	for _, child := range e.Content {
		switch c := child.(type) {
		case *element:
			prefixAttributes(c)
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("your_file.xml") // Replace with your actual file path
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var root element
	err = xml.Unmarshal(data, &root)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	prefixAttributes(&root)

	output, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling modified XML:", err)
		return
	}

	fmt.Println(string(output))
}
