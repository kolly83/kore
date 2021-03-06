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

// UpdateServiceCredentialsReader is a Reader for the UpdateServiceCredentials structure.
type UpdateServiceCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateServiceCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateServiceCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceCredentialsOK creates a UpdateServiceCredentialsOK with default headers values
func NewUpdateServiceCredentialsOK() *UpdateServiceCredentialsOK {
	return &UpdateServiceCredentialsOK{}
}

/*UpdateServiceCredentialsOK handles this case with default header values.

The service credentail details
*/
type UpdateServiceCredentialsOK struct {
	Payload *models.V1ServiceCredentials
}

func (o *UpdateServiceCredentialsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] updateServiceCredentialsOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceCredentialsOK) GetPayload() *models.V1ServiceCredentials {
	return o.Payload
}

func (o *UpdateServiceCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ServiceCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceCredentialsBadRequest creates a UpdateServiceCredentialsBadRequest with default headers values
func NewUpdateServiceCredentialsBadRequest() *UpdateServiceCredentialsBadRequest {
	return &UpdateServiceCredentialsBadRequest{}
}

/*UpdateServiceCredentialsBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type UpdateServiceCredentialsBadRequest struct {
	Payload *models.ValidationError
}

func (o *UpdateServiceCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] updateServiceCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceCredentialsBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateServiceCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceCredentialsUnauthorized creates a UpdateServiceCredentialsUnauthorized with default headers values
func NewUpdateServiceCredentialsUnauthorized() *UpdateServiceCredentialsUnauthorized {
	return &UpdateServiceCredentialsUnauthorized{}
}

/*UpdateServiceCredentialsUnauthorized handles this case with default header values.

If not authenticated
*/
type UpdateServiceCredentialsUnauthorized struct {
}

func (o *UpdateServiceCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] updateServiceCredentialsUnauthorized ", 401)
}

func (o *UpdateServiceCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceCredentialsForbidden creates a UpdateServiceCredentialsForbidden with default headers values
func NewUpdateServiceCredentialsForbidden() *UpdateServiceCredentialsForbidden {
	return &UpdateServiceCredentialsForbidden{}
}

/*UpdateServiceCredentialsForbidden handles this case with default header values.

If authenticated but not authorized
*/
type UpdateServiceCredentialsForbidden struct {
}

func (o *UpdateServiceCredentialsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] updateServiceCredentialsForbidden ", 403)
}

func (o *UpdateServiceCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceCredentialsInternalServerError creates a UpdateServiceCredentialsInternalServerError with default headers values
func NewUpdateServiceCredentialsInternalServerError() *UpdateServiceCredentialsInternalServerError {
	return &UpdateServiceCredentialsInternalServerError{}
}

/*UpdateServiceCredentialsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateServiceCredentialsInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *UpdateServiceCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] updateServiceCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateServiceCredentialsInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateServiceCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
