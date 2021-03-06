// Code generated by go-swagger; DO NOT EDIT.

package commands

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

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParams creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams object
// with the default values initialized.
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParams() *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	var ()
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParamsWithTimeout creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParamsWithTimeout(timeout time.Duration) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	var ()
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams{

		timeout: timeout,
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParamsWithContext creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParamsWithContext(ctx context.Context) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	var ()
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams{

		Context: ctx,
	}
}

// NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParamsWithHTTPClient creates a new PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNetworksNetworkIDGatewaysGatewayIDCommandPingParamsWithHTTPClient(client *http.Client) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	var ()
	return &PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams{
		HTTPClient: client,
	}
}

/*PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams contains all the parameters to send to the API endpoint
for the post networks network ID gateways gateway ID command ping operation typically these are written to a http.Request
*/
type PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams struct {

	/*PingRequest
	  Ping request

	*/
	PingRequest *models.PingRequest
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

// WithTimeout adds the timeout to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WithTimeout(timeout time.Duration) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WithContext(ctx context.Context) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WithHTTPClient(client *http.Client) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPingRequest adds the pingRequest to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WithPingRequest(pingRequest *models.PingRequest) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	o.SetPingRequest(pingRequest)
	return o
}

// SetPingRequest adds the pingRequest to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) SetPingRequest(pingRequest *models.PingRequest) {
	o.PingRequest = pingRequest
}

// WithGatewayID adds the gatewayID to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WithGatewayID(gatewayID string) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WithNetworkID(networkID string) *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post networks network ID gateways gateway ID command ping params
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksNetworkIDGatewaysGatewayIDCommandPingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PingRequest != nil {
		if err := r.SetBodyParam(o.PingRequest); err != nil {
			return err
		}
	}

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
