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

// PutFegNetworkIDReader is a Reader for the PutFegNetworkID structure.
type PutFegNetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFegNetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutFegNetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutFegNetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutFegNetworkIDNoContent creates a PutFegNetworkIDNoContent with default headers values
func NewPutFegNetworkIDNoContent() *PutFegNetworkIDNoContent {
	return &PutFegNetworkIDNoContent{}
}

/*PutFegNetworkIDNoContent handles this case with default header values.

Success
*/
type PutFegNetworkIDNoContent struct {
}

func (o *PutFegNetworkIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /feg/{network_id}][%d] putFegNetworkIdNoContent ", 204)
}

func (o *PutFegNetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFegNetworkIDDefault creates a PutFegNetworkIDDefault with default headers values
func NewPutFegNetworkIDDefault(code int) *PutFegNetworkIDDefault {
	return &PutFegNetworkIDDefault{
		_statusCode: code,
	}
}

/*PutFegNetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type PutFegNetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put feg network ID default response
func (o *PutFegNetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *PutFegNetworkIDDefault) Error() string {
	return fmt.Sprintf("[PUT /feg/{network_id}][%d] PutFegNetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *PutFegNetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutFegNetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
