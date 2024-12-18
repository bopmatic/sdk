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

// DescribeServiceReader is a Reader for the DescribeService structure.
type DescribeServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeServiceOK creates a DescribeServiceOK with default headers values
func NewDescribeServiceOK() *DescribeServiceOK {
	return &DescribeServiceOK{}
}

/*
DescribeServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DescribeServiceOK struct {
	Payload *models.DescribeServiceReply
}

// IsSuccess returns true when this describe service o k response has a 2xx status code
func (o *DescribeServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe service o k response has a 3xx status code
func (o *DescribeServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe service o k response has a 4xx status code
func (o *DescribeServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe service o k response has a 5xx status code
func (o *DescribeServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe service o k response a status code equal to that given
func (o *DescribeServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe service o k response
func (o *DescribeServiceOK) Code() int {
	return 200
}

func (o *DescribeServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeService][%d] describeServiceOK %s", 200, payload)
}

func (o *DescribeServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeService][%d] describeServiceOK %s", 200, payload)
}

func (o *DescribeServiceOK) GetPayload() *models.DescribeServiceReply {
	return o.Payload
}

func (o *DescribeServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeServiceReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServiceDefault creates a DescribeServiceDefault with default headers values
func NewDescribeServiceDefault(code int) *DescribeServiceDefault {
	return &DescribeServiceDefault{
		_statusCode: code,
	}
}

/*
DescribeServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DescribeServiceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this describe service default response has a 2xx status code
func (o *DescribeServiceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe service default response has a 3xx status code
func (o *DescribeServiceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe service default response has a 4xx status code
func (o *DescribeServiceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe service default response has a 5xx status code
func (o *DescribeServiceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe service default response a status code equal to that given
func (o *DescribeServiceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe service default response
func (o *DescribeServiceDefault) Code() int {
	return o._statusCode
}

func (o *DescribeServiceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeService][%d] DescribeService default %s", o._statusCode, payload)
}

func (o *DescribeServiceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ServiceRunner/DescribeService][%d] DescribeService default %s", o._statusCode, payload)
}

func (o *DescribeServiceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DescribeServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
