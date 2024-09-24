// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceHeader service header
//
// swagger:model ServiceHeader
type ServiceHeader struct {

	// a project/environment tuple
	ProjEnvHeader *ProjEnvHeader `json:"projEnvHeader,omitempty"`

	// a unique service name within the project
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this service header
func (m *ServiceHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjEnvHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceHeader) validateProjEnvHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjEnvHeader) { // not required
		return nil
	}

	if m.ProjEnvHeader != nil {
		if err := m.ProjEnvHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projEnvHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("projEnvHeader")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service header based on the context it is used
func (m *ServiceHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjEnvHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceHeader) contextValidateProjEnvHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.ProjEnvHeader != nil {

		if swag.IsZero(m.ProjEnvHeader) { // not required
			return nil
		}

		if err := m.ProjEnvHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projEnvHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("projEnvHeader")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceHeader) UnmarshalBinary(b []byte) error {
	var res ServiceHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
