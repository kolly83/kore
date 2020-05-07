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

// NewGetServicePlanSchemaParams creates a new GetServicePlanSchemaParams object
// with the default values initialized.
func NewGetServicePlanSchemaParams() *GetServicePlanSchemaParams {
	var ()
	return &GetServicePlanSchemaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicePlanSchemaParamsWithTimeout creates a new GetServicePlanSchemaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicePlanSchemaParamsWithTimeout(timeout time.Duration) *GetServicePlanSchemaParams {
	var ()
	return &GetServicePlanSchemaParams{

		timeout: timeout,
	}
}

// NewGetServicePlanSchemaParamsWithContext creates a new GetServicePlanSchemaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicePlanSchemaParamsWithContext(ctx context.Context) *GetServicePlanSchemaParams {
	var ()
	return &GetServicePlanSchemaParams{

		Context: ctx,
	}
}

// NewGetServicePlanSchemaParamsWithHTTPClient creates a new GetServicePlanSchemaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicePlanSchemaParamsWithHTTPClient(client *http.Client) *GetServicePlanSchemaParams {
	var ()
	return &GetServicePlanSchemaParams{
		HTTPClient: client,
	}
}

/*GetServicePlanSchemaParams contains all the parameters to send to the API endpoint
for the get service plan schema operation typically these are written to a http.Request
*/
type GetServicePlanSchemaParams struct {

	/*Name
	  The name of the service plan

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service plan schema params
func (o *GetServicePlanSchemaParams) WithTimeout(timeout time.Duration) *GetServicePlanSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service plan schema params
func (o *GetServicePlanSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service plan schema params
func (o *GetServicePlanSchemaParams) WithContext(ctx context.Context) *GetServicePlanSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service plan schema params
func (o *GetServicePlanSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service plan schema params
func (o *GetServicePlanSchemaParams) WithHTTPClient(client *http.Client) *GetServicePlanSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service plan schema params
func (o *GetServicePlanSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get service plan schema params
func (o *GetServicePlanSchemaParams) WithName(name string) *GetServicePlanSchemaParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get service plan schema params
func (o *GetServicePlanSchemaParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicePlanSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
