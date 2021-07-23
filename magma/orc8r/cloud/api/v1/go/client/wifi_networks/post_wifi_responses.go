// Code generated by go-swagger; DO NOT EDIT.

package wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostWifiReader is a Reader for the PostWifi structure.
type PostWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWifiCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWifiCreated creates a PostWifiCreated with default headers values
func NewPostWifiCreated() *PostWifiCreated {
	return &PostWifiCreated{}
}

/*PostWifiCreated handles this case with default header values.

Success
*/
type PostWifiCreated struct {
}

func (o *PostWifiCreated) Error() string {
	return fmt.Sprintf("[POST /wifi][%d] postWifiCreated ", 201)
}

func (o *PostWifiCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWifiDefault creates a PostWifiDefault with default headers values
func NewPostWifiDefault(code int) *PostWifiDefault {
	return &PostWifiDefault{
		_statusCode: code,
	}
}

/*PostWifiDefault handles this case with default header values.

Unexpected Error
*/
type PostWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post wifi default response
func (o *PostWifiDefault) Code() int {
	return o._statusCode
}

func (o *PostWifiDefault) Error() string {
	return fmt.Sprintf("[POST /wifi][%d] PostWifi default  %+v", o._statusCode, o.Payload)
}

func (o *PostWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
