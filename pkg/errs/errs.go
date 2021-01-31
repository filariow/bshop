package errs

type ErrValidation struct {
	Errors map[string]string
}

func (e *ErrValidation) Error() string {
	return "Validation errors occurred"
}
