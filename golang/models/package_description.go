// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageDescription package description
//
// swagger:model PackageDescription
type PackageDescription struct {

	// hex string of first 4 bytes of packageXsum
	PackageID string `json:"packageId,omitempty"`

	// package tarball content in .tar.xz format (limited to 6MiB);
	// Format: byte
	PackageTarballData strfmt.Base64 `json:"packageTarballData,omitempty"`

	// URL to package tarball (when larger than 6MiB)
	PackageTarballURL string `json:"packageTarballURL,omitempty"`

	// sha256 checksum of packageTarballData
	// Format: byte
	PackageXsum strfmt.Base64 `json:"packageXsum,omitempty"`

	// name of the Bopmatic project; must be unique
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this package description
func (m *PackageDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this package description based on context it is used
func (m *PackageDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDescription) UnmarshalBinary(b []byte) error {
	var res PackageDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
