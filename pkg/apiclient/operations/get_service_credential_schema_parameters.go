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

// NewGetServiceCredentialSchemaParams creates a new GetServiceCredentialSchemaParams object
// with the default values initialized.
func NewGetServiceCredentialSchemaParams() *GetServiceCredentialSchemaParams {
	var ()
	return &GetServiceCredentialSchemaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceCredentialSchemaParamsWithTimeout creates a new GetServiceCredentialSchemaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceCredentialSchemaParamsWithTimeout(timeout time.Duration) *GetServiceCredentialSchemaParams {
	var ()
	return &GetServiceCredentialSchemaParams{

		timeout: timeout,
	}
}

// NewGetServiceCredentialSchemaParamsWithContext creates a new GetServiceCredentialSchemaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceCredentialSchemaParamsWithContext(ctx context.Context) *GetServiceCredentialSchemaParams {
	var ()
	return &GetServiceCredentialSchemaParams{

		Context: ctx,
	}
}

// NewGetServiceCredentialSchemaParamsWithHTTPClient creates a new GetServiceCredentialSchemaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceCredentialSchemaParamsWithHTTPClient(client *http.Client) *GetServiceCredentialSchemaParams {
	var ()
	return &GetServiceCredentialSchemaParams{
		HTTPClient: client,
	}
}

/*GetServiceCredentialSchemaParams contains all the parameters to send to the API endpoint
for the get service credential schema operation typically these are written to a http.Request
*/
type GetServiceCredentialSchemaParams struct {

	/*Name
	  The name of the service plan

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) WithTimeout(timeout time.Duration) *GetServiceCredentialSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) WithContext(ctx context.Context) *GetServiceCredentialSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) WithHTTPClient(client *http.Client) *GetServiceCredentialSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) WithName(name string) *GetServiceCredentialSchemaParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get service credential schema params
func (o *GetServiceCredentialSchemaParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceCredentialSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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