// Code generated by go-swagger; DO NOT EDIT.

package service_runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/bopmatic/sdk/golang/models"
)

// NewDeployPackageParams creates a new DeployPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeployPackageParams() *DeployPackageParams {
	return &DeployPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeployPackageParamsWithTimeout creates a new DeployPackageParams object
// with the ability to set a timeout on a request.
func NewDeployPackageParamsWithTimeout(timeout time.Duration) *DeployPackageParams {
	return &DeployPackageParams{
		timeout: timeout,
	}
}

// NewDeployPackageParamsWithContext creates a new DeployPackageParams object
// with the ability to set a context for a request.
func NewDeployPackageParamsWithContext(ctx context.Context) *DeployPackageParams {
	return &DeployPackageParams{
		Context: ctx,
	}
}

// NewDeployPackageParamsWithHTTPClient creates a new DeployPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeployPackageParamsWithHTTPClient(client *http.Client) *DeployPackageParams {
	return &DeployPackageParams{
		HTTPClient: client,
	}
}

/*
DeployPackageParams contains all the parameters to send to the API endpoint

	for the deploy package operation.

	Typically these are written to a http.Request.
*/
type DeployPackageParams struct {

	// Body.
	Body *models.DeployPackageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deploy package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeployPackageParams) WithDefaults() *DeployPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deploy package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeployPackageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deploy package params
func (o *DeployPackageParams) WithTimeout(timeout time.Duration) *DeployPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deploy package params
func (o *DeployPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deploy package params
func (o *DeployPackageParams) WithContext(ctx context.Context) *DeployPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deploy package params
func (o *DeployPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deploy package params
func (o *DeployPackageParams) WithHTTPClient(client *http.Client) *DeployPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deploy package params
func (o *DeployPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the deploy package params
func (o *DeployPackageParams) WithBody(body *models.DeployPackageRequest) *DeployPackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the deploy package params
func (o *DeployPackageParams) SetBody(body *models.DeployPackageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeployPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
