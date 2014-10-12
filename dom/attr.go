package dom

type IAttr interface {
	INode

	// accessors
	GetName() string
	SetName(name string)
	GetSpecified() bool
	SetSpecified(specified bool)
	GetValue() string
	SetValue(value string)
	GetOwnerElement() IElement
	SetOwnerElement(ownerElement IElement)
	// no spec-defined functions
}

type Attr struct {
	Node
	name         string
	specified    bool
	value        string
	ownerElement IElement
}

func (a Attr) GetName() string {
	return a.name
}

func (a Attr) SetName(name string) {
	a.name = name
}

func (a Attr) GetSpecified() bool {
	return a.specified
}

func (a Attr) SetSpecified(specified bool) {
	a.specified = specified
}

func (a Attr) GetValue() string {
	return a.value
}

func (a Attr) SetValue(value string) {
	a.value = value
}

func (a Attr) GetOwnerElement() IElement {
	return a.ownerElement
}

func (a Attr) SetOwnerElement(ownerElement IElement) {
	a.ownerElement = ownerElement
}
