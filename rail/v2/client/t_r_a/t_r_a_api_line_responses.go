// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// TRAAPILineReader is a Reader for the TRAAPILine structure.
type TRAAPILineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPILineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPILineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTRAAPILineOK creates a TRAAPILineOK with default headers values
func NewTRAAPILineOK() *TRAAPILineOK {
	return &TRAAPILineOK{}
}

/*TRAAPILineOK handles this case with default header values.

OK
*/
type TRAAPILineOK struct {
	Payload []*models.ServiceDTOVersion2RailTRALine
}

func (o *TRAAPILineOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/Line][%d] tRAApiLineOK  %+v", 200, o.Payload)
}

func (o *TRAAPILineOK) GetPayload() []*models.ServiceDTOVersion2RailTRALine {
	return o.Payload
}

func (o *TRAAPILineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
