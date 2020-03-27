// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1GKECredentialsStatus v1alpha1 g k e credentials status
//
// swagger:model v1alpha1.GKECredentialsStatus
type V1alpha1GKECredentialsStatus struct {

	// conditions
	Conditions []*V1Condition `json:"conditions"`

	// status
	Status string `json:"status,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this v1alpha1 g k e credentials status
func (m *V1alpha1GKECredentialsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1GKECredentialsStatus) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1GKECredentialsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1GKECredentialsStatus) UnmarshalBinary(b []byte) error {
	var res V1alpha1GKECredentialsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}