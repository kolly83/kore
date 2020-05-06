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

// V1beta1AccountsRule v1beta1 accounts rule
//
// swagger:model v1beta1.AccountsRule
type V1beta1AccountsRule struct {

	// description
	Description string `json:"description,omitempty"`

	// exact
	Exact string `json:"exact,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// plans
	// Required: true
	Plans []string `json:"plans"`

	// prefix
	Prefix string `json:"prefix,omitempty"`

	// suffix
	Suffix string `json:"suffix,omitempty"`
}

// Validate validates this v1beta1 accounts rule
func (m *V1beta1AccountsRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1AccountsRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1AccountsRule) validatePlans(formats strfmt.Registry) error {

	if err := validate.Required("plans", "body", m.Plans); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1AccountsRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1AccountsRule) UnmarshalBinary(b []byte) error {
	var res V1beta1AccountsRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
