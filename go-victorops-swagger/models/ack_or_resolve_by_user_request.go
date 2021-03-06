// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AckOrResolveByUserRequest ack or resolve by user request
// swagger:model AckOrResolveByUserRequest
type AckOrResolveByUserRequest struct {

	// message
	Message string `json:"message,omitempty"`

	// user name
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this ack or resolve by user request
func (m *AckOrResolveByUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AckOrResolveByUserRequest) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AckOrResolveByUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AckOrResolveByUserRequest) UnmarshalBinary(b []byte) error {
	var res AckOrResolveByUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
