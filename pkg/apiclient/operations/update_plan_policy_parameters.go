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

// NewUpdatePlanPolicyParams creates a new UpdatePlanPolicyParams object
// with the default values initialized.
func NewUpdatePlanPolicyParams() *UpdatePlanPolicyParams {
	var ()
	return &UpdatePlanPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePlanPolicyParamsWithTimeout creates a new UpdatePlanPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePlanPolicyParamsWithTimeout(timeout time.Duration) *UpdatePlanPolicyParams {
	var ()
	return &UpdatePlanPolicyParams{

		timeout: timeout,
	}
}

// NewUpdatePlanPolicyParamsWithContext creates a new UpdatePlanPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePlanPolicyParamsWithContext(ctx context.Context) *UpdatePlanPolicyParams {
	var ()
	return &UpdatePlanPolicyParams{

		Context: ctx,
	}
}

// NewUpdatePlanPolicyParamsWithHTTPClient creates a new UpdatePlanPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePlanPolicyParamsWithHTTPClient(client *http.Client) *UpdatePlanPolicyParams {
	var ()
	return &UpdatePlanPolicyParams{
		HTTPClient: client,
	}
}

/*UpdatePlanPolicyParams contains all the parameters to send to the API endpoint
for the update plan policy operation typically these are written to a http.Request
*/
type UpdatePlanPolicyParams struct {

	/*Body
	  The specification for the plan policy you are updating

	*/
	Body *models.V1PlanPolicy
	/*Name
	  The name of the plan policy you wish to update

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update plan policy params
func (o *UpdatePlanPolicyParams) WithTimeout(timeout time.Duration) *UpdatePlanPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update plan policy params
func (o *UpdatePlanPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update plan policy params
func (o *UpdatePlanPolicyParams) WithContext(ctx context.Context) *UpdatePlanPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update plan policy params
func (o *UpdatePlanPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update plan policy params
func (o *UpdatePlanPolicyParams) WithHTTPClient(client *http.Client) *UpdatePlanPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update plan policy params
func (o *UpdatePlanPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update plan policy params
func (o *UpdatePlanPolicyParams) WithBody(body *models.V1PlanPolicy) *UpdatePlanPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update plan policy params
func (o *UpdatePlanPolicyParams) SetBody(body *models.V1PlanPolicy) {
	o.Body = body
}

// WithName adds the name to the update plan policy params
func (o *UpdatePlanPolicyParams) WithName(name string) *UpdatePlanPolicyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update plan policy params
func (o *UpdatePlanPolicyParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePlanPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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