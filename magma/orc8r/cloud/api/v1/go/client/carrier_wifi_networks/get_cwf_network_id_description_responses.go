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

// GetCwfNetworkIDDescriptionReader is a Reader for the GetCwfNetworkIDDescription structure.
type GetCwfNetworkIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDDescriptionOK creates a GetCwfNetworkIDDescriptionOK with default headers values
func NewGetCwfNetworkIDDescriptionOK() *GetCwfNetworkIDDescriptionOK {
	return &GetCwfNetworkIDDescriptionOK{}
}

/*GetCwfNetworkIDDescriptionOK handles this case with default header values.

Description of the network
*/
type GetCwfNetworkIDDescriptionOK struct {
	Payload models.NetworkDescription
}

func (o *GetCwfNetworkIDDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/description][%d] getCwfNetworkIdDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDDescriptionOK) GetPayload() models.NetworkDescription {
	return o.Payload
}

func (o *GetCwfNetworkIDDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDDescriptionDefault creates a GetCwfNetworkIDDescriptionDefault with default headers values
func NewGetCwfNetworkIDDescriptionDefault(code int) *GetCwfNetworkIDDescriptionDefault {
	return &GetCwfNetworkIDDescriptionDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID description default response
func (o *GetCwfNetworkIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/description][%d] GetCwfNetworkIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
