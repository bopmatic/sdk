// Code generated by go-swagger; DO NOT EDIT.

package service_runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bopmatic/sdk/golang/models"
)

// ListMetricsReader is a Reader for the ListMetrics structure.
type ListMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMetricsOK creates a ListMetricsOK with default headers values
func NewListMetricsOK() *ListMetricsOK {
	return &ListMetricsOK{}
}

/*
ListMetricsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListMetricsOK struct {
	Payload *models.ListMetricsReply
}

// IsSuccess returns true when this list metrics o k response has a 2xx status code
func (o *ListMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list metrics o k response has a 3xx status code
func (o *ListMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list metrics o k response has a 4xx status code
func (o *ListMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list metrics o k response has a 5xx status code
func (o *ListMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list metrics o k response a status code equal to that given
func (o *ListMetricsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list metrics o k response
func (o *ListMetricsOK) Code() int {
	return 200
}

func (o *ListMetricsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/ListMetrics][%d] listMetricsOK %s", 200, payload)
}

func (o *ListMetricsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/ListMetrics][%d] listMetricsOK %s", 200, payload)
}

func (o *ListMetricsOK) GetPayload() *models.ListMetricsReply {
	return o.Payload
}

func (o *ListMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListMetricsReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMetricsDefault creates a ListMetricsDefault with default headers values
func NewListMetricsDefault(code int) *ListMetricsDefault {
	return &ListMetricsDefault{
		_statusCode: code,
	}
}

/*
ListMetricsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListMetricsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this list metrics default response has a 2xx status code
func (o *ListMetricsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list metrics default response has a 3xx status code
func (o *ListMetricsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list metrics default response has a 4xx status code
func (o *ListMetricsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list metrics default response has a 5xx status code
func (o *ListMetricsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list metrics default response a status code equal to that given
func (o *ListMetricsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list metrics default response
func (o *ListMetricsDefault) Code() int {
	return o._statusCode
}

func (o *ListMetricsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/ListMetrics][%d] ListMetrics default %s", o._statusCode, payload)
}

func (o *ListMetricsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/ListMetrics][%d] ListMetrics default %s", o._statusCode, payload)
}

func (o *ListMetricsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ListMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
