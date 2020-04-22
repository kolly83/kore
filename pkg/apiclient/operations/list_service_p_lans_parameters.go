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

// NewListServicePLansParams creates a new ListServicePLansParams object
// with the default values initialized.
func NewListServicePLansParams() *ListServicePLansParams {
	var ()
	return &ListServicePLansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServicePLansParamsWithTimeout creates a new ListServicePLansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServicePLansParamsWithTimeout(timeout time.Duration) *ListServicePLansParams {
	var ()
	return &ListServicePLansParams{

		timeout: timeout,
	}
}

// NewListServicePLansParamsWithContext creates a new ListServicePLansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServicePLansParamsWithContext(ctx context.Context) *ListServicePLansParams {
	var ()
	return &ListServicePLansParams{

		Context: ctx,
	}
}

// NewListServicePLansParamsWithHTTPClient creates a new ListServicePLansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServicePLansParamsWithHTTPClient(client *http.Client) *ListServicePLansParams {
	var ()
	return &ListServicePLansParams{
		HTTPClient: client,
	}
}

/*ListServicePLansParams contains all the parameters to send to the API endpoint
for the list service p lans operation typically these are written to a http.Request
*/
type ListServicePLansParams struct {

	/*Kind
	  Filters service plans for a specific kind

	*/
	Kind *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list service p lans params
func (o *ListServicePLansParams) WithTimeout(timeout time.Duration) *ListServicePLansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service p lans params
func (o *ListServicePLansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service p lans params
func (o *ListServicePLansParams) WithContext(ctx context.Context) *ListServicePLansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service p lans params
func (o *ListServicePLansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service p lans params
func (o *ListServicePLansParams) WithHTTPClient(client *http.Client) *ListServicePLansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service p lans params
func (o *ListServicePLansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKind adds the kind to the list service p lans params
func (o *ListServicePLansParams) WithKind(kind *string) *ListServicePLansParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the list service p lans params
func (o *ListServicePLansParams) SetKind(kind *string) {
	o.Kind = kind
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicePLansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
