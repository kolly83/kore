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

// NewListServicesParams creates a new ListServicesParams object
// with the default values initialized.
func NewListServicesParams() *ListServicesParams {
	var ()
	return &ListServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServicesParamsWithTimeout creates a new ListServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServicesParamsWithTimeout(timeout time.Duration) *ListServicesParams {
	var ()
	return &ListServicesParams{

		timeout: timeout,
	}
}

// NewListServicesParamsWithContext creates a new ListServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServicesParamsWithContext(ctx context.Context) *ListServicesParams {
	var ()
	return &ListServicesParams{

		Context: ctx,
	}
}

// NewListServicesParamsWithHTTPClient creates a new ListServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServicesParamsWithHTTPClient(client *http.Client) *ListServicesParams {
	var ()
	return &ListServicesParams{
		HTTPClient: client,
	}
}

/*ListServicesParams contains all the parameters to send to the API endpoint
for the list services operation typically these are written to a http.Request
*/
type ListServicesParams struct {

	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list services params
func (o *ListServicesParams) WithTimeout(timeout time.Duration) *ListServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list services params
func (o *ListServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list services params
func (o *ListServicesParams) WithContext(ctx context.Context) *ListServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list services params
func (o *ListServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list services params
func (o *ListServicesParams) WithHTTPClient(client *http.Client) *ListServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list services params
func (o *ListServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeam adds the team to the list services params
func (o *ListServicesParams) WithTeam(team string) *ListServicesParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the list services params
func (o *ListServicesParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
