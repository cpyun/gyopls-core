package contract

// ErrorCoder error code
type ErrorCoder interface {
	String() string
	Code() int32
}
