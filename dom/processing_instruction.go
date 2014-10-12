package dom

type IProcessingInstruction interface {
	INode

	// accessors
	GetTarget() string
	SetTarget(target string)
	GetData() string
	SetData(data string)
	// no spec-defined functions
}

type ProcessingInstruction struct {
	Node
	target string
	data   string
}

func (p ProcessingInstruction) GetTarget() string {
	return p.target
}

func (p ProcessingInstruction) SetTarget(target string) {
	p.target = target
}

func (p ProcessingInstruction) GetData() string {
	return p.data
}

func (p ProcessingInstruction) SetData(data string) {
	p.data = data
}
