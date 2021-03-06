// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDGatewaysGatewayIDCellularDNSReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDCellularDNS structure.
type PutLTENetworkIDGatewaysGatewayIDCellularDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDCellularDNSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent creates a PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent() *PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/dns][%d] putLteNetworkIdGatewaysGatewayIdCellularDnsNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDCellularDNSDefault creates a PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDCellularDNSDefault(code int) *PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault {
	return &PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID cellular DNS default response
func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/cellular/dns][%d] PutLTENetworkIDGatewaysGatewayIDCellularDNS default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDCellularDNSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
