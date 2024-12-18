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

// ProjectState - INVALID_PROJ_STATE: unused
//   - INACTIVE: the project has been created but no packages are presently active in any environment
//
// and there are no pending deployments
//   - ACTIVE: the project has at least one package presently active in at least one environment, or
//
// has at least one pending deployment
//   - PROJ_STATE_DELETED: the project has been deleted
//   - UNKNOWN_PROJ_STATE: MAX_INT
//
// swagger:model ProjectState
type ProjectState string

func NewProjectState(value ProjectState) *ProjectState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ProjectState.
func (m ProjectState) Pointer() *ProjectState {
	return &m
}

const (

	// ProjectStateINVALIDPROJSTATE captures enum value "INVALID_PROJ_STATE"
	ProjectStateINVALIDPROJSTATE ProjectState = "INVALID_PROJ_STATE"

	// ProjectStateINACTIVE captures enum value "INACTIVE"
	ProjectStateINACTIVE ProjectState = "INACTIVE"

	// ProjectStateACTIVE captures enum value "ACTIVE"
	ProjectStateACTIVE ProjectState = "ACTIVE"

	// ProjectStatePROJSTATEDELETED captures enum value "PROJ_STATE_DELETED"
	ProjectStatePROJSTATEDELETED ProjectState = "PROJ_STATE_DELETED"

	// ProjectStateUNKNOWNPROJSTATE captures enum value "UNKNOWN_PROJ_STATE"
	ProjectStateUNKNOWNPROJSTATE ProjectState = "UNKNOWN_PROJ_STATE"
)

// for schema
var projectStateEnum []interface{}

func init() {
	var res []ProjectState
	if err := json.Unmarshal([]byte(`["INVALID_PROJ_STATE","INACTIVE","ACTIVE","PROJ_STATE_DELETED","UNKNOWN_PROJ_STATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectStateEnum = append(projectStateEnum, v)
	}
}

func (m ProjectState) validateProjectStateEnum(path, location string, value ProjectState) error {
	if err := validate.EnumCase(path, location, value, projectStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this project state
func (m ProjectState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProjectStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this project state based on context it is used
func (m ProjectState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
