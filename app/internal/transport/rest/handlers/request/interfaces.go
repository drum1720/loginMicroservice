package request

type Validator interface {
	Validate() error
}
