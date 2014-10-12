package dom

type Text interface {
	CharacterData

	// accessors
	GetIsElementContentWhitespace() bool
	SetIsElementContentWhitespace(isElementContentWhitespace bool)
	GetWholeText() string
	SetWholeText(wholeText string)
	// spec-defined functions
	SplitText(offset int32) Text
	ReplaceWholeText(content string) Text
}

type TextImpl struct {
	CharacterDataImpl
	isElementContentWhitespace bool
	wholeText                  string
}

func (t *TextImpl) GetIsElementContentWhitespace() bool {
	return t.isElementContentWhitespace
}

func (t *TextImpl) SetIsElementContentWhitespace(isElementContentWhitespace bool) {
	t.isElementContentWhitespace = isElementContentWhitespace
}

func (t *TextImpl) GetWholeText() string {
	return t.wholeText
}

func (t *TextImpl) SetWholeText(wholeText string) {
	t.wholeText = wholeText
}

func (t *TextImpl) SplitText(offset int32) Text {
	panic("unimplemented") // TODO

	return nil
}

func (t *TextImpl) ReplaceWholeText(content string) Text {
	panic("unimplemented") // TODO

	return nil
}
