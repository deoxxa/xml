package dom

type DOMImplementation interface {
	// no accessors
	// spec-defined functions
	HasFeature(feature string, version string) bool
	CreateDocumentType(qualifiedName string, publicId string, systemId string) DocumentType
	CreateDocument(namespaceURI string, qualifiedName string, doctype DocumentType) Document
	GetFeature(feature string, version string) DOMObject
}

type DOMImplementationImpl struct {
	features map[string][]string
}

func NewDOMImplementation() *DOMImplementationImpl {
	return &DOMImplementationImpl{
		features: map[string][]string{},
	}
}

func (d *DOMImplementationImpl) HasFeature(feature string, version string) bool {
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

func (d *DOMImplementationImpl) CreateDocumentType(qualifiedName string, publicId string, systemId string) DocumentType {
	return &DocumentTypeImpl{
		name:     qualifiedName,
		publicId: publicId,
		systemId: systemId,
	}
}

func (d *DOMImplementationImpl) CreateDocument(namespaceURI string, qualifiedName string, doctype DocumentType) Document {
	document := &DocumentImpl{
		NodeImpl: NodeImpl{
			nodeName:   "#document",
			childNodes: NodeListImpl{},
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

func (d *DOMImplementationImpl) GetFeature(feature string, version string) DOMObject {
	panic("unimplemented") // TODO

	return nil
}
