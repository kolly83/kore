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

// V1alpha1ServicePlanSpec v1alpha1 service plan spec
//
// swagger:model v1alpha1.ServicePlanSpec
type V1alpha1ServicePlanSpec struct {

	// configuration
	// Required: true
	Configuration *string `json:"configuration"`

	// description
	// Required: true
	Description *string `json:"description"`

	// kind
	// Required: true
	Kind *string `json:"kind"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// summary
	// Required: true
	Summary *string `json:"summary"`
}

// Validate validates this v1alpha1 service plan spec
func (m *V1alpha1ServicePlanSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ServicePlanSpec) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1ServicePlanSpec) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1ServicePlanSpec) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1ServicePlanSpec) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ServicePlanSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ServicePlanSpec) UnmarshalBinary(b []byte) error {
	var res V1alpha1ServicePlanSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
