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

// FindEKSsReader is a Reader for the FindEKSs structure.
type FindEKSsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindEKSsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindEKSsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindEKSsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindEKSsOK creates a FindEKSsOK with default headers values
func NewFindEKSsOK() *FindEKSsOK {
	return &FindEKSsOK{}
}

/*FindEKSsOK handles this case with default header values.

Contains the former team definition from the kore
*/
type FindEKSsOK struct {
	Payload *models.V1alpha1EKSList
}

func (o *FindEKSsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/eks][%d] findEKSsOK  %+v", 200, o.Payload)
}

func (o *FindEKSsOK) GetPayload() *models.V1alpha1EKSList {
	return o.Payload
}

func (o *FindEKSsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1EKSList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindEKSsDefault creates a FindEKSsDefault with default headers values
func NewFindEKSsDefault(code int) *FindEKSsDefault {
	return &FindEKSsDefault{
		_statusCode: code,
	}
}

/*FindEKSsDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type FindEKSsDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the find e k ss default response
func (o *FindEKSsDefault) Code() int {
	return o._statusCode
}

func (o *FindEKSsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/eks][%d] findEKSs default  %+v", o._statusCode, o.Payload)
}

func (o *FindEKSsDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *FindEKSsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
