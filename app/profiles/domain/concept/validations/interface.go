package validations

// ValidationProvider provides validation behaviour
type ValidationProvider interface {
	Validate() error
}
