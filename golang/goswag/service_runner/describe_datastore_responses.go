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

// DescribeDatastoreReader is a Reader for the DescribeDatastore structure.
type DescribeDatastoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDatastoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeDatastoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeDatastoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeDatastoreOK creates a DescribeDatastoreOK with default headers values
func NewDescribeDatastoreOK() *DescribeDatastoreOK {
	return &DescribeDatastoreOK{}
}

/*
DescribeDatastoreOK describes a response with status code 200, with default header values.

A successful response.
*/
type DescribeDatastoreOK struct {
	Payload *models.DescribeDatastoreReply
}

// IsSuccess returns true when this describe datastore o k response has a 2xx status code
func (o *DescribeDatastoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe datastore o k response has a 3xx status code
func (o *DescribeDatastoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe datastore o k response has a 4xx status code
func (o *DescribeDatastoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe datastore o k response has a 5xx status code
func (o *DescribeDatastoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe datastore o k response a status code equal to that given
func (o *DescribeDatastoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe datastore o k response
func (o *DescribeDatastoreOK) Code() int {
	return 200
}

func (o *DescribeDatastoreOK) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribeDatastore][%d] describeDatastoreOK  %+v", 200, o.Payload)
}

func (o *DescribeDatastoreOK) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribeDatastore][%d] describeDatastoreOK  %+v", 200, o.Payload)
}

func (o *DescribeDatastoreOK) GetPayload() *models.DescribeDatastoreReply {
	return o.Payload
}

func (o *DescribeDatastoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeDatastoreReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDatastoreDefault creates a DescribeDatastoreDefault with default headers values
func NewDescribeDatastoreDefault(code int) *DescribeDatastoreDefault {
	return &DescribeDatastoreDefault{
		_statusCode: code,
	}
}

/*
DescribeDatastoreDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DescribeDatastoreDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this describe datastore default response has a 2xx status code
func (o *DescribeDatastoreDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe datastore default response has a 3xx status code
func (o *DescribeDatastoreDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe datastore default response has a 4xx status code
func (o *DescribeDatastoreDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe datastore default response has a 5xx status code
func (o *DescribeDatastoreDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe datastore default response a status code equal to that given
func (o *DescribeDatastoreDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe datastore default response
func (o *DescribeDatastoreDefault) Code() int {
	return o._statusCode
}

func (o *DescribeDatastoreDefault) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribeDatastore][%d] DescribeDatastore default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeDatastoreDefault) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribeDatastore][%d] DescribeDatastore default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeDatastoreDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DescribeDatastoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
