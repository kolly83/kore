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

// V1alpha1GKECredentialsSpec v1alpha1 g k e credentials spec
//
// swagger:model v1alpha1.GKECredentialsSpec
type V1alpha1GKECredentialsSpec struct {

	// account
	// Required: true
	Account *string `json:"account"`

	// project
	// Required: true
	Project *string `json:"project"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this v1alpha1 g k e credentials spec
func (m *V1alpha1GKECredentialsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1GKECredentialsSpec) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1GKECredentialsSpec) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1GKECredentialsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1GKECredentialsSpec) UnmarshalBinary(b []byte) error {
	var res V1alpha1GKECredentialsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}