package dom

type ProcessingInstruction interface {
	Node

	// accessors
	GetTarget() string
	SetTarget(target string)
	GetData() string
	SetData(data string)
	// no spec-defined functions
}

type ProcessingInstructionImpl struct {
	NodeImpl
	target string
	data   string
}

func (p *ProcessingInstructionImpl) GetTarget() string {
	return p.target
}

func (p *ProcessingInstructionImpl) SetTarget(target string) {
	p.target = target
}

func (p *ProcessingInstructionImpl) GetData() string {
	return p.data
}

func (p *ProcessingInstructionImpl) SetData(data string) {
	p.data = data
}
