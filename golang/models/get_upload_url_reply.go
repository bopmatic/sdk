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

// GetUploadURLReply get upload URL reply
//
// swagger:model GetUploadURLReply
type GetUploadURLReply struct {

	// URL
	URL string `json:"URL,omitempty"`

	// result
	Result *ServiceRunnerResult `json:"result,omitempty"`
}

// Validate validates this get upload URL reply
func (m *GetUploadURLReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUploadURLReply) validateResult(formats strfmt.Registry) error {
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

// ContextValidate validate this get upload URL reply based on the context it is used
func (m *GetUploadURLReply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUploadURLReply) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetUploadURLReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUploadURLReply) UnmarshalBinary(b []byte) error {
	var res GetUploadURLReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
