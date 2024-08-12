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

// GetUploadURLReader is a Reader for the GetUploadURL structure.
type GetUploadURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUploadURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUploadURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUploadURLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUploadURLOK creates a GetUploadURLOK with default headers values
func NewGetUploadURLOK() *GetUploadURLOK {
	return &GetUploadURLOK{}
}

/*
GetUploadURLOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetUploadURLOK struct {
	Payload *models.GetUploadURLReply
}

// IsSuccess returns true when this get upload Url o k response has a 2xx status code
func (o *GetUploadURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upload Url o k response has a 3xx status code
func (o *GetUploadURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upload Url o k response has a 4xx status code
func (o *GetUploadURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upload Url o k response has a 5xx status code
func (o *GetUploadURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upload Url o k response a status code equal to that given
func (o *GetUploadURLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get upload Url o k response
func (o *GetUploadURLOK) Code() int {
	return 200
}

func (o *GetUploadURLOK) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetUploadURL][%d] getUploadUrlOK  %+v", 200, o.Payload)
}

func (o *GetUploadURLOK) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetUploadURL][%d] getUploadUrlOK  %+v", 200, o.Payload)
}

func (o *GetUploadURLOK) GetPayload() *models.GetUploadURLReply {
	return o.Payload
}

func (o *GetUploadURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUploadURLReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUploadURLDefault creates a GetUploadURLDefault with default headers values
func NewGetUploadURLDefault(code int) *GetUploadURLDefault {
	return &GetUploadURLDefault{
		_statusCode: code,
	}
}

/*
GetUploadURLDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetUploadURLDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get upload URL default response has a 2xx status code
func (o *GetUploadURLDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get upload URL default response has a 3xx status code
func (o *GetUploadURLDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get upload URL default response has a 4xx status code
func (o *GetUploadURLDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get upload URL default response has a 5xx status code
func (o *GetUploadURLDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get upload URL default response a status code equal to that given
func (o *GetUploadURLDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get upload URL default response
func (o *GetUploadURLDefault) Code() int {
	return o._statusCode
}

func (o *GetUploadURLDefault) Error() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetUploadURL][%d] GetUploadURL default  %+v", o._statusCode, o.Payload)
}

func (o *GetUploadURLDefault) String() string {
	return fmt.Sprintf("[POST /ServiceRunner/GetUploadURL][%d] GetUploadURL default  %+v", o._statusCode, o.Payload)
}

func (o *GetUploadURLDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetUploadURLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
