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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutNetworksNetworkIDGatewaysGatewayIDTierParams creates a new PutNetworksNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized.
func NewPutNetworksNetworkIDGatewaysGatewayIDTierParams() *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDTierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDTierParamsWithTimeout creates a new PutNetworksNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDGatewaysGatewayIDTierParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDTierParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDTierParamsWithContext creates a new PutNetworksNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDGatewaysGatewayIDTierParamsWithContext(ctx context.Context) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDTierParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDTierParamsWithHTTPClient creates a new PutNetworksNetworkIDGatewaysGatewayIDTierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDGatewaysGatewayIDTierParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	var ()
	return &PutNetworksNetworkIDGatewaysGatewayIDTierParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDGatewaysGatewayIDTierParams contains all the parameters to send to the API endpoint
for the put networks network ID gateways gateway ID tier operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDGatewaysGatewayIDTierParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*TierID*/
	TierID models.TierID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WithContext(ctx context.Context) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WithGatewayID(gatewayID string) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WithNetworkID(networkID string) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTierID adds the tierID to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WithTierID(tierID models.TierID) *PutNetworksNetworkIDGatewaysGatewayIDTierParams {
	o.SetTierID(tierID)
	return o
}

// SetTierID adds the tierId to the put networks network ID gateways gateway ID tier params
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) SetTierID(tierID models.TierID) {
	o.TierID = tierID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.TierID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
