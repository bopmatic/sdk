// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetLogsRequest get logs request
//
// swagger:model GetLogsRequest
type GetLogsRequest struct {

	// of seconds since Jan 1, 1970 00:00:00 UTC
	//
	// latest log message to retrieve expressed as the number of
	EndTime string `json:"endTime,omitempty"`

	// name of the Bopmatic project
	ProjectName string `json:"projectName,omitempty"`

	// name of an RPC defined within the service
	RPCName string `json:"rpcName,omitempty"`

	// name of a service defined within the Bopmatic project
	ServiceName string `json:"serviceName,omitempty"`

	// earliest log message to retrieve expressed as the number
	StartTime string `json:"startTime,omitempty"`
}

// Validate validates this get logs request
func (m *GetLogsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get logs request based on context it is used
func (m *GetLogsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetLogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogsRequest) UnmarshalBinary(b []byte) error {
	var res GetLogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
