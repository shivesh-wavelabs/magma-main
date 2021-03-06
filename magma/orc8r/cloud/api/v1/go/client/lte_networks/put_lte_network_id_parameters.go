// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewPutLTENetworkIDParams creates a new PutLTENetworkIDParams object
// with the default values initialized.
func NewPutLTENetworkIDParams() *PutLTENetworkIDParams {
	var ()
	return &PutLTENetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDParamsWithTimeout creates a new PutLTENetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDParams {
	var ()
	return &PutLTENetworkIDParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDParamsWithContext creates a new PutLTENetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDParamsWithContext(ctx context.Context) *PutLTENetworkIDParams {
	var ()
	return &PutLTENetworkIDParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDParamsWithHTTPClient creates a new PutLTENetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDParams {
	var ()
	return &PutLTENetworkIDParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDParams contains all the parameters to send to the API endpoint
for the put LTE network ID operation typically these are written to a http.Request
*/
type PutLTENetworkIDParams struct {

	/*LTENetwork
	  Full desired configuration of the network

	*/
	LTENetwork *models.LTENetwork
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID params
func (o *PutLTENetworkIDParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID params
func (o *PutLTENetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID params
func (o *PutLTENetworkIDParams) WithContext(ctx context.Context) *PutLTENetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID params
func (o *PutLTENetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID params
func (o *PutLTENetworkIDParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID params
func (o *PutLTENetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLTENetwork adds the lTENetwork to the put LTE network ID params
func (o *PutLTENetworkIDParams) WithLTENetwork(lTENetwork *models.LTENetwork) *PutLTENetworkIDParams {
	o.SetLTENetwork(lTENetwork)
	return o
}

// SetLTENetwork adds the lteNetwork to the put LTE network ID params
func (o *PutLTENetworkIDParams) SetLTENetwork(lTENetwork *models.LTENetwork) {
	o.LTENetwork = lTENetwork
}

// WithNetworkID adds the networkID to the put LTE network ID params
func (o *PutLTENetworkIDParams) WithNetworkID(networkID string) *PutLTENetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID params
func (o *PutLTENetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LTENetwork != nil {
		if err := r.SetBodyParam(o.LTENetwork); err != nil {
			return err
		}
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
