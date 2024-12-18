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

// CreateEnvironmentRequest create environment request
//
// swagger:model CreateEnvironmentRequest
type CreateEnvironmentRequest struct {

	// the environment header
	Header *EnvironmentHeader `json:"header,omitempty"`
}

// Validate validates this create environment request
func (m *CreateEnvironmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateEnvironmentRequest) validateHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create environment request based on the context it is used
func (m *CreateEnvironmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateEnvironmentRequest) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.Header != nil {

		if swag.IsZero(m.Header) { // not required
			return nil
		}

		if err := m.Header.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateEnvironmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEnvironmentRequest) UnmarshalBinary(b []byte) error {
	var res CreateEnvironmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
