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

// PutLTENetworkIDSubscriberConfigRuleNamesReader is a Reader for the PutLTENetworkIDSubscriberConfigRuleNames structure.
type PutLTENetworkIDSubscriberConfigRuleNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDSubscriberConfigRuleNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDSubscriberConfigRuleNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDSubscriberConfigRuleNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDSubscriberConfigRuleNamesNoContent creates a PutLTENetworkIDSubscriberConfigRuleNamesNoContent with default headers values
func NewPutLTENetworkIDSubscriberConfigRuleNamesNoContent() *PutLTENetworkIDSubscriberConfigRuleNamesNoContent {
	return &PutLTENetworkIDSubscriberConfigRuleNamesNoContent{}
}

/*PutLTENetworkIDSubscriberConfigRuleNamesNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDSubscriberConfigRuleNamesNoContent struct {
}

func (o *PutLTENetworkIDSubscriberConfigRuleNamesNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/subscriber_config/rule_names][%d] putLteNetworkIdSubscriberConfigRuleNamesNoContent ", 204)
}

func (o *PutLTENetworkIDSubscriberConfigRuleNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDSubscriberConfigRuleNamesDefault creates a PutLTENetworkIDSubscriberConfigRuleNamesDefault with default headers values
func NewPutLTENetworkIDSubscriberConfigRuleNamesDefault(code int) *PutLTENetworkIDSubscriberConfigRuleNamesDefault {
	return &PutLTENetworkIDSubscriberConfigRuleNamesDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDSubscriberConfigRuleNamesDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDSubscriberConfigRuleNamesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID subscriber config rule names default response
func (o *PutLTENetworkIDSubscriberConfigRuleNamesDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDSubscriberConfigRuleNamesDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/subscriber_config/rule_names][%d] PutLTENetworkIDSubscriberConfigRuleNames default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDSubscriberConfigRuleNamesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDSubscriberConfigRuleNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
