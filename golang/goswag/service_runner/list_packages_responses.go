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

// ListPackagesReader is a Reader for the ListPackages structure.
type ListPackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPackagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPackagesOK creates a ListPackagesOK with default headers values
func NewListPackagesOK() *ListPackagesOK {
	return &ListPackagesOK{}
}

/* ListPackagesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListPackagesOK struct {
	Payload *models.ListPackagesReply
}

// IsSuccess returns true when this list packages o k response has a 2xx status code
func (o *ListPackagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list packages o k response has a 3xx status code
func (o *ListPackagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list packages o k response has a 4xx status code
func (o *ListPackagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list packages o k response has a 5xx status code
func (o *ListPackagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list packages o k response a status code equal to that given
func (o *ListPackagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListPackagesOK) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/ListPackages][%d] listPackagesOK  %+v", 200, o.Payload)
}

func (o *ListPackagesOK) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/ListPackages][%d] listPackagesOK  %+v", 200, o.Payload)
}

func (o *ListPackagesOK) GetPayload() *models.ListPackagesReply {
	return o.Payload
}

func (o *ListPackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListPackagesReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPackagesDefault creates a ListPackagesDefault with default headers values
func NewListPackagesDefault(code int) *ListPackagesDefault {
	return &ListPackagesDefault{
		_statusCode: code,
	}
}

/* ListPackagesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListPackagesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the list packages default response
func (o *ListPackagesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list packages default response has a 2xx status code
func (o *ListPackagesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list packages default response has a 3xx status code
func (o *ListPackagesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list packages default response has a 4xx status code
func (o *ListPackagesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list packages default response has a 5xx status code
func (o *ListPackagesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list packages default response a status code equal to that given
func (o *ListPackagesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListPackagesDefault) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/ListPackages][%d] ListPackages default  %+v", o._statusCode, o.Payload)
}

func (o *ListPackagesDefault) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/ListPackages][%d] ListPackages default  %+v", o._statusCode, o.Payload)
}

func (o *ListPackagesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ListPackagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
