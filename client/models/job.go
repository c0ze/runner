package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Job job

swagger:model Job
*/
type Job struct {
	NewJob

	IDStatus

	/* Time when job completed, whether it was successul or failed. Always in UTC.
	 */
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	/* Time when job was submitted. Always in UTC.

	Read Only: true
	*/
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* Image to execute to run this Job. Get details via the /image/{id} endpoint.

	Required: true
	Read Only: true
	*/
	ImageID string `json:"image_id"`

	/* reason
	 */
	Reason Reason `json:"reason,omitempty"`

	/* If this field is set, then this job was retried by the job referenced in this field.
	 */
	RetryID string `json:"retry_id,omitempty"`

	/* If this field is set, then this job is a retry of the ID in this field.

	Read Only: true
	*/
	RetryOf string `json:"retry_of,omitempty"`

	/* Time when job started execution. Always in UTC.
	 */
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.NewJob.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.IDStatus.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateImageID(formats strfmt.Registry) error {

	if err := validate.RequiredString("image_id", "body", string(m.ImageID)); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateReason(formats strfmt.Registry) error {

	if err := m.Reason.Validate(formats); err != nil {
		return err
	}

	return nil
}
