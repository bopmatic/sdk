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

// ServiceRunnerStatus - STATUS_INVALID: unused
//   - STATUS_UNKNOWN: MAX_INT
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

	// ServiceRunnerStatusSTATUSINVALID captures enum value "STATUS_INVALID"
	ServiceRunnerStatusSTATUSINVALID ServiceRunnerStatus = "STATUS_INVALID"

	// ServiceRunnerStatusSTATUSOK captures enum value "STATUS_OK"
	ServiceRunnerStatusSTATUSOK ServiceRunnerStatus = "STATUS_OK"

	// ServiceRunnerStatusSTATUSINTERNALERR captures enum value "STATUS_INTERNAL_ERR"
	ServiceRunnerStatusSTATUSINTERNALERR ServiceRunnerStatus = "STATUS_INTERNAL_ERR"

	// ServiceRunnerStatusSTATUSDNSCONFLICT captures enum value "STATUS_DNS_CONFLICT"
	ServiceRunnerStatusSTATUSDNSCONFLICT ServiceRunnerStatus = "STATUS_DNS_CONFLICT"

	// ServiceRunnerStatusSTATUSACCTLIMITREACHED captures enum value "STATUS_ACCT_LIMIT_REACHED"
	ServiceRunnerStatusSTATUSACCTLIMITREACHED ServiceRunnerStatus = "STATUS_ACCT_LIMIT_REACHED"

	// ServiceRunnerStatusSTATUSEXISTS captures enum value "STATUS_EXISTS"
	ServiceRunnerStatusSTATUSEXISTS ServiceRunnerStatus = "STATUS_EXISTS"

	// ServiceRunnerStatusSTATUSNOTEXISTS captures enum value "STATUS_NOT_EXISTS"
	ServiceRunnerStatusSTATUSNOTEXISTS ServiceRunnerStatus = "STATUS_NOT_EXISTS"

	// ServiceRunnerStatusSTATUSINVALIDREQUEST captures enum value "STATUS_INVALID_REQUEST"
	ServiceRunnerStatusSTATUSINVALIDREQUEST ServiceRunnerStatus = "STATUS_INVALID_REQUEST"

	// ServiceRunnerStatusSTATUSUNKNOWN captures enum value "STATUS_UNKNOWN"
	ServiceRunnerStatusSTATUSUNKNOWN ServiceRunnerStatus = "STATUS_UNKNOWN"
)

// for schema
var serviceRunnerStatusEnum []interface{}

func init() {
	var res []ServiceRunnerStatus
	if err := json.Unmarshal([]byte(`["STATUS_INVALID","STATUS_OK","STATUS_INTERNAL_ERR","STATUS_DNS_CONFLICT","STATUS_ACCT_LIMIT_REACHED","STATUS_EXISTS","STATUS_NOT_EXISTS","STATUS_INVALID_REQUEST","STATUS_UNKNOWN"]`), &res); err != nil {
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
