package custom_error

type LinkNotFoundError struct {
	Error string
}

func NewLinkNotFound(error string) LinkNotFoundError {
	return LinkNotFoundError{Error: error}
}
