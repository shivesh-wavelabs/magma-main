// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDHaPairsHaPairIDReader is a Reader for the GetCwfNetworkIDHaPairsHaPairID structure.
type GetCwfNetworkIDHaPairsHaPairIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDHaPairsHaPairIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDHaPairsHaPairIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDHaPairsHaPairIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDHaPairsHaPairIDOK creates a GetCwfNetworkIDHaPairsHaPairIDOK with default headers values
func NewGetCwfNetworkIDHaPairsHaPairIDOK() *GetCwfNetworkIDHaPairsHaPairIDOK {
	return &GetCwfNetworkIDHaPairsHaPairIDOK{}
}

/*GetCwfNetworkIDHaPairsHaPairIDOK handles this case with default header values.

Retrieve high availability gateway pair for Carrier Wifi network
*/
type GetCwfNetworkIDHaPairsHaPairIDOK struct {
	Payload *models.CwfHaPair
}

func (o *GetCwfNetworkIDHaPairsHaPairIDOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/ha_pairs/{ha_pair_id}][%d] getCwfNetworkIdHaPairsHaPairIdOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDHaPairsHaPairIDOK) GetPayload() *models.CwfHaPair {
	return o.Payload
}

func (o *GetCwfNetworkIDHaPairsHaPairIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CwfHaPair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDHaPairsHaPairIDDefault creates a GetCwfNetworkIDHaPairsHaPairIDDefault with default headers values
func NewGetCwfNetworkIDHaPairsHaPairIDDefault(code int) *GetCwfNetworkIDHaPairsHaPairIDDefault {
	return &GetCwfNetworkIDHaPairsHaPairIDDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDHaPairsHaPairIDDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDHaPairsHaPairIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID ha pairs ha pair ID default response
func (o *GetCwfNetworkIDHaPairsHaPairIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDHaPairsHaPairIDDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/ha_pairs/{ha_pair_id}][%d] GetCwfNetworkIDHaPairsHaPairID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDHaPairsHaPairIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDHaPairsHaPairIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
