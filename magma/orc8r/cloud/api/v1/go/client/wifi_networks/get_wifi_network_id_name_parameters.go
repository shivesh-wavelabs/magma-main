// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWifiNetworkIDNameParams creates a new GetWifiNetworkIDNameParams object
// with the default values initialized.
func NewGetWifiNetworkIDNameParams() *GetWifiNetworkIDNameParams {
	var ()
	return &GetWifiNetworkIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWifiNetworkIDNameParamsWithTimeout creates a new GetWifiNetworkIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWifiNetworkIDNameParamsWithTimeout(timeout time.Duration) *GetWifiNetworkIDNameParams {
	var ()
	return &GetWifiNetworkIDNameParams{

		timeout: timeout,
	}
}

// NewGetWifiNetworkIDNameParamsWithContext creates a new GetWifiNetworkIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWifiNetworkIDNameParamsWithContext(ctx context.Context) *GetWifiNetworkIDNameParams {
	var ()
	return &GetWifiNetworkIDNameParams{

		Context: ctx,
	}
}

// NewGetWifiNetworkIDNameParamsWithHTTPClient creates a new GetWifiNetworkIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWifiNetworkIDNameParamsWithHTTPClient(client *http.Client) *GetWifiNetworkIDNameParams {
	var ()
	return &GetWifiNetworkIDNameParams{
		HTTPClient: client,
	}
}

/*GetWifiNetworkIDNameParams contains all the parameters to send to the API endpoint
for the get wifi network ID name operation typically these are written to a http.Request
*/
type GetWifiNetworkIDNameParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) WithTimeout(timeout time.Duration) *GetWifiNetworkIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) WithContext(ctx context.Context) *GetWifiNetworkIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) WithHTTPClient(client *http.Client) *GetWifiNetworkIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) WithNetworkID(networkID string) *GetWifiNetworkIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get wifi network ID name params
func (o *GetWifiNetworkIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWifiNetworkIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
