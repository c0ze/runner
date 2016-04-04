package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*NewJobsWrapper new jobs wrapper

swagger:model NewJobsWrapper
*/
type NewJobsWrapper struct {

	/* jobs

	Required: true
	*/
	Jobs []*NewJob `json:"jobs"`
}

// Validate validates this new jobs wrapper
func (m *NewJobsWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewJobsWrapper) validateJobs(formats strfmt.Registry) error {

	if err := validate.Required("jobs", "body", m.Jobs); err != nil {
		return err
	}

	for i := 0; i < len(m.Jobs); i++ {

		if err := m.Jobs[i].Validate(formats); err != nil {
			return err
		}

	}

	return nil
}
