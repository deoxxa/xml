package dom

type IDOMError interface {
	// accessors
	GetSeverity() int16
	SetSeverity(severity int16)
	GetMessage() string
	SetMessage(message string)
	GetType() string
	SetType(theType string)
	GetRelatedException() IDOMObject
	SetRelatedException(relatedException IDOMObject)
	GetRelatedData() IDOMObject
	SetRelatedData(relatedData IDOMObject)
	GetLocation() IDOMLocator
	SetLocation(location IDOMLocator)
	// no spec-defined functions
}

const (
	SEVERITY_WARNING     = 1
	SEVERITY_ERROR       = 2
	SEVERITY_FATAL_ERROR = 3
)

type DOMError struct {
	severity         int16
	message          string
	theType          string
	relatedException IDOMObject
	relatedData      IDOMObject
	location         IDOMLocator
}

func (d *DOMError) GetSeverity() int16 {
	return d.severity
}

func (d *DOMError) SetSeverity(severity int16) {
	d.severity = severity
}

func (d *DOMError) GetMessage() string {
	return d.message
}

func (d *DOMError) SetMessage(message string) {
	d.message = message
}

func (d *DOMError) GetType() string {
	return d.theType
}

func (d *DOMError) SetType(theType string) {
	d.theType = theType
}

func (d *DOMError) GetRelatedException() IDOMObject {
	return d.relatedException
}

func (d *DOMError) SetRelatedException(relatedException IDOMObject) {
	d.relatedException = relatedException
}

func (d *DOMError) GetRelatedData() IDOMObject {
	return d.relatedData
}

func (d *DOMError) SetRelatedData(relatedData IDOMObject) {
	d.relatedData = relatedData
}

func (d *DOMError) GetLocation() IDOMLocator {
	return d.location
}

func (d *DOMError) SetLocation(location IDOMLocator) {
	d.location = location
}
