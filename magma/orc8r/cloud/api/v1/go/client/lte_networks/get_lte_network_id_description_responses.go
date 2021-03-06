// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetLTENetworkIDDescriptionReader is a Reader for the GetLTENetworkIDDescription structure.
type GetLTENetworkIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDDescriptionOK creates a GetLTENetworkIDDescriptionOK with default headers values
func NewGetLTENetworkIDDescriptionOK() *GetLTENetworkIDDescriptionOK {
	return &GetLTENetworkIDDescriptionOK{}
}

/*GetLTENetworkIDDescriptionOK handles this case with default header values.

Description of the network
*/
type GetLTENetworkIDDescriptionOK struct {
	Payload models.NetworkDescription
}

func (o *GetLTENetworkIDDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/description][%d] getLteNetworkIdDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDDescriptionOK) GetPayload() models.NetworkDescription {
	return o.Payload
}

func (o *GetLTENetworkIDDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDDescriptionDefault creates a GetLTENetworkIDDescriptionDefault with default headers values
func NewGetLTENetworkIDDescriptionDefault(code int) *GetLTENetworkIDDescriptionDefault {
	return &GetLTENetworkIDDescriptionDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID description default response
func (o *GetLTENetworkIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}/description][%d] GetLTENetworkIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
