package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*Error error

swagger:model Error
*/
type Error struct {

	/* the HTTP status of the response

	Required: true
	*/
	Code int32 `json:"code,omitempty"`

	/* details about why the request failed

	Required: true
	*/
	Message string `json:"message,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", int32(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}