// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetFegNetworkIDGatewaysReader is a Reader for the GetFegNetworkIDGateways structure.
type GetFegNetworkIDGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegNetworkIDGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegNetworkIDGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegNetworkIDGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegNetworkIDGatewaysOK creates a GetFegNetworkIDGatewaysOK with default headers values
func NewGetFegNetworkIDGatewaysOK() *GetFegNetworkIDGatewaysOK {
	return &GetFegNetworkIDGatewaysOK{}
}

/*GetFegNetworkIDGatewaysOK handles this case with default header values.

Map of all federated gateways inside the network by gatewayID
*/
type GetFegNetworkIDGatewaysOK struct {
	Payload map[string]models.FederationGateway
}

func (o *GetFegNetworkIDGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/gateways][%d] getFegNetworkIdGatewaysOK  %+v", 200, o.Payload)
}

func (o *GetFegNetworkIDGatewaysOK) GetPayload() map[string]models.FederationGateway {
	return o.Payload
}

func (o *GetFegNetworkIDGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegNetworkIDGatewaysDefault creates a GetFegNetworkIDGatewaysDefault with default headers values
func NewGetFegNetworkIDGatewaysDefault(code int) *GetFegNetworkIDGatewaysDefault {
	return &GetFegNetworkIDGatewaysDefault{
		_statusCode: code,
	}
}

/*GetFegNetworkIDGatewaysDefault handles this case with default header values.

Unexpected Error
*/
type GetFegNetworkIDGatewaysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg network ID gateways default response
func (o *GetFegNetworkIDGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *GetFegNetworkIDGatewaysDefault) Error() string {
	return fmt.Sprintf("[GET /feg/{network_id}/gateways][%d] GetFegNetworkIDGateways default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegNetworkIDGatewaysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegNetworkIDGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
