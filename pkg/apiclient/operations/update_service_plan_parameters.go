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

// NewUpdateServicePlanParams creates a new UpdateServicePlanParams object
// with the default values initialized.
func NewUpdateServicePlanParams() *UpdateServicePlanParams {
	var ()
	return &UpdateServicePlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServicePlanParamsWithTimeout creates a new UpdateServicePlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServicePlanParamsWithTimeout(timeout time.Duration) *UpdateServicePlanParams {
	var ()
	return &UpdateServicePlanParams{

		timeout: timeout,
	}
}

// NewUpdateServicePlanParamsWithContext creates a new UpdateServicePlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServicePlanParamsWithContext(ctx context.Context) *UpdateServicePlanParams {
	var ()
	return &UpdateServicePlanParams{

		Context: ctx,
	}
}

// NewUpdateServicePlanParamsWithHTTPClient creates a new UpdateServicePlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServicePlanParamsWithHTTPClient(client *http.Client) *UpdateServicePlanParams {
	var ()
	return &UpdateServicePlanParams{
		HTTPClient: client,
	}
}

/*UpdateServicePlanParams contains all the parameters to send to the API endpoint
for the update service plan operation typically these are written to a http.Request
*/
type UpdateServicePlanParams struct {

	/*Body
	  The specification for the service plan you are creating or updating

	*/
	Body *models.V1ServicePlan
	/*Name
	  The name of the service plan you wish to create or update

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service plan params
func (o *UpdateServicePlanParams) WithTimeout(timeout time.Duration) *UpdateServicePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service plan params
func (o *UpdateServicePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service plan params
func (o *UpdateServicePlanParams) WithContext(ctx context.Context) *UpdateServicePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service plan params
func (o *UpdateServicePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service plan params
func (o *UpdateServicePlanParams) WithHTTPClient(client *http.Client) *UpdateServicePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service plan params
func (o *UpdateServicePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service plan params
func (o *UpdateServicePlanParams) WithBody(body *models.V1ServicePlan) *UpdateServicePlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service plan params
func (o *UpdateServicePlanParams) SetBody(body *models.V1ServicePlan) {
	o.Body = body
}

// WithName adds the name to the update service plan params
func (o *UpdateServicePlanParams) WithName(name string) *UpdateServicePlanParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update service plan params
func (o *UpdateServicePlanParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServicePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
