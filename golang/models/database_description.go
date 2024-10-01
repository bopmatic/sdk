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

// DatabaseDescription database description
//
// swagger:model DatabaseDescription
type DatabaseDescription struct {

	// service header
	DatabaseHeader *DatabaseHeader `json:"databaseHeader,omitempty"`

	// a list of service names from the project deployed in the environment
	TableNames []string `json:"tableNames"`
}

// Validate validates this database description
func (m *DatabaseDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseDescription) validateDatabaseHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseHeader) { // not required
		return nil
	}

	if m.DatabaseHeader != nil {
		if err := m.DatabaseHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseHeader")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this database description based on the context it is used
func (m *DatabaseDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatabaseHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseDescription) contextValidateDatabaseHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.DatabaseHeader != nil {

		if swag.IsZero(m.DatabaseHeader) { // not required
			return nil
		}

		if err := m.DatabaseHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseHeader")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseDescription) UnmarshalBinary(b []byte) error {
	var res DatabaseDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}