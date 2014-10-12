package dom

type DOMError interface {
	// accessors
	GetSeverity() int16
	SetSeverity(severity int16)
	GetMessage() string
	SetMessage(message string)
	GetType() string
	SetType(theType string)
	GetRelatedException() DOMObject
	SetRelatedException(relatedException DOMObject)
	GetRelatedData() DOMObject
	SetRelatedData(relatedData DOMObject)
	GetLocation() DOMLocator
	SetLocation(location DOMLocator)
	// no spec-defined functions
}

const (
	SEVERITY_WARNING     = 1
	SEVERITY_ERROR       = 2
	SEVERITY_FATAL_ERROR = 3
)

type DOMErrorImpl struct {
	severity         int16
	message          string
	theType          string
	relatedException DOMObject
	relatedData      DOMObject
	location         DOMLocator
}

func (d *DOMErrorImpl) GetSeverity() int16 {
	return d.severity
}

func (d *DOMErrorImpl) SetSeverity(severity int16) {
	d.severity = severity
}

func (d *DOMErrorImpl) GetMessage() string {
	return d.message
}

func (d *DOMErrorImpl) SetMessage(message string) {
	d.message = message
}

func (d *DOMErrorImpl) GetType() string {
	return d.theType
}

func (d *DOMErrorImpl) SetType(theType string) {
	d.theType = theType
}

func (d *DOMErrorImpl) GetRelatedException() DOMObject {
	return d.relatedException
}

func (d *DOMErrorImpl) SetRelatedException(relatedException DOMObject) {
	d.relatedException = relatedException
}

func (d *DOMErrorImpl) GetRelatedData() DOMObject {
	return d.relatedData
}

func (d *DOMErrorImpl) SetRelatedData(relatedData DOMObject) {
	d.relatedData = relatedData
}

func (d *DOMErrorImpl) GetLocation() DOMLocator {
	return d.location
}

func (d *DOMErrorImpl) SetLocation(location DOMLocator) {
	d.location = location
}
