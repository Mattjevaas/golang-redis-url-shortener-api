package custom_error

type InvalidUrl struct {
	Error string
}

func NewInvalidUrl(error string) InvalidUrl {
	return InvalidUrl{Error: error}
}
