// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParams creates a new GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized.
func NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParams() *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout creates a new GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithContext creates a new GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient creates a new GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams contains all the parameters to send to the API endpoint
for the get networks network ID gateways gateway ID description operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithContext(ctx context.Context) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithGatewayID(gatewayID string) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WithNetworkID(networkID string) *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID gateways gateway ID description params
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDGatewaysGatewayIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
