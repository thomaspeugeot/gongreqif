package schema

import (
	"encoding/xml"
	"log"
	"time"
)

type CustomDate time.Time

func (c *CustomDate) UnmarshalXMLAttr(attr xml.Attr) error {
	const myFormat = "2006-01-02T15:04:05.000-07:00"
	parse, err := time.Parse(myFormat, attr.Value)
	if err != nil {
		return err
	}
	*c = CustomDate(parse)
	log.Println("Date time", time.Time(*c).Local().Format(myFormat))
	return nil
}

func (c CustomDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	const myFormat = "2006-01-02T15:04:05.000-07:00"
	return xml.Attr{
		Name:  name,
		Value: time.Time(c).Format(myFormat),
	}, nil
}
