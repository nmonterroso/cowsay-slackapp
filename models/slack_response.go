package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*SLACKResponse slack response

swagger:model SlackResponse
*/
type SLACKResponse struct {

	/* ResponseType response type

	Required: true
	*/
	ResponseType string `json:"response_type,omitempty"`

	/* Text text

	Required: true
	*/
	Text string `json:"text,omitempty"`
}

// Validate validates this slack response
func (m *SLACKResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponseType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var slackResponseResponseTypeEnum []interface{}

func (m *SLACKResponse) validateResponseTypeEnum(path, location string, value string) error {
	if slackResponseResponseTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["in_channel","ephemeral"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			slackResponseResponseTypeEnum = append(slackResponseResponseTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, slackResponseResponseTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *SLACKResponse) validateResponseType(formats strfmt.Registry) error {

	if err := validate.Required("response_type", "body", string(m.ResponseType)); err != nil {
		return err
	}

	if err := m.validateResponseTypeEnum("response_type", "body", m.ResponseType); err != nil {
		return err
	}

	return nil
}

func (m *SLACKResponse) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", string(m.Text)); err != nil {
		return err
	}

	return nil
}