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

// GetLTENetworkIDReader is a Reader for the GetLTENetworkID structure.
type GetLTENetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTENetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTENetworkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTENetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTENetworkIDOK creates a GetLTENetworkIDOK with default headers values
func NewGetLTENetworkIDOK() *GetLTENetworkIDOK {
	return &GetLTENetworkIDOK{}
}

/*GetLTENetworkIDOK handles this case with default header values.

Full description of an LTE network
*/
type GetLTENetworkIDOK struct {
	Payload *models.LTENetwork
}

func (o *GetLTENetworkIDOK) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}][%d] getLteNetworkIdOK  %+v", 200, o.Payload)
}

func (o *GetLTENetworkIDOK) GetPayload() *models.LTENetwork {
	return o.Payload
}

func (o *GetLTENetworkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LTENetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTENetworkIDDefault creates a GetLTENetworkIDDefault with default headers values
func NewGetLTENetworkIDDefault(code int) *GetLTENetworkIDDefault {
	return &GetLTENetworkIDDefault{
		_statusCode: code,
	}
}

/*GetLTENetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type GetLTENetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE network ID default response
func (o *GetLTENetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *GetLTENetworkIDDefault) Error() string {
	return fmt.Sprintf("[GET /lte/{network_id}][%d] GetLTENetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTENetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTENetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
