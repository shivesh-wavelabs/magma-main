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

// DeleteLTENetworkIDReader is a Reader for the DeleteLTENetworkID structure.
type DeleteLTENetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDNoContent creates a DeleteLTENetworkIDNoContent with default headers values
func NewDeleteLTENetworkIDNoContent() *DeleteLTENetworkIDNoContent {
	return &DeleteLTENetworkIDNoContent{}
}

/*DeleteLTENetworkIDNoContent handles this case with default header values.

Success
*/
type DeleteLTENetworkIDNoContent struct {
}

func (o *DeleteLTENetworkIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}][%d] deleteLteNetworkIdNoContent ", 204)
}

func (o *DeleteLTENetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDDefault creates a DeleteLTENetworkIDDefault with default headers values
func NewDeleteLTENetworkIDDefault(code int) *DeleteLTENetworkIDDefault {
	return &DeleteLTENetworkIDDefault{
		_statusCode: code,
	}
}

/*DeleteLTENetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID default response
func (o *DeleteLTENetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}][%d] DeleteLTENetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLTENetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
