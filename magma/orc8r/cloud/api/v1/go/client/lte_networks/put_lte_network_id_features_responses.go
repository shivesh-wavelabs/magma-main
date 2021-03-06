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

// PutLTENetworkIDFeaturesReader is a Reader for the PutLTENetworkIDFeatures structure.
type PutLTENetworkIDFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDFeaturesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDFeaturesNoContent creates a PutLTENetworkIDFeaturesNoContent with default headers values
func NewPutLTENetworkIDFeaturesNoContent() *PutLTENetworkIDFeaturesNoContent {
	return &PutLTENetworkIDFeaturesNoContent{}
}

/*PutLTENetworkIDFeaturesNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDFeaturesNoContent struct {
}

func (o *PutLTENetworkIDFeaturesNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/features][%d] putLteNetworkIdFeaturesNoContent ", 204)
}

func (o *PutLTENetworkIDFeaturesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDFeaturesDefault creates a PutLTENetworkIDFeaturesDefault with default headers values
func NewPutLTENetworkIDFeaturesDefault(code int) *PutLTENetworkIDFeaturesDefault {
	return &PutLTENetworkIDFeaturesDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDFeaturesDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID features default response
func (o *PutLTENetworkIDFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDFeaturesDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/features][%d] PutLTENetworkIDFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDFeaturesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
