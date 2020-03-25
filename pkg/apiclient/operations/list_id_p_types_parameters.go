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

// NewListIDPTypesParams creates a new ListIDPTypesParams object
// with the default values initialized.
func NewListIDPTypesParams() *ListIDPTypesParams {

	return &ListIDPTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListIDPTypesParamsWithTimeout creates a new ListIDPTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIDPTypesParamsWithTimeout(timeout time.Duration) *ListIDPTypesParams {

	return &ListIDPTypesParams{

		timeout: timeout,
	}
}

// NewListIDPTypesParamsWithContext creates a new ListIDPTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIDPTypesParamsWithContext(ctx context.Context) *ListIDPTypesParams {

	return &ListIDPTypesParams{

		Context: ctx,
	}
}

// NewListIDPTypesParamsWithHTTPClient creates a new ListIDPTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIDPTypesParamsWithHTTPClient(client *http.Client) *ListIDPTypesParams {

	return &ListIDPTypesParams{
		HTTPClient: client,
	}
}

/*ListIDPTypesParams contains all the parameters to send to the API endpoint
for the list ID p types operation typically these are written to a http.Request
*/
type ListIDPTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list ID p types params
func (o *ListIDPTypesParams) WithTimeout(timeout time.Duration) *ListIDPTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ID p types params
func (o *ListIDPTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ID p types params
func (o *ListIDPTypesParams) WithContext(ctx context.Context) *ListIDPTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ID p types params
func (o *ListIDPTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ID p types params
func (o *ListIDPTypesParams) WithHTTPClient(client *http.Client) *ListIDPTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ID p types params
func (o *ListIDPTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListIDPTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
