// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListPackagesRequest list packages request
//
// swagger:model ListPackagesRequest
type ListPackagesRequest struct {

	// leave empty for all projects
	ProjID string `json:"projId,omitempty"`
}

// Validate validates this list packages request
func (m *ListPackagesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list packages request based on context it is used
func (m *ListPackagesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListPackagesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPackagesRequest) UnmarshalBinary(b []byte) error {
	var res ListPackagesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
