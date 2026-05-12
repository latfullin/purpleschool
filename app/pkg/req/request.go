package req

type Validatable interface {
	Validate() error
}
