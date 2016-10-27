package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aws/amazon-ecs-event-stream-handler/internal/models"
)

// GetInstanceReader is a Reader for the GetInstance structure.
type GetInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetInstanceOK creates a GetInstanceOK with default headers values
func NewGetInstanceOK() *GetInstanceOK {
	return &GetInstanceOK{}
}

/*GetInstanceOK handles this case with default header values.

Get instance by ARN response
*/
type GetInstanceOK struct {
	Payload *models.ContainerInstanceModel
}

func (o *GetInstanceOK) Error() string {
	return fmt.Sprintf("[GET /instance/{arn}][%d] getInstanceOK  %+v", 200, o.Payload)
}

func (o *GetInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerInstanceModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceDefault creates a GetInstanceDefault with default headers values
func NewGetInstanceDefault(code int) *GetInstanceDefault {
	return &GetInstanceDefault{
		_statusCode: code,
	}
}

/*GetInstanceDefault handles this case with default header values.

Unexpected error getting instance by ARN
*/
type GetInstanceDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the get instance default response
func (o *GetInstanceDefault) Code() int {
	return o._statusCode
}

func (o *GetInstanceDefault) Error() string {
	return fmt.Sprintf("[GET /instance/{arn}][%d] GetInstance default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}