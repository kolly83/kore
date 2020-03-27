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

// UpdateOrganizationReader is a Reader for the UpdateOrganization structure.
type UpdateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOrganizationOK creates a UpdateOrganizationOK with default headers values
func NewUpdateOrganizationOK() *UpdateOrganizationOK {
	return &UpdateOrganizationOK{}
}

/*UpdateOrganizationOK handles this case with default header values.

Contains the former team definition from the kore
*/
type UpdateOrganizationOK struct {
	Payload *models.V1alpha1Organization
}

func (o *UpdateOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/organizations/{name}][%d] updateOrganizationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationOK) GetPayload() *models.V1alpha1Organization {
	return o.Payload
}

func (o *UpdateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationDefault creates a UpdateOrganizationDefault with default headers values
func NewUpdateOrganizationDefault(code int) *UpdateOrganizationDefault {
	return &UpdateOrganizationDefault{
		_statusCode: code,
	}
}

/*UpdateOrganizationDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateOrganizationDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the update organization default response
func (o *UpdateOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOrganizationDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/organizations/{name}][%d] updateOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateOrganizationDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}