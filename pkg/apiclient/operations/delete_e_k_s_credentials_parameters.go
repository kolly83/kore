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

// NewDeleteEKSCredentialsParams creates a new DeleteEKSCredentialsParams object
// with the default values initialized.
func NewDeleteEKSCredentialsParams() *DeleteEKSCredentialsParams {
	var ()
	return &DeleteEKSCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEKSCredentialsParamsWithTimeout creates a new DeleteEKSCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEKSCredentialsParamsWithTimeout(timeout time.Duration) *DeleteEKSCredentialsParams {
	var ()
	return &DeleteEKSCredentialsParams{

		timeout: timeout,
	}
}

// NewDeleteEKSCredentialsParamsWithContext creates a new DeleteEKSCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEKSCredentialsParamsWithContext(ctx context.Context) *DeleteEKSCredentialsParams {
	var ()
	return &DeleteEKSCredentialsParams{

		Context: ctx,
	}
}

// NewDeleteEKSCredentialsParamsWithHTTPClient creates a new DeleteEKSCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEKSCredentialsParamsWithHTTPClient(client *http.Client) *DeleteEKSCredentialsParams {
	var ()
	return &DeleteEKSCredentialsParams{
		HTTPClient: client,
	}
}

/*DeleteEKSCredentialsParams contains all the parameters to send to the API endpoint
for the delete e k s credentials operation typically these are written to a http.Request
*/
type DeleteEKSCredentialsParams struct {

	/*Name
	  Is name the of the EKS credentials you are acting upon

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

// WithTimeout adds the timeout to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) WithTimeout(timeout time.Duration) *DeleteEKSCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) WithContext(ctx context.Context) *DeleteEKSCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) WithHTTPClient(client *http.Client) *DeleteEKSCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) WithName(name string) *DeleteEKSCredentialsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) WithTeam(team string) *DeleteEKSCredentialsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the delete e k s credentials params
func (o *DeleteEKSCredentialsParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEKSCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
