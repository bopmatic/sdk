// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeDeploymentRequest describe deployment request
//
// swagger:model DescribeDeploymentRequest
type DescribeDeploymentRequest struct {

	// a unique identifier associated with this deployment
	ID string `json:"id,omitempty"`
}

// Validate validates this describe deployment request
func (m *DescribeDeploymentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe deployment request based on context it is used
func (m *DescribeDeploymentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeDeploymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDeploymentRequest) UnmarshalBinary(b []byte) error {
	var res DescribeDeploymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}