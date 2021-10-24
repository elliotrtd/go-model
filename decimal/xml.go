package decimal

import "encoding/xml"

// UnmarshalXML parses an XML string
func (s *Decimal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	x, err := FromString(str)
	if err != nil {
		return err
	}
	*s = x
	return nil
}
