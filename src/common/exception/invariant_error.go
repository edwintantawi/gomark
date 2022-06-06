package exception

func NewInvariantError(error string) Error {
	return Error{message: error, code: 400}
}
