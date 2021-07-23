// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutCwfNetworkIDGatewaysGatewayIDCarrierWifiReader is a Reader for the PutCwfNetworkIDGatewaysGatewayIDCarrierWifi structure.
type PutCwfNetworkIDGatewaysGatewayIDCarrierWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent creates a PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent() *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent {
	return &PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent{}
}

/*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent handles this case with default header values.

Success
*/
type PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent struct {
}

func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/carrier_wifi][%d] putCwfNetworkIdGatewaysGatewayIdCarrierWifiNoContent ", 204)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault creates a PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault with default headers values
func NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault(code int) *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault {
	return &PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault{
		_statusCode: code,
	}
}

/*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault handles this case with default header values.

Unexpected Error
*/
type PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put cwf network ID gateways gateway ID carrier wifi default response
func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault) Code() int {
	return o._statusCode
}

func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault) Error() string {
	return fmt.Sprintf("[PUT /cwf/{network_id}/gateways/{gateway_id}/carrier_wifi][%d] PutCwfNetworkIDGatewaysGatewayIDCarrierWifi default  %+v", o._statusCode, o.Payload)
}

func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
