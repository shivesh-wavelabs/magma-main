// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDGatewaysGatewayIDTierReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayIDTier structure.
type GetNetworksNetworkIDGatewaysGatewayIDTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDTierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDTierOK creates a GetNetworksNetworkIDGatewaysGatewayIDTierOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDTierOK() *GetNetworksNetworkIDGatewaysGatewayIDTierOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDTierOK{}
}

/*GetNetworksNetworkIDGatewaysGatewayIDTierOK handles this case with default header values.

The ID of the upgrade tier
*/
type GetNetworksNetworkIDGatewaysGatewayIDTierOK struct {
	Payload models.TierID
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDTierOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/tier][%d] getNetworksNetworkIdGatewaysGatewayIdTierOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDTierOK) GetPayload() models.TierID {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDTierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDTierDefault creates a GetNetworksNetworkIDGatewaysGatewayIDTierDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDTierDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDTierDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDTierDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDTierDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID tier default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDTierDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDTierDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}/tier][%d] GetNetworksNetworkIDGatewaysGatewayIDTier default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDTierDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
