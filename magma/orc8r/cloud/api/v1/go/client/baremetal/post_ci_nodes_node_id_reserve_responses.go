// Code generated by go-swagger; DO NOT EDIT.

package baremetal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostCiNodesNodeIDReserveReader is a Reader for the PostCiNodesNodeIDReserve structure.
type PostCiNodesNodeIDReserveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCiNodesNodeIDReserveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCiNodesNodeIDReserveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCiNodesNodeIDReserveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCiNodesNodeIDReserveOK creates a PostCiNodesNodeIDReserveOK with default headers values
func NewPostCiNodesNodeIDReserveOK() *PostCiNodesNodeIDReserveOK {
	return &PostCiNodesNodeIDReserveOK{}
}

/*PostCiNodesNodeIDReserveOK handles this case with default header values.

Lease for the requested node
*/
type PostCiNodesNodeIDReserveOK struct {
	Payload *models.NodeLease
}

func (o *PostCiNodesNodeIDReserveOK) Error() string {
	return fmt.Sprintf("[POST /ci/nodes/{node_id}/reserve][%d] postCiNodesNodeIdReserveOK  %+v", 200, o.Payload)
}

func (o *PostCiNodesNodeIDReserveOK) GetPayload() *models.NodeLease {
	return o.Payload
}

func (o *PostCiNodesNodeIDReserveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeLease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCiNodesNodeIDReserveDefault creates a PostCiNodesNodeIDReserveDefault with default headers values
func NewPostCiNodesNodeIDReserveDefault(code int) *PostCiNodesNodeIDReserveDefault {
	return &PostCiNodesNodeIDReserveDefault{
		_statusCode: code,
	}
}

/*PostCiNodesNodeIDReserveDefault handles this case with default header values.

Unexpected Error
*/
type PostCiNodesNodeIDReserveDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post ci nodes node ID reserve default response
func (o *PostCiNodesNodeIDReserveDefault) Code() int {
	return o._statusCode
}

func (o *PostCiNodesNodeIDReserveDefault) Error() string {
	return fmt.Sprintf("[POST /ci/nodes/{node_id}/reserve][%d] PostCiNodesNodeIDReserve default  %+v", o._statusCode, o.Payload)
}

func (o *PostCiNodesNodeIDReserveDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostCiNodesNodeIDReserveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
