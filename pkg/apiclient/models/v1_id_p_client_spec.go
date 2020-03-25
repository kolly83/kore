// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IDPClientSpec v1 ID p client spec
//
// swagger:model v1.IDPClientSpec
type V1IDPClientSpec struct {

	// display name
	// Required: true
	DisplayName *string `json:"displayName"`

	// id
	// Required: true
	ID *string `json:"id"`

	// redirect u r is
	// Required: true
	RedirectURIs []string `json:"redirectURIs"`

	// secret
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this v1 ID p client spec
func (m *V1IDPClientSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectURIs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IDPClientSpec) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *V1IDPClientSpec) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1IDPClientSpec) validateRedirectURIs(formats strfmt.Registry) error {

	if err := validate.Required("redirectURIs", "body", m.RedirectURIs); err != nil {
		return err
	}

	return nil
}

func (m *V1IDPClientSpec) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", m.Secret); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IDPClientSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IDPClientSpec) UnmarshalBinary(b []byte) error {
	var res V1IDPClientSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
