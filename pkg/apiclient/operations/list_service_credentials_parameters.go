// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListServiceCredentialsParams creates a new ListServiceCredentialsParams object
// with the default values initialized.
func NewListServiceCredentialsParams() *ListServiceCredentialsParams {
	var ()
	return &ListServiceCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceCredentialsParamsWithTimeout creates a new ListServiceCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServiceCredentialsParamsWithTimeout(timeout time.Duration) *ListServiceCredentialsParams {
	var ()
	return &ListServiceCredentialsParams{

		timeout: timeout,
	}
}

// NewListServiceCredentialsParamsWithContext creates a new ListServiceCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServiceCredentialsParamsWithContext(ctx context.Context) *ListServiceCredentialsParams {
	var ()
	return &ListServiceCredentialsParams{

		Context: ctx,
	}
}

// NewListServiceCredentialsParamsWithHTTPClient creates a new ListServiceCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServiceCredentialsParamsWithHTTPClient(client *http.Client) *ListServiceCredentialsParams {
	var ()
	return &ListServiceCredentialsParams{
		HTTPClient: client,
	}
}

/*ListServiceCredentialsParams contains all the parameters to send to the API endpoint
for the list service credentials operation typically these are written to a http.Request
*/
type ListServiceCredentialsParams struct {

	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list service credentials params
func (o *ListServiceCredentialsParams) WithTimeout(timeout time.Duration) *ListServiceCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service credentials params
func (o *ListServiceCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service credentials params
func (o *ListServiceCredentialsParams) WithContext(ctx context.Context) *ListServiceCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service credentials params
func (o *ListServiceCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service credentials params
func (o *ListServiceCredentialsParams) WithHTTPClient(client *http.Client) *ListServiceCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service credentials params
func (o *ListServiceCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeam adds the team to the list service credentials params
func (o *ListServiceCredentialsParams) WithTeam(team string) *ListServiceCredentialsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the list service credentials params
func (o *ListServiceCredentialsParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}