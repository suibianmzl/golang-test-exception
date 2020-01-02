package exception

type UserError interface {
	error
	Message() string
}