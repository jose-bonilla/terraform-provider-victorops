// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetAPIPublicV1ProfileUsernamePoliciesOKBody get Api public v1 profile username policies o k body
// swagger:model getApiPublicV1ProfileUsernamePoliciesOKBody
type GetAPIPublicV1ProfileUsernamePoliciesOKBody struct {

	// self Url
	SelfURL string `json:"_selfUrl,omitempty"`

	// steps
	Steps UserPagingPolicy `json:"steps"`
}

// Validate validates this get Api public v1 profile username policies o k body
func (m *GetAPIPublicV1ProfileUsernamePoliciesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAPIPublicV1ProfileUsernamePoliciesOKBody) validateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	if err := m.Steps.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("steps")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAPIPublicV1ProfileUsernamePoliciesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAPIPublicV1ProfileUsernamePoliciesOKBody) UnmarshalBinary(b []byte) error {
	var res GetAPIPublicV1ProfileUsernamePoliciesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}