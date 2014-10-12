package dom

type IComment interface {
	ICharacterData

	// no accessors
	// no spec-defined functions
}

type Comment struct {
	CharacterData
}
