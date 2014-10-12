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
}

func (d DOMImplementation) HasFeature(feature string, version string) bool {
	panic("unimplemented") // TODO

	return false
}

func (d DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) IDocumentType {
	panic("unimplemented") // TODO

	return nil
}

func (d DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype IDocumentType) IDocument {
	panic("unimplemented") // TODO

	return nil
}

func (d DOMImplementation) GetFeature(feature string, version string) IDOMObject {
	panic("unimplemented") // TODO

	return nil
}
