package dom

type IText interface {
	ICharacterData

	// no accessors
	// spec-defined functions
	SplitText(offset int32) IText
}

type Text struct {
	CharacterData
}

func (t Text) SplitText(offset int32) IText {
	return nil
}
