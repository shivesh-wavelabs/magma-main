// Code generated by go-swagger; DO NOT EDIT.

package wifi_meshes

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

// NewPutWifiNetworkIDMeshesMeshIDNameParams creates a new PutWifiNetworkIDMeshesMeshIDNameParams object
// with the default values initialized.
func NewPutWifiNetworkIDMeshesMeshIDNameParams() *PutWifiNetworkIDMeshesMeshIDNameParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutWifiNetworkIDMeshesMeshIDNameParamsWithTimeout creates a new PutWifiNetworkIDMeshesMeshIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutWifiNetworkIDMeshesMeshIDNameParamsWithTimeout(timeout time.Duration) *PutWifiNetworkIDMeshesMeshIDNameParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDNameParams{

		timeout: timeout,
	}
}

// NewPutWifiNetworkIDMeshesMeshIDNameParamsWithContext creates a new PutWifiNetworkIDMeshesMeshIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutWifiNetworkIDMeshesMeshIDNameParamsWithContext(ctx context.Context) *PutWifiNetworkIDMeshesMeshIDNameParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDNameParams{

		Context: ctx,
	}
}

// NewPutWifiNetworkIDMeshesMeshIDNameParamsWithHTTPClient creates a new PutWifiNetworkIDMeshesMeshIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutWifiNetworkIDMeshesMeshIDNameParamsWithHTTPClient(client *http.Client) *PutWifiNetworkIDMeshesMeshIDNameParams {
	var ()
	return &PutWifiNetworkIDMeshesMeshIDNameParams{
		HTTPClient: client,
	}
}

/*PutWifiNetworkIDMeshesMeshIDNameParams contains all the parameters to send to the API endpoint
for the put wifi network ID meshes mesh ID name operation typically these are written to a http.Request
*/
type PutWifiNetworkIDMeshesMeshIDNameParams struct {

	/*MeshID
	  Mesh ID

	*/
	MeshID string
	/*MeshName
	  The updated name of the mesh

	*/
	MeshName models.MeshName
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WithTimeout(timeout time.Duration) *PutWifiNetworkIDMeshesMeshIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WithContext(ctx context.Context) *PutWifiNetworkIDMeshesMeshIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WithHTTPClient(client *http.Client) *PutWifiNetworkIDMeshesMeshIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeshID adds the meshID to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WithMeshID(meshID string) *PutWifiNetworkIDMeshesMeshIDNameParams {
	o.SetMeshID(meshID)
	return o
}

// SetMeshID adds the meshId to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) SetMeshID(meshID string) {
	o.MeshID = meshID
}

// WithMeshName adds the meshName to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WithMeshName(meshName models.MeshName) *PutWifiNetworkIDMeshesMeshIDNameParams {
	o.SetMeshName(meshName)
	return o
}

// SetMeshName adds the meshName to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) SetMeshName(meshName models.MeshName) {
	o.MeshName = meshName
}

// WithNetworkID adds the networkID to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WithNetworkID(networkID string) *PutWifiNetworkIDMeshesMeshIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put wifi network ID meshes mesh ID name params
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutWifiNetworkIDMeshesMeshIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mesh_id
	if err := r.SetPathParam("mesh_id", o.MeshID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.MeshName); err != nil {
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
