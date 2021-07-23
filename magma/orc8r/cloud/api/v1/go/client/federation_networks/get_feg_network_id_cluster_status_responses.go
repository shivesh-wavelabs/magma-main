// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetFegNetworkIDClusterStatusReader is a Reader for the GetFegNetworkIDClusterStatus structure.
type GetFegNetworkIDClusterStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegNetworkIDClusterStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegNetworkIDClusterStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegNetworkIDClusterStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegNetworkIDClusterStatusOK creates a GetFegNetworkIDClusterStatusOK with default headers values
func NewGetFegNetworkIDClusterStatusOK() *GetFegNetworkIDClusterStatusOK {
	return &GetFegNetworkIDClusterStatusOK{}
}

/*GetFegNetworkIDClusterStatusOK handles this case with default header values.

Cluster status of Federation Network
*/
type GetFegNetworkIDClusterStatusOK struct {
	Payload *models.FederationNetworkClusterStatus
}

func (o *GetFegNetworkIDClusterStatusOK) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/cluster_status][%d] getFegNetworkIdClusterStatusOK  %+v", 200, o.Payload)
}

func (o *GetFegNetworkIDClusterStatusOK) GetPayload() *models.FederationNetworkClusterStatus {
	return o.Payload
}

func (o *GetFegNetworkIDClusterStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FederationNetworkClusterStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegNetworkIDClusterStatusDefault creates a GetFegNetworkIDClusterStatusDefault with default headers values
func NewGetFegNetworkIDClusterStatusDefault(code int) *GetFegNetworkIDClusterStatusDefault {
	return &GetFegNetworkIDClusterStatusDefault{
		_statusCode: code,
	}
}

/*GetFegNetworkIDClusterStatusDefault handles this case with default header values.

Unexpected Error
*/
type GetFegNetworkIDClusterStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg network ID cluster status default response
func (o *GetFegNetworkIDClusterStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetFegNetworkIDClusterStatusDefault) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/cluster_status][%d] GetFegNetworkIDClusterStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegNetworkIDClusterStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegNetworkIDClusterStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
