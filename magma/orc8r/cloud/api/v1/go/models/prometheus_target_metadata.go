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

// PrometheusTargetMetadata prometheus target metadata
// swagger:model prometheus_target_metadata
type PrometheusTargetMetadata struct {

	// instance
	// Required: true
	Instance *string `json:"instance"`

	// job
	// Required: true
	Job *string `json:"job"`
}

// Validate validates this prometheus target metadata
func (m *PrometheusTargetMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrometheusTargetMetadata) validateInstance(formats strfmt.Registry) error {

	if err := validate.Required("instance", "body", m.Instance); err != nil {
		return err
	}

	return nil
}

func (m *PrometheusTargetMetadata) validateJob(formats strfmt.Registry) error {

	if err := validate.Required("job", "body", m.Job); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusTargetMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusTargetMetadata) UnmarshalBinary(b []byte) error {
	var res PrometheusTargetMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
