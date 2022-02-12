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

// NewListPackagesParams creates a new ListPackagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPackagesParams() *ListPackagesParams {
	return &ListPackagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPackagesParamsWithTimeout creates a new ListPackagesParams object
// with the ability to set a timeout on a request.
func NewListPackagesParamsWithTimeout(timeout time.Duration) *ListPackagesParams {
	return &ListPackagesParams{
		timeout: timeout,
	}
}

// NewListPackagesParamsWithContext creates a new ListPackagesParams object
// with the ability to set a context for a request.
func NewListPackagesParamsWithContext(ctx context.Context) *ListPackagesParams {
	return &ListPackagesParams{
		Context: ctx,
	}
}

// NewListPackagesParamsWithHTTPClient creates a new ListPackagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPackagesParamsWithHTTPClient(client *http.Client) *ListPackagesParams {
	return &ListPackagesParams{
		HTTPClient: client,
	}
}

/* ListPackagesParams contains all the parameters to send to the API endpoint
   for the list packages operation.

   Typically these are written to a http.Request.
*/
type ListPackagesParams struct {

	// Body.
	Body *models.ListPackagesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPackagesParams) WithDefaults() *ListPackagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPackagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list packages params
func (o *ListPackagesParams) WithTimeout(timeout time.Duration) *ListPackagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list packages params
func (o *ListPackagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list packages params
func (o *ListPackagesParams) WithContext(ctx context.Context) *ListPackagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list packages params
func (o *ListPackagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list packages params
func (o *ListPackagesParams) WithHTTPClient(client *http.Client) *ListPackagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list packages params
func (o *ListPackagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list packages params
func (o *ListPackagesParams) WithBody(body *models.ListPackagesRequest) *ListPackagesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list packages params
func (o *ListPackagesParams) SetBody(body *models.ListPackagesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListPackagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
