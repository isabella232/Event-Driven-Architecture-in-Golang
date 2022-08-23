// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"eda-in-golang/stores/storesclient/models"
)

// RemoveProductReader is a Reader for the RemoveProduct structure.
type RemoveProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveProductOK creates a RemoveProductOK with default headers values
func NewRemoveProductOK() *RemoveProductOK {
	return &RemoveProductOK{}
}

/* RemoveProductOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveProductOK struct {
	Payload models.StorespbRemoveProductResponse
}

func (o *RemoveProductOK) Error() string {
	return fmt.Sprintf("[DELETE /api/stores/products/{id}][%d] removeProductOK  %+v", 200, o.Payload)
}
func (o *RemoveProductOK) GetPayload() models.StorespbRemoveProductResponse {
	return o.Payload
}

func (o *RemoveProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProductDefault creates a RemoveProductDefault with default headers values
func NewRemoveProductDefault(code int) *RemoveProductDefault {
	return &RemoveProductDefault{
		_statusCode: code,
	}
}

/* RemoveProductDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveProductDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the remove product default response
func (o *RemoveProductDefault) Code() int {
	return o._statusCode
}

func (o *RemoveProductDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/stores/products/{id}][%d] removeProduct default  %+v", o._statusCode, o.Payload)
}
func (o *RemoveProductDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RemoveProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}