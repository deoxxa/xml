package dom

type IDOMImplementation interface {
	// no accessors
	// spec-defined functions
	HasFeature(feature string, version string) bool
	CreateDocumentType(qualifiedName string, publicId string, systemId string) IDocumentType
	CreateDocument(namespaceURI string, qualifiedName string, doctype IDocumentType) IDocument
	GetFeature(feature string, version string) IDOMObject
}

type DOMImplementation struct {
	features map[string][]string
}

func NewDOMImplementation() *DOMImplementation {
	return &DOMImplementation{
		features: map[string][]string{},
	}
}

func (d *DOMImplementation) HasFeature(feature string, version string) bool {
	if _feature, ok := d.features[feature]; ok {
		if version == "" {
			return true
		}

		for _, _version := range _feature {
			if _version == version {
				return true
			}
		}
	}

	return false
}

func (d *DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) IDocumentType {
	return &DocumentType{
		name:     qualifiedName,
		publicId: publicId,
		systemId: systemId,
	}
}

func (d *DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype IDocumentType) IDocument {
	document := &Document{
		Node: Node{
			nodeName:   "#document",
			childNodes: NodeList{},
		},
		implementation: d,
		doctype:        doctype,
	}

	if doctype != nil {
		document.AppendChild(doctype)
	}

	if qualifiedName != "" {
		document.AppendChild(document.CreateElementNS(namespaceURI, qualifiedName))
	}

	return document
}

func (d *DOMImplementation) GetFeature(feature string, version string) IDOMObject {
	panic("unimplemented") // TODO

	return nil
}
