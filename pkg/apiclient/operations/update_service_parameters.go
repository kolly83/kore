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

	"github.com/appvia/kore/pkg/apiclient/models"
)

// NewUpdateServiceParams creates a new UpdateServiceParams object
// with the default values initialized.
func NewUpdateServiceParams() *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceParamsWithTimeout creates a new UpdateServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceParamsWithTimeout(timeout time.Duration) *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{

		timeout: timeout,
	}
}

// NewUpdateServiceParamsWithContext creates a new UpdateServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceParamsWithContext(ctx context.Context) *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{

		Context: ctx,
	}
}

// NewUpdateServiceParamsWithHTTPClient creates a new UpdateServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceParamsWithHTTPClient(client *http.Client) *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{
		HTTPClient: client,
	}
}

/*UpdateServiceParams contains all the parameters to send to the API endpoint
for the update service operation typically these are written to a http.Request
*/
type UpdateServiceParams struct {

	/*Body
	  The definition for the service

	*/
	Body *models.V1Service
	/*Name
	  Is name the of the service

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

// WithTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) WithTimeout(timeout time.Duration) *UpdateServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service params
func (o *UpdateServiceParams) WithContext(ctx context.Context) *UpdateServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service params
func (o *UpdateServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) WithHTTPClient(client *http.Client) *UpdateServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service params
func (o *UpdateServiceParams) WithBody(body *models.V1Service) *UpdateServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service params
func (o *UpdateServiceParams) SetBody(body *models.V1Service) {
	o.Body = body
}

// WithName adds the name to the update service params
func (o *UpdateServiceParams) WithName(name string) *UpdateServiceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update service params
func (o *UpdateServiceParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the update service params
func (o *UpdateServiceParams) WithTeam(team string) *UpdateServiceParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update service params
func (o *UpdateServiceParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
