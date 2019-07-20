// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIShapeReader is a Reader for the THSRAPIShape structure.
type THSRAPIShapeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIShapeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIShapeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTHSRAPIShapeOK creates a THSRAPIShapeOK with default headers values
func NewTHSRAPIShapeOK() *THSRAPIShapeOK {
	return &THSRAPIShapeOK{}
}

/*THSRAPIShapeOK handles this case with default header values.

OK
*/
type THSRAPIShapeOK struct {
	Payload []*models.ServiceDTOVersion2RailTHSRTHSRShape
}

func (o *THSRAPIShapeOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/Shape][%d] tHSRApiShapeOK  %+v", 200, o.Payload)
}

func (o *THSRAPIShapeOK) GetPayload() []*models.ServiceDTOVersion2RailTHSRTHSRShape {
	return o.Payload
}

func (o *THSRAPIShapeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
