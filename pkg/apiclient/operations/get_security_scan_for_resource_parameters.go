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

// NewGetSecurityScanForResourceParams creates a new GetSecurityScanForResourceParams object
// with the default values initialized.
func NewGetSecurityScanForResourceParams() *GetSecurityScanForResourceParams {
	var ()
	return &GetSecurityScanForResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityScanForResourceParamsWithTimeout creates a new GetSecurityScanForResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecurityScanForResourceParamsWithTimeout(timeout time.Duration) *GetSecurityScanForResourceParams {
	var ()
	return &GetSecurityScanForResourceParams{

		timeout: timeout,
	}
}

// NewGetSecurityScanForResourceParamsWithContext creates a new GetSecurityScanForResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecurityScanForResourceParamsWithContext(ctx context.Context) *GetSecurityScanForResourceParams {
	var ()
	return &GetSecurityScanForResourceParams{

		Context: ctx,
	}
}

// NewGetSecurityScanForResourceParamsWithHTTPClient creates a new GetSecurityScanForResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecurityScanForResourceParamsWithHTTPClient(client *http.Client) *GetSecurityScanForResourceParams {
	var ()
	return &GetSecurityScanForResourceParams{
		HTTPClient: client,
	}
}

/*GetSecurityScanForResourceParams contains all the parameters to send to the API endpoint
for the get security scan for resource operation typically these are written to a http.Request
*/
type GetSecurityScanForResourceParams struct {

	/*Group
	  Is the group of the kind

	*/
	Group string
	/*Kind
	  Is the kind of the resource

	*/
	Kind string
	/*Name
	  Is the name of the resource

	*/
	Name string
	/*Namespace
	  Is the namespace of the resource

	*/
	Namespace string
	/*Version
	  Is the version of the kind

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithTimeout(timeout time.Duration) *GetSecurityScanForResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithContext(ctx context.Context) *GetSecurityScanForResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithHTTPClient(client *http.Client) *GetSecurityScanForResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithGroup(group string) *GetSecurityScanForResourceParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetGroup(group string) {
	o.Group = group
}

// WithKind adds the kind to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithKind(kind string) *GetSecurityScanForResourceParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetKind(kind string) {
	o.Kind = kind
}

// WithName adds the name to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithName(name string) *GetSecurityScanForResourceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithNamespace(namespace string) *GetSecurityScanForResourceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVersion adds the version to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) WithVersion(version string) *GetSecurityScanForResourceParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get security scan for resource params
func (o *GetSecurityScanForResourceParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityScanForResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group
	if err := r.SetPathParam("group", o.Group); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
