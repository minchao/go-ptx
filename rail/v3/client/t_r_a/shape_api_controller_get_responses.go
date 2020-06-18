// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// ShapeAPIControllerGetReader is a Reader for the ShapeAPIControllerGet structure.
type ShapeAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShapeAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShapeAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewShapeAPIControllerGetOK creates a ShapeAPIControllerGetOK with default headers values
func NewShapeAPIControllerGetOK() *ShapeAPIControllerGetOK {
	return &ShapeAPIControllerGetOK{}
}

/*ShapeAPIControllerGetOK handles this case with default header values.

OK
*/
type ShapeAPIControllerGetOK struct {
	Payload *models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAShape
}

func (o *ShapeAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Shape][%d] shapeApiControllerGetOK  %+v", 200, o.Payload)
}

func (o *ShapeAPIControllerGetOK) GetPayload() *models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAShape {
	return o.Payload
}

func (o *ShapeAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAShape)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
