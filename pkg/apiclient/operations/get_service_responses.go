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

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK handles this case with default header values.

The requested service details
*/
type GetServiceOK struct {
	Payload *models.V1Service
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services/{name}][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) GetPayload() *models.V1Service {
	return o.Payload
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceUnauthorized creates a GetServiceUnauthorized with default headers values
func NewGetServiceUnauthorized() *GetServiceUnauthorized {
	return &GetServiceUnauthorized{}
}

/*GetServiceUnauthorized handles this case with default header values.

If not authenticated
*/
type GetServiceUnauthorized struct {
}

func (o *GetServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services/{name}][%d] getServiceUnauthorized ", 401)
}

func (o *GetServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceForbidden creates a GetServiceForbidden with default headers values
func NewGetServiceForbidden() *GetServiceForbidden {
	return &GetServiceForbidden{}
}

/*GetServiceForbidden handles this case with default header values.

If authenticated but not authorized
*/
type GetServiceForbidden struct {
}

func (o *GetServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services/{name}][%d] getServiceForbidden ", 403)
}

func (o *GetServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceNotFound creates a GetServiceNotFound with default headers values
func NewGetServiceNotFound() *GetServiceNotFound {
	return &GetServiceNotFound{}
}

/*GetServiceNotFound handles this case with default header values.

the service with the given name doesn't exist
*/
type GetServiceNotFound struct {
}

func (o *GetServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services/{name}][%d] getServiceNotFound ", 404)
}

func (o *GetServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceInternalServerError creates a GetServiceInternalServerError with default headers values
func NewGetServiceInternalServerError() *GetServiceInternalServerError {
	return &GetServiceInternalServerError{}
}

/*GetServiceInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetServiceInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetServiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services/{name}][%d] getServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
