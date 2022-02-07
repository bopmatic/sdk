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

// NewDescribePackageParams creates a new DescribePackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribePackageParams() *DescribePackageParams {
	return &DescribePackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribePackageParamsWithTimeout creates a new DescribePackageParams object
// with the ability to set a timeout on a request.
func NewDescribePackageParamsWithTimeout(timeout time.Duration) *DescribePackageParams {
	return &DescribePackageParams{
		timeout: timeout,
	}
}

// NewDescribePackageParamsWithContext creates a new DescribePackageParams object
// with the ability to set a context for a request.
func NewDescribePackageParamsWithContext(ctx context.Context) *DescribePackageParams {
	return &DescribePackageParams{
		Context: ctx,
	}
}

// NewDescribePackageParamsWithHTTPClient creates a new DescribePackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribePackageParamsWithHTTPClient(client *http.Client) *DescribePackageParams {
	return &DescribePackageParams{
		HTTPClient: client,
	}
}

/* DescribePackageParams contains all the parameters to send to the API endpoint
   for the describe package operation.

   Typically these are written to a http.Request.
*/
type DescribePackageParams struct {

	// Body.
	Body *models.DescribePackageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribePackageParams) WithDefaults() *DescribePackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribePackageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe package params
func (o *DescribePackageParams) WithTimeout(timeout time.Duration) *DescribePackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe package params
func (o *DescribePackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe package params
func (o *DescribePackageParams) WithContext(ctx context.Context) *DescribePackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe package params
func (o *DescribePackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe package params
func (o *DescribePackageParams) WithHTTPClient(client *http.Client) *DescribePackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe package params
func (o *DescribePackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the describe package params
func (o *DescribePackageParams) WithBody(body *models.DescribePackageRequest) *DescribePackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the describe package params
func (o *DescribePackageParams) SetBody(body *models.DescribePackageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DescribePackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
