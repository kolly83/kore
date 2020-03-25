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

// NewUpdateNamespaceParams creates a new UpdateNamespaceParams object
// with the default values initialized.
func NewUpdateNamespaceParams() *UpdateNamespaceParams {
	var ()
	return &UpdateNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNamespaceParamsWithTimeout creates a new UpdateNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNamespaceParamsWithTimeout(timeout time.Duration) *UpdateNamespaceParams {
	var ()
	return &UpdateNamespaceParams{

		timeout: timeout,
	}
}

// NewUpdateNamespaceParamsWithContext creates a new UpdateNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNamespaceParamsWithContext(ctx context.Context) *UpdateNamespaceParams {
	var ()
	return &UpdateNamespaceParams{

		Context: ctx,
	}
}

// NewUpdateNamespaceParamsWithHTTPClient creates a new UpdateNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNamespaceParamsWithHTTPClient(client *http.Client) *UpdateNamespaceParams {
	var ()
	return &UpdateNamespaceParams{
		HTTPClient: client,
	}
}

/*UpdateNamespaceParams contains all the parameters to send to the API endpoint
for the update namespace operation typically these are written to a http.Request
*/
type UpdateNamespaceParams struct {

	/*Body
	  The definition for namespace claim

	*/
	Body *models.V1NamespaceClaim
	/*Name
	  Is name the of the namespace claim you are acting upon

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

// WithTimeout adds the timeout to the update namespace params
func (o *UpdateNamespaceParams) WithTimeout(timeout time.Duration) *UpdateNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update namespace params
func (o *UpdateNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update namespace params
func (o *UpdateNamespaceParams) WithContext(ctx context.Context) *UpdateNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update namespace params
func (o *UpdateNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update namespace params
func (o *UpdateNamespaceParams) WithHTTPClient(client *http.Client) *UpdateNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update namespace params
func (o *UpdateNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update namespace params
func (o *UpdateNamespaceParams) WithBody(body *models.V1NamespaceClaim) *UpdateNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update namespace params
func (o *UpdateNamespaceParams) SetBody(body *models.V1NamespaceClaim) {
	o.Body = body
}

// WithName adds the name to the update namespace params
func (o *UpdateNamespaceParams) WithName(name string) *UpdateNamespaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update namespace params
func (o *UpdateNamespaceParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the update namespace params
func (o *UpdateNamespaceParams) WithTeam(team string) *UpdateNamespaceParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update namespace params
func (o *UpdateNamespaceParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
