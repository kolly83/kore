// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ServiceSpec v1 service spec
//
// swagger:model v1.ServiceSpec
type V1ServiceSpec struct {

	// cluster
	Cluster *V1Ownership `json:"cluster,omitempty"`

	// cluster namespace
	ClusterNamespace string `json:"clusterNamespace,omitempty"`

	// configuration
	Configuration interface{} `json:"configuration,omitempty"`

	// configuration from
	ConfigurationFrom []*V1ConfigurationFromSource `json:"configurationFrom"`

	// kind
	// Required: true
	Kind *string `json:"kind"`

	// plan
	// Required: true
	Plan *string `json:"plan"`
}

// Validate validates this v1 service spec
func (m *V1ServiceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServiceSpec) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *V1ServiceSpec) validateConfigurationFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigurationFrom) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigurationFrom); i++ {
		if swag.IsZero(m.ConfigurationFrom[i]) { // not required
			continue
		}

		if m.ConfigurationFrom[i] != nil {
			if err := m.ConfigurationFrom[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurationFrom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ServiceSpec) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *V1ServiceSpec) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("plan", "body", m.Plan); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ServiceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServiceSpec) UnmarshalBinary(b []byte) error {
	var res V1ServiceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
