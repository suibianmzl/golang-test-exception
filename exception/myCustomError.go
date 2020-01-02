package exception

type MyCustomError string

func (e MyCustomError) Error() string {
	return e.Message()
}

func (e MyCustomError) Message() string {
	return string(e)
}
