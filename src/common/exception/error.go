package exception

type Error struct {
	message string
	code    int
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}
