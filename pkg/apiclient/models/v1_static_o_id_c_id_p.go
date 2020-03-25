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

// V1StaticOIDCIDP v1 static o ID c ID p
//
// swagger:model v1.StaticOIDCIDP
type V1StaticOIDCIDP struct {

	// client ID
	// Required: true
	ClientID *string `json:"clientID"`

	// client scopes
	// Required: true
	ClientScopes []string `json:"clientScopes"`

	// client secret
	// Required: true
	ClientSecret *string `json:"clientSecret"`

	// issuer
	// Required: true
	Issuer *string `json:"issuer"`

	// user claims
	// Required: true
	UserClaims []string `json:"userClaims"`
}

// Validate validates this v1 static o ID c ID p
func (m *V1StaticOIDCIDP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserClaims(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1StaticOIDCIDP) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("clientID", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *V1StaticOIDCIDP) validateClientScopes(formats strfmt.Registry) error {

	if err := validate.Required("clientScopes", "body", m.ClientScopes); err != nil {
		return err
	}

	return nil
}

func (m *V1StaticOIDCIDP) validateClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("clientSecret", "body", m.ClientSecret); err != nil {
		return err
	}

	return nil
}

func (m *V1StaticOIDCIDP) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

func (m *V1StaticOIDCIDP) validateUserClaims(formats strfmt.Registry) error {

	if err := validate.Required("userClaims", "body", m.UserClaims); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1StaticOIDCIDP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1StaticOIDCIDP) UnmarshalBinary(b []byte) error {
	var res V1StaticOIDCIDP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
