package validations

import (
	"errors"

	e "../../entity"
)

const (
	msgNameInvalid    string = "Name needs to be between 2 and 40 characters"
	msgAddressInvalid string = "Address needs to be between 20 and 200 characters"
)

// Provide is a wrapper which provides Validations around an entity
type Profile struct {
	entity *e.Profile
}

// Validate validates a profile entity
func (p Profile) Validate() error {
	return validateProfile(p.entity)
}

// validate profile cycles through all validation requirements and produces an error if they fail
func validateProfile(profile *e.Profile) error {

	// validate firstname
	if len(profile.FirstName) < 2 || len(profile.FirstName) > 40 {
		return errors.New(msgNameInvalid)
	}

	// validate lastname
	if len(profile.LastName) < 2 || len(profile.LastName) > 40 {
		return errors.New(msgNameInvalid)
	}

	// validate address
	if len(profile.Address) < 20 || len(profile.Address) > 200 {
		return errors.New(msgAddressInvalid)
	}

	return nil
}
