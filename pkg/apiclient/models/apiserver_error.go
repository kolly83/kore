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

// ApiserverError apiserver error
//
// swagger:model apiserver.Error
type ApiserverError struct {

	// code
	// Required: true
	Code *int32 `json:"code"`

	// detail
	// Required: true
	Detail *string `json:"detail"`

	// message
	// Required: true
	Message *string `json:"message"`

	// uri
	// Required: true
	URI *string `json:"uri"`

	// verb
	// Required: true
	Verb *string `json:"verb"`
}

// Validate validates this apiserver error
func (m *ApiserverError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApiserverError) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *ApiserverError) validateDetail(formats strfmt.Registry) error {

	if err := validate.Required("detail", "body", m.Detail); err != nil {
		return err
	}

	return nil
}

func (m *ApiserverError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *ApiserverError) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("uri", "body", m.URI); err != nil {
		return err
	}

	return nil
}

func (m *ApiserverError) validateVerb(formats strfmt.Registry) error {

	if err := validate.Required("verb", "body", m.Verb); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApiserverError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApiserverError) UnmarshalBinary(b []byte) error {
	var res ApiserverError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
