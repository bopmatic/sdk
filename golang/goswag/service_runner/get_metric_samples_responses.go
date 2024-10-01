// Code generated by go-swagger; DO NOT EDIT.

package service_runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bopmatic/sdk/golang/models"
)

// GetMetricSamplesReader is a Reader for the GetMetricSamples structure.
type GetMetricSamplesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricSamplesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricSamplesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMetricSamplesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricSamplesOK creates a GetMetricSamplesOK with default headers values
func NewGetMetricSamplesOK() *GetMetricSamplesOK {
	return &GetMetricSamplesOK{}
}

/*
GetMetricSamplesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetMetricSamplesOK struct {
	Payload *models.GetMetricSamplesReply
}

// IsSuccess returns true when this get metric samples o k response has a 2xx status code
func (o *GetMetricSamplesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metric samples o k response has a 3xx status code
func (o *GetMetricSamplesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metric samples o k response has a 4xx status code
func (o *GetMetricSamplesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metric samples o k response has a 5xx status code
func (o *GetMetricSamplesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metric samples o k response a status code equal to that given
func (o *GetMetricSamplesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get metric samples o k response
func (o *GetMetricSamplesOK) Code() int {
	return 200
}

func (o *GetMetricSamplesOK) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetMetricSamples][%d] getMetricSamplesOK  %+v", 200, o.Payload)
}

func (o *GetMetricSamplesOK) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetMetricSamples][%d] getMetricSamplesOK  %+v", 200, o.Payload)
}

func (o *GetMetricSamplesOK) GetPayload() *models.GetMetricSamplesReply {
	return o.Payload
}

func (o *GetMetricSamplesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMetricSamplesReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricSamplesDefault creates a GetMetricSamplesDefault with default headers values
func NewGetMetricSamplesDefault(code int) *GetMetricSamplesDefault {
	return &GetMetricSamplesDefault{
		_statusCode: code,
	}
}

/*
GetMetricSamplesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetMetricSamplesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get metric samples default response has a 2xx status code
func (o *GetMetricSamplesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get metric samples default response has a 3xx status code
func (o *GetMetricSamplesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get metric samples default response has a 4xx status code
func (o *GetMetricSamplesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get metric samples default response has a 5xx status code
func (o *GetMetricSamplesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get metric samples default response a status code equal to that given
func (o *GetMetricSamplesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get metric samples default response
func (o *GetMetricSamplesDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricSamplesDefault) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetMetricSamples][%d] GetMetricSamples default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetricSamplesDefault) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetMetricSamples][%d] GetMetricSamples default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetricSamplesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetMetricSamplesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}