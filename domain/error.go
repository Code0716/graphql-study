package domain

// Error is Error struct
type Error struct {
	Message string
}

// NewError create error
func NewError(msg string) Error {
	return Error{
		Message: msg,
	}
}
