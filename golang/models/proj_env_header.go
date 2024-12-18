// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjEnvHeader proj env header
//
// swagger:model ProjEnvHeader
type ProjEnvHeader struct {

	// a unique identifier associated with the environment to query; leave blank
	// for all
	EnvID string `json:"envId,omitempty"`

	// a unique identifier associated with the project to query; leave blank
	// for all
	ProjID string `json:"projId,omitempty"`
}

// Validate validates this proj env header
func (m *ProjEnvHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proj env header based on context it is used
func (m *ProjEnvHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjEnvHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjEnvHeader) UnmarshalBinary(b []byte) error {
	var res ProjEnvHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
