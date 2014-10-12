package serialiser

import (
	"errors"
	"strings"

	"fknsrs.biz/p/xml/dom"
)

var (
	ERR_UNHANDLED_NODE = errors.New("unhandled node type")
)

func Serialise(node dom.INode) (string, error) {
	switch t := node.(type) {
	case *dom.Document:
		return Serialise(t.GetDocumentElement())
	default:
		if xml, err := serialise(node); err != nil {
			return "", err
		} else {
			return strings.Join(xml, ""), nil
		}
	}
}

func serialise(node dom.INode) ([]string, error) {
	switch t := node.(type) {
	case *dom.Element:
		xml := []string{}

		xml = append(xml, "<"+t.GetNodeName()+">")

		childNodes := t.GetChildNodes()
		l := childNodes.GetLength()
		for i := int32(0); i < l; i++ {
			if childXml, err := serialise(childNodes.Item(i)); err != nil {
				return nil, err
			} else {
				xml = append(xml, childXml...)
			}
		}

		xml = append(xml, "</"+t.GetNodeName()+">")

		return xml, nil
	}

	return nil, ERR_UNHANDLED_NODE
}
