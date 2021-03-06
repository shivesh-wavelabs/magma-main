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

// DeleteCwfNetworkIDReader is a Reader for the DeleteCwfNetworkID structure.
type DeleteCwfNetworkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCwfNetworkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCwfNetworkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCwfNetworkIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCwfNetworkIDNoContent creates a DeleteCwfNetworkIDNoContent with default headers values
func NewDeleteCwfNetworkIDNoContent() *DeleteCwfNetworkIDNoContent {
	return &DeleteCwfNetworkIDNoContent{}
}

/*DeleteCwfNetworkIDNoContent handles this case with default header values.

Success
*/
type DeleteCwfNetworkIDNoContent struct {
}

func (o *DeleteCwfNetworkIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}][%d] deleteCwfNetworkIdNoContent ", 204)
}

func (o *DeleteCwfNetworkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCwfNetworkIDDefault creates a DeleteCwfNetworkIDDefault with default headers values
func NewDeleteCwfNetworkIDDefault(code int) *DeleteCwfNetworkIDDefault {
	return &DeleteCwfNetworkIDDefault{
		_statusCode: code,
	}
}

/*DeleteCwfNetworkIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteCwfNetworkIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete cwf network ID default response
func (o *DeleteCwfNetworkIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCwfNetworkIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cwf/{network_id}][%d] DeleteCwfNetworkID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCwfNetworkIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCwfNetworkIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
