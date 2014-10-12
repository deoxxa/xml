package dom

type Comment interface {
	CharacterData

	// no accessors
	// no spec-defined functions
}

type CommentImpl struct {
	CharacterDataImpl
}
