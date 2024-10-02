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

// DeactivateProjectRequest deactivate project request
//
// swagger:model DeactivateProjectRequest
type DeactivateProjectRequest struct {

	// the project & environment id to deactivate
	ProjEnvHeader *ProjEnvHeader `json:"projEnvHeader,omitempty"`
}

// Validate validates this deactivate project request
func (m *DeactivateProjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjEnvHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeactivateProjectRequest) validateProjEnvHeader(formats strfmt.Registry) error {
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

// ContextValidate validate this deactivate project request based on the context it is used
func (m *DeactivateProjectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjEnvHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeactivateProjectRequest) contextValidateProjEnvHeader(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DeactivateProjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeactivateProjectRequest) UnmarshalBinary(b []byte) error {
	var res DeactivateProjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
