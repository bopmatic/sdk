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

// ListDatastoresRequest list datastores request
//
// swagger:model ListDatastoresRequest
type ListDatastoresRequest struct {

	// a project/environment tuple
	Header *ProjEnvHeader `json:"header,omitempty"`
}

// Validate validates this list datastores request
func (m *ListDatastoresRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatastoresRequest) validateHeader(formats strfmt.Registry) error {
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

// ContextValidate validate this list datastores request based on the context it is used
func (m *ListDatastoresRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatastoresRequest) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ListDatastoresRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatastoresRequest) UnmarshalBinary(b []byte) error {
	var res ListDatastoresRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}