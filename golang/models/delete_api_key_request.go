// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteAPIKeyRequest delete Api key request
//
// swagger:model DeleteApiKeyRequest
type DeleteAPIKeyRequest struct {

	// the id of the api key to delete
	KeyID string `json:"keyId,omitempty"`
}

// Validate validates this delete Api key request
func (m *DeleteAPIKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete Api key request based on context it is used
func (m *DeleteAPIKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteAPIKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteAPIKeyRequest) UnmarshalBinary(b []byte) error {
	var res DeleteAPIKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}