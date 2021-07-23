// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

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

// NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams creates a new DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized.
func NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams() *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithTimeout creates a new DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithTimeout(timeout time.Duration) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams{

		timeout: timeout,
	}
}

// NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithContext creates a new DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithContext(ctx context.Context) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams{

		Context: ctx,
	}
}

// NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithHTTPClient creates a new DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParamsWithHTTPClient(client *http.Client) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams{
		HTTPClient: client,
	}
}

/*DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams contains all the parameters to send to the API endpoint
for the delete feg network ID subscriber config rule names rule ID operation typically these are written to a http.Request
*/
type DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*RuleID
	  Rule Id

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithTimeout(timeout time.Duration) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithContext(ctx context.Context) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithHTTPClient(client *http.Client) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithNetworkID(networkID string) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRuleID adds the ruleID to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) WithRuleID(ruleID string) *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete feg network ID subscriber config rule names rule ID params
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFegNetworkIDSubscriberConfigRuleNamesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
