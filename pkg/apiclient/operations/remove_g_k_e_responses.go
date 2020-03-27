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

// RemoveGKEReader is a Reader for the RemoveGKE structure.
type RemoveGKEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveGKEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveGKEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRemoveGKEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveGKEOK creates a RemoveGKEOK with default headers values
func NewRemoveGKEOK() *RemoveGKEOK {
	return &RemoveGKEOK{}
}

/*RemoveGKEOK handles this case with default header values.

Contains the former team definition from the kore
*/
type RemoveGKEOK struct {
	Payload *models.V1alpha1GKE
}

func (o *RemoveGKEOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/gkes/{name}][%d] removeGKEOK  %+v", 200, o.Payload)
}

func (o *RemoveGKEOK) GetPayload() *models.V1alpha1GKE {
	return o.Payload
}

func (o *RemoveGKEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1GKE)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGKEInternalServerError creates a RemoveGKEInternalServerError with default headers values
func NewRemoveGKEInternalServerError() *RemoveGKEInternalServerError {
	return &RemoveGKEInternalServerError{}
}

/*RemoveGKEInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type RemoveGKEInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *RemoveGKEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/gkes/{name}][%d] removeGKEInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveGKEInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *RemoveGKEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}