// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/models"
)

// LineAPIControllerGetReader is a Reader for the LineAPIControllerGet structure.
type LineAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LineAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLineAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLineAPIControllerGetOK creates a LineAPIControllerGetOK with default headers values
func NewLineAPIControllerGetOK() *LineAPIControllerGetOK {
	return &LineAPIControllerGetOK{}
}

/*LineAPIControllerGetOK handles this case with default header values.

OK
*/
type LineAPIControllerGetOK struct {
	Payload *models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRALine
}

func (o *LineAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/Line][%d] lineApiControllerGetOK  %+v", 200, o.Payload)
}

func (o *LineAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRALine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
