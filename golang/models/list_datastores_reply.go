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

// ListDatastoresReply list datastores reply
//
// swagger:model ListDatastoresReply
type ListDatastoresReply struct {

	// a list of service names from the project deployed in the environment
	DatastoreNames []string `json:"datastoreNames"`

	// a project/environment tuple
	Header *ProjEnvHeader `json:"header,omitempty"`

	// result
	Result *ServiceRunnerResult `json:"result,omitempty"`
}

// Validate validates this list datastores reply
func (m *ListDatastoresReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatastoresReply) validateHeader(formats strfmt.Registry) error {
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

func (m *ListDatastoresReply) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list datastores reply based on the context it is used
func (m *ListDatastoresReply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatastoresReply) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ListDatastoresReply) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {

		if swag.IsZero(m.Result) { // not required
			return nil
		}

		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDatastoresReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatastoresReply) UnmarshalBinary(b []byte) error {
	var res ListDatastoresReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
