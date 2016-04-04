package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Image image

swagger:model Image
*/
type Image struct {

	/* Time when image first used/created.

	Read Only: true
	*/
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* Docker image to use for job, including the tag.

	Required: true
	Read Only: true
	*/
	Image string `json:"image"`

	/* Name of this image/package/code. Can and should be different than the image name (shoudn't include tag). TODO: Should we strip tag automatically if only image is passed in?

	Read Only: true
	*/
	Name string `json:"name,omitempty"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	return nil
}
