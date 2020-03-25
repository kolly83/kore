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

// NewDeleteOrganizationParams creates a new DeleteOrganizationParams object
// with the default values initialized.
func NewDeleteOrganizationParams() *DeleteOrganizationParams {
	var ()
	return &DeleteOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationParamsWithTimeout creates a new DeleteOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrganizationParamsWithTimeout(timeout time.Duration) *DeleteOrganizationParams {
	var ()
	return &DeleteOrganizationParams{

		timeout: timeout,
	}
}

// NewDeleteOrganizationParamsWithContext creates a new DeleteOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrganizationParamsWithContext(ctx context.Context) *DeleteOrganizationParams {
	var ()
	return &DeleteOrganizationParams{

		Context: ctx,
	}
}

// NewDeleteOrganizationParamsWithHTTPClient creates a new DeleteOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrganizationParamsWithHTTPClient(client *http.Client) *DeleteOrganizationParams {
	var ()
	return &DeleteOrganizationParams{
		HTTPClient: client,
	}
}

/*DeleteOrganizationParams contains all the parameters to send to the API endpoint
for the delete organization operation typically these are written to a http.Request
*/
type DeleteOrganizationParams struct {

	/*Name
	  Is name the of the resource you are acting on

	*/
	Name string
	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete organization params
func (o *DeleteOrganizationParams) WithTimeout(timeout time.Duration) *DeleteOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization params
func (o *DeleteOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization params
func (o *DeleteOrganizationParams) WithContext(ctx context.Context) *DeleteOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization params
func (o *DeleteOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization params
func (o *DeleteOrganizationParams) WithHTTPClient(client *http.Client) *DeleteOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization params
func (o *DeleteOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete organization params
func (o *DeleteOrganizationParams) WithName(name string) *DeleteOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete organization params
func (o *DeleteOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the delete organization params
func (o *DeleteOrganizationParams) WithTeam(team string) *DeleteOrganizationParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the delete organization params
func (o *DeleteOrganizationParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
