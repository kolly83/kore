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

// NewListAuditEventsParams creates a new ListAuditEventsParams object
// with the default values initialized.
func NewListAuditEventsParams() *ListAuditEventsParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListAuditEventsParams{
		Since: &sinceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListAuditEventsParamsWithTimeout creates a new ListAuditEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAuditEventsParamsWithTimeout(timeout time.Duration) *ListAuditEventsParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListAuditEventsParams{
		Since: &sinceDefault,

		timeout: timeout,
	}
}

// NewListAuditEventsParamsWithContext creates a new ListAuditEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAuditEventsParamsWithContext(ctx context.Context) *ListAuditEventsParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListAuditEventsParams{
		Since: &sinceDefault,

		Context: ctx,
	}
}

// NewListAuditEventsParamsWithHTTPClient creates a new ListAuditEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAuditEventsParamsWithHTTPClient(client *http.Client) *ListAuditEventsParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListAuditEventsParams{
		Since:      &sinceDefault,
		HTTPClient: client,
	}
}

/*ListAuditEventsParams contains all the parameters to send to the API endpoint
for the list audit events operation typically these are written to a http.Request
*/
type ListAuditEventsParams struct {

	/*Since
	  The time duration to return the events within

	*/
	Since *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list audit events params
func (o *ListAuditEventsParams) WithTimeout(timeout time.Duration) *ListAuditEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list audit events params
func (o *ListAuditEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list audit events params
func (o *ListAuditEventsParams) WithContext(ctx context.Context) *ListAuditEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list audit events params
func (o *ListAuditEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list audit events params
func (o *ListAuditEventsParams) WithHTTPClient(client *http.Client) *ListAuditEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list audit events params
func (o *ListAuditEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSince adds the since to the list audit events params
func (o *ListAuditEventsParams) WithSince(since *string) *ListAuditEventsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the list audit events params
func (o *ListAuditEventsParams) SetSince(since *string) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *ListAuditEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
