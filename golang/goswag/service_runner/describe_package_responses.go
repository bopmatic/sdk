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

// DescribePackageReader is a Reader for the DescribePackage structure.
type DescribePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribePackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribePackageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribePackageOK creates a DescribePackageOK with default headers values
func NewDescribePackageOK() *DescribePackageOK {
	return &DescribePackageOK{}
}

/*
DescribePackageOK describes a response with status code 200, with default header values.

A successful response.
*/
type DescribePackageOK struct {
	Payload *models.DescribePackageReply
}

// IsSuccess returns true when this describe package o k response has a 2xx status code
func (o *DescribePackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe package o k response has a 3xx status code
func (o *DescribePackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe package o k response has a 4xx status code
func (o *DescribePackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe package o k response has a 5xx status code
func (o *DescribePackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe package o k response a status code equal to that given
func (o *DescribePackageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe package o k response
func (o *DescribePackageOK) Code() int {
	return 200
}

func (o *DescribePackageOK) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribePackage][%d] describePackageOK  %+v", 200, o.Payload)
}

func (o *DescribePackageOK) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribePackage][%d] describePackageOK  %+v", 200, o.Payload)
}

func (o *DescribePackageOK) GetPayload() *models.DescribePackageReply {
	return o.Payload
}

func (o *DescribePackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribePackageReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribePackageDefault creates a DescribePackageDefault with default headers values
func NewDescribePackageDefault(code int) *DescribePackageDefault {
	return &DescribePackageDefault{
		_statusCode: code,
	}
}

/*
DescribePackageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DescribePackageDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this describe package default response has a 2xx status code
func (o *DescribePackageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe package default response has a 3xx status code
func (o *DescribePackageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe package default response has a 4xx status code
func (o *DescribePackageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe package default response has a 5xx status code
func (o *DescribePackageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe package default response a status code equal to that given
func (o *DescribePackageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe package default response
func (o *DescribePackageDefault) Code() int {
	return o._statusCode
}

func (o *DescribePackageDefault) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribePackage][%d] DescribePackage default  %+v", o._statusCode, o.Payload)
}

func (o *DescribePackageDefault) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/DescribePackage][%d] DescribePackage default  %+v", o._statusCode, o.Payload)
}

func (o *DescribePackageDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DescribePackageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
