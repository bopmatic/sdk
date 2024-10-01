// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ServiceRunnerStatus - STATUS_UNKNOWN: MAX_INT
//
// swagger:model ServiceRunnerStatus
type ServiceRunnerStatus string

func NewServiceRunnerStatus(value ServiceRunnerStatus) *ServiceRunnerStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ServiceRunnerStatus.
func (m ServiceRunnerStatus) Pointer() *ServiceRunnerStatus {
	return &m
}

const (

	// ServiceRunnerStatusSTATUSOK captures enum value "STATUS_OK"
	ServiceRunnerStatusSTATUSOK ServiceRunnerStatus = "STATUS_OK"

	// ServiceRunnerStatusSTATUSNOTFOUND captures enum value "STATUS_NOT_FOUND"
	ServiceRunnerStatusSTATUSNOTFOUND ServiceRunnerStatus = "STATUS_NOT_FOUND"

	// ServiceRunnerStatusSTATUSDNSCONFLICT captures enum value "STATUS_DNS_CONFLICT"
	ServiceRunnerStatusSTATUSDNSCONFLICT ServiceRunnerStatus = "STATUS_DNS_CONFLICT"

	// ServiceRunnerStatusSTATUSACCTLIMITREACHED captures enum value "STATUS_ACCT_LIMIT_REACHED"
	ServiceRunnerStatusSTATUSACCTLIMITREACHED ServiceRunnerStatus = "STATUS_ACCT_LIMIT_REACHED"

	// ServiceRunnerStatusSTATUSEXISTS captures enum value "STATUS_EXISTS"
	ServiceRunnerStatusSTATUSEXISTS ServiceRunnerStatus = "STATUS_EXISTS"

	// ServiceRunnerStatusSTATUSNOTEXISTS captures enum value "STATUS_NOT_EXISTS"
	ServiceRunnerStatusSTATUSNOTEXISTS ServiceRunnerStatus = "STATUS_NOT_EXISTS"

	// ServiceRunnerStatusSTATUSUNKNOWN captures enum value "STATUS_UNKNOWN"
	ServiceRunnerStatusSTATUSUNKNOWN ServiceRunnerStatus = "STATUS_UNKNOWN"
)

// for schema
var serviceRunnerStatusEnum []interface{}

func init() {
	var res []ServiceRunnerStatus
	if err := json.Unmarshal([]byte(`["STATUS_OK","STATUS_NOT_FOUND","STATUS_DNS_CONFLICT","STATUS_ACCT_LIMIT_REACHED","STATUS_EXISTS","STATUS_NOT_EXISTS","STATUS_UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceRunnerStatusEnum = append(serviceRunnerStatusEnum, v)
	}
}

func (m ServiceRunnerStatus) validateServiceRunnerStatusEnum(path, location string, value ServiceRunnerStatus) error {
	if err := validate.EnumCase(path, location, value, serviceRunnerStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this service runner status
func (m ServiceRunnerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceRunnerStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this service runner status based on context it is used
func (m ServiceRunnerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}