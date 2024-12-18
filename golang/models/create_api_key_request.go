// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAPIKeyRequest create Api key request
//
// swagger:model CreateApiKeyRequest
type CreateAPIKeyRequest struct {

	// a description of the API key
	Description string `json:"description,omitempty"`

	// time the key should expire expressed as the number of milliseconds since
	// Jan 1, 1970 00:00:00 UTC. A value of 0 indicates the key should never expire.
	ExpireTime string `json:"expireTime,omitempty"`

	// the name of the API key
	Name string `json:"name,omitempty"`
}

// Validate validates this create Api key request
func (m *CreateAPIKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create Api key request based on context it is used
func (m *CreateAPIKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAPIKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAPIKeyRequest) UnmarshalBinary(b []byte) error {
	var res CreateAPIKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
