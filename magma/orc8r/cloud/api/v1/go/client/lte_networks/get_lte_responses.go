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

// GetLTEReader is a Reader for the GetLTE structure.
type GetLTEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLTEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLTEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLTEDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLTEOK creates a GetLTEOK with default headers values
func NewGetLTEOK() *GetLTEOK {
	return &GetLTEOK{}
}

/*GetLTEOK handles this case with default header values.

List of LTE network IDs
*/
type GetLTEOK struct {
	Payload []string
}

func (o *GetLTEOK) Error() string {
	return fmt.Sprintf("[GET /lte][%d] getLteOK  %+v", 200, o.Payload)
}

func (o *GetLTEOK) GetPayload() []string {
	return o.Payload
}

func (o *GetLTEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLTEDefault creates a GetLTEDefault with default headers values
func NewGetLTEDefault(code int) *GetLTEDefault {
	return &GetLTEDefault{
		_statusCode: code,
	}
}

/*GetLTEDefault handles this case with default header values.

Unexpected Error
*/
type GetLTEDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get LTE default response
func (o *GetLTEDefault) Code() int {
	return o._statusCode
}

func (o *GetLTEDefault) Error() string {
	return fmt.Sprintf("[GET /lte][%d] GetLTE default  %+v", o._statusCode, o.Payload)
}

func (o *GetLTEDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLTEDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
