// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListPackagesReplyListPackagesItem list packages reply list packages item
//
// swagger:model ListPackagesReplyListPackagesItem
type ListPackagesReplyListPackagesItem struct {

	// package Id
	PackageID string `json:"packageId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this list packages reply list packages item
func (m *ListPackagesReplyListPackagesItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list packages reply list packages item based on context it is used
func (m *ListPackagesReplyListPackagesItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListPackagesReplyListPackagesItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPackagesReplyListPackagesItem) UnmarshalBinary(b []byte) error {
	var res ListPackagesReplyListPackagesItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
