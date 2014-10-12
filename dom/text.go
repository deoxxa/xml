package dom

type IText interface {
	ICharacterData

	// accessors
	GetIsElementContentWhitespace() bool
	SetIsElementContentWhitespace(isElementContentWhitespace bool)
	GetWholeText() string
	SetWholeText(wholeText string)
	// spec-defined functions
	SplitText(offset int32) IText
	ReplaceWholeText(content string) IText
}

type Text struct {
	CharacterData
	isElementContentWhitespace bool
	wholeText                  string
}

func (t *Text) GetIsElementContentWhitespace() bool {
	return t.isElementContentWhitespace
}

func (t *Text) SetIsElementContentWhitespace(isElementContentWhitespace bool) {
	t.isElementContentWhitespace = isElementContentWhitespace
}

func (t *Text) GetWholeText() string {
	return t.wholeText
}

func (t *Text) SetWholeText(wholeText string) {
	t.wholeText = wholeText
}

func (t *Text) SplitText(offset int32) IText {
	panic("unimplemented") // TODO

	return nil
}

func (t *Text) ReplaceWholeText(content string) IText {
	panic("unimplemented") // TODO

	return nil
}
