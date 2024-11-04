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

// DescribeAPIKeyReader is a Reader for the DescribeAPIKey structure.
type DescribeAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeAPIKeyOK creates a DescribeAPIKeyOK with default headers values
func NewDescribeAPIKeyOK() *DescribeAPIKeyOK {
	return &DescribeAPIKeyOK{}
}

/*
DescribeAPIKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type DescribeAPIKeyOK struct {
	Payload *models.DescribeAPIKeyReply
}

// IsSuccess returns true when this describe Api key o k response has a 2xx status code
func (o *DescribeAPIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe Api key o k response has a 3xx status code
func (o *DescribeAPIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe Api key o k response has a 4xx status code
func (o *DescribeAPIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe Api key o k response has a 5xx status code
func (o *DescribeAPIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe Api key o k response a status code equal to that given
func (o *DescribeAPIKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe Api key o k response
func (o *DescribeAPIKeyOK) Code() int {
	return 200
}

func (o *DescribeAPIKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeApiKey][%d] describeApiKeyOK %s", 200, payload)
}

func (o *DescribeAPIKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeApiKey][%d] describeApiKeyOK %s", 200, payload)
}

func (o *DescribeAPIKeyOK) GetPayload() *models.DescribeAPIKeyReply {
	return o.Payload
}

func (o *DescribeAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeAPIKeyReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeAPIKeyDefault creates a DescribeAPIKeyDefault with default headers values
func NewDescribeAPIKeyDefault(code int) *DescribeAPIKeyDefault {
	return &DescribeAPIKeyDefault{
		_statusCode: code,
	}
}

/*
DescribeAPIKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DescribeAPIKeyDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this describe Api key default response has a 2xx status code
func (o *DescribeAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe Api key default response has a 3xx status code
func (o *DescribeAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe Api key default response has a 4xx status code
func (o *DescribeAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe Api key default response has a 5xx status code
func (o *DescribeAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe Api key default response a status code equal to that given
func (o *DescribeAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe Api key default response
func (o *DescribeAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *DescribeAPIKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeApiKey][%d] DescribeApiKey default %s", o._statusCode, payload)
}

func (o *DescribeAPIKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeApiKey][%d] DescribeApiKey default %s", o._statusCode, payload)
}

func (o *DescribeAPIKeyDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DescribeAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
