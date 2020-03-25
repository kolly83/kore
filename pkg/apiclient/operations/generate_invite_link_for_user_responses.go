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

// GenerateInviteLinkForUserReader is a Reader for the GenerateInviteLinkForUser structure.
type GenerateInviteLinkForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateInviteLinkForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateInviteLinkForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGenerateInviteLinkForUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenerateInviteLinkForUserOK creates a GenerateInviteLinkForUserOK with default headers values
func NewGenerateInviteLinkForUserOK() *GenerateInviteLinkForUserOK {
	return &GenerateInviteLinkForUserOK{}
}

/*GenerateInviteLinkForUserOK handles this case with default header values.

A generated URI which users can use to join the team
*/
type GenerateInviteLinkForUserOK struct {
	Payload string
}

func (o *GenerateInviteLinkForUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/invites/generate/{user}][%d] generateInviteLinkForUserOK  %+v", 200, o.Payload)
}

func (o *GenerateInviteLinkForUserOK) GetPayload() string {
	return o.Payload
}

func (o *GenerateInviteLinkForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateInviteLinkForUserInternalServerError creates a GenerateInviteLinkForUserInternalServerError with default headers values
func NewGenerateInviteLinkForUserInternalServerError() *GenerateInviteLinkForUserInternalServerError {
	return &GenerateInviteLinkForUserInternalServerError{}
}

/*GenerateInviteLinkForUserInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GenerateInviteLinkForUserInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GenerateInviteLinkForUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/invites/generate/{user}][%d] generateInviteLinkForUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateInviteLinkForUserInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GenerateInviteLinkForUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
