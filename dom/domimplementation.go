package dom

type IDOMImplementation interface {
	// no accessors
	// spec-defined functions
	HasFeature(feature string, version string) bool
	CreateDocumentType(qualifiedName string, publicId string, systemId string) IDocumentType
	CreateDocument(namespaceURI string, qualifiedName string, doctype IDocumentType) IDocument
}

type DOMImplementation struct {
}

func (d DOMImplementation) HasFeature(feature string, version string) bool {
	return false
}

func (d DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) IDocumentType {
	return nil
}

func (d DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype IDocumentType) IDocument {
	return nil
}
