// Code generated by go-swagger; DO NOT EDIT.

package sign_up

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

// RequestAccessReader is a Reader for the RequestAccess structure.
type RequestAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRequestAccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRequestAccessOK creates a RequestAccessOK with default headers values
func NewRequestAccessOK() *RequestAccessOK {
	return &RequestAccessOK{}
}

/*
RequestAccessOK describes a response with status code 200, with default header values.

A successful response.
*/
type RequestAccessOK struct {
	Payload models.RequestAccessReply
}

// IsSuccess returns true when this request access o k response has a 2xx status code
func (o *RequestAccessOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this request access o k response has a 3xx status code
func (o *RequestAccessOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request access o k response has a 4xx status code
func (o *RequestAccessOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this request access o k response has a 5xx status code
func (o *RequestAccessOK) IsServerError() bool {
	return false
}

// IsCode returns true when this request access o k response a status code equal to that given
func (o *RequestAccessOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the request access o k response
func (o *RequestAccessOK) Code() int {
	return 200
}

func (o *RequestAccessOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SignUp/RequestAccess][%d] requestAccessOK %s", 200, payload)
}

func (o *RequestAccessOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SignUp/RequestAccess][%d] requestAccessOK %s", 200, payload)
}

func (o *RequestAccessOK) GetPayload() models.RequestAccessReply {
	return o.Payload
}

func (o *RequestAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestAccessDefault creates a RequestAccessDefault with default headers values
func NewRequestAccessDefault(code int) *RequestAccessDefault {
	return &RequestAccessDefault{
		_statusCode: code,
	}
}

/*
RequestAccessDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RequestAccessDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this request access default response has a 2xx status code
func (o *RequestAccessDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this request access default response has a 3xx status code
func (o *RequestAccessDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this request access default response has a 4xx status code
func (o *RequestAccessDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this request access default response has a 5xx status code
func (o *RequestAccessDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this request access default response a status code equal to that given
func (o *RequestAccessDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the request access default response
func (o *RequestAccessDefault) Code() int {
	return o._statusCode
}

func (o *RequestAccessDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SignUp/RequestAccess][%d] RequestAccess default %s", o._statusCode, payload)
}

func (o *RequestAccessDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /SignUp/RequestAccess][%d] RequestAccess default %s", o._statusCode, payload)
}

func (o *RequestAccessDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RequestAccessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
