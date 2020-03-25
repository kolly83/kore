// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// LoginAttemptedReader is a Reader for the LoginAttempted structure.
type LoginAttemptedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginAttemptedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginAttemptedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLoginAttemptedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginAttemptedOK creates a LoginAttemptedOK with default headers values
func NewLoginAttemptedOK() *LoginAttemptedOK {
	return &LoginAttemptedOK{}
}

/*LoginAttemptedOK handles this case with default header values.

OK
*/
type LoginAttemptedOK struct {
}

func (o *LoginAttemptedOK) Error() string {
	return fmt.Sprintf("[GET /oauth/authorize][%d] loginAttemptedOK ", 200)
}

func (o *LoginAttemptedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginAttemptedDefault creates a LoginAttemptedDefault with default headers values
func NewLoginAttemptedDefault(code int) *LoginAttemptedDefault {
	return &LoginAttemptedDefault{
		_statusCode: code,
	}
}

/*LoginAttemptedDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type LoginAttemptedDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the login attempted default response
func (o *LoginAttemptedDefault) Code() int {
	return o._statusCode
}

func (o *LoginAttemptedDefault) Error() string {
	return fmt.Sprintf("[GET /oauth/authorize][%d] LoginAttempted default  %+v", o._statusCode, o.Payload)
}

func (o *LoginAttemptedDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *LoginAttemptedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
