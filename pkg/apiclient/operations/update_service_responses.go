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

// UpdateServiceReader is a Reader for the UpdateService structure.
type UpdateServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceOK creates a UpdateServiceOK with default headers values
func NewUpdateServiceOK() *UpdateServiceOK {
	return &UpdateServiceOK{}
}

/*UpdateServiceOK handles this case with default header values.

The service details
*/
type UpdateServiceOK struct {
	Payload *models.V1Service
}

func (o *UpdateServiceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/services/{name}][%d] updateServiceOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceOK) GetPayload() *models.V1Service {
	return o.Payload
}

func (o *UpdateServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceBadRequest creates a UpdateServiceBadRequest with default headers values
func NewUpdateServiceBadRequest() *UpdateServiceBadRequest {
	return &UpdateServiceBadRequest{}
}

/*UpdateServiceBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type UpdateServiceBadRequest struct {
	Payload *models.ValidationError
}

func (o *UpdateServiceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/services/{name}][%d] updateServiceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceUnauthorized creates a UpdateServiceUnauthorized with default headers values
func NewUpdateServiceUnauthorized() *UpdateServiceUnauthorized {
	return &UpdateServiceUnauthorized{}
}

/*UpdateServiceUnauthorized handles this case with default header values.

If not authenticated
*/
type UpdateServiceUnauthorized struct {
}

func (o *UpdateServiceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/services/{name}][%d] updateServiceUnauthorized ", 401)
}

func (o *UpdateServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceForbidden creates a UpdateServiceForbidden with default headers values
func NewUpdateServiceForbidden() *UpdateServiceForbidden {
	return &UpdateServiceForbidden{}
}

/*UpdateServiceForbidden handles this case with default header values.

If authenticated but not authorized
*/
type UpdateServiceForbidden struct {
}

func (o *UpdateServiceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/services/{name}][%d] updateServiceForbidden ", 403)
}

func (o *UpdateServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceInternalServerError creates a UpdateServiceInternalServerError with default headers values
func NewUpdateServiceInternalServerError() *UpdateServiceInternalServerError {
	return &UpdateServiceInternalServerError{}
}

/*UpdateServiceInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateServiceInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *UpdateServiceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/services/{name}][%d] updateServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateServiceInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
