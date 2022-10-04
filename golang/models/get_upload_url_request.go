// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetUploadURLRequest get upload URL request
//
// swagger:model GetUploadURLRequest
type GetUploadURLRequest struct {

	// key
	Key string `json:"key,omitempty"`
}

// Validate validates this get upload URL request
func (m *GetUploadURLRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get upload URL request based on context it is used
func (m *GetUploadURLRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetUploadURLRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUploadURLRequest) UnmarshalBinary(b []byte) error {
	var res GetUploadURLRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}