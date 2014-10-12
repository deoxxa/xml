package main

import (
	"fmt"
	"log"

	"fknsrs.biz/p/xml/dom"
	"fknsrs.biz/p/xml/serialiser"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var d dom.DOMImplementation

	doc := d.CreateDocument("urn:ns:1.0", "ns1:Whatever", nil)

	doc.GetDocumentElement().AppendChild(doc.CreateElementNS("urn:ns:1.0", "ns1:Attribute"))
	doc.GetDocumentElement().AppendChild(doc.CreateElementNS("urn:ns:1.0", "ns1:Attribute"))

	if xml, err := serialiser.Serialise(doc); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(xml)
	}
}
