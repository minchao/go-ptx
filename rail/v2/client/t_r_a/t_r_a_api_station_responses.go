// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// TRAAPIStationReader is a Reader for the TRAAPIStation structure.
type TRAAPIStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIStationOK creates a TRAAPIStationOK with default headers values
func NewTRAAPIStationOK() *TRAAPIStationOK {
	return &TRAAPIStationOK{}
}

/*TRAAPIStationOK handles this case with default header values.

Success
*/
type TRAAPIStationOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailStation
}

func (o *TRAAPIStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/Station][%d] tRAApiStationOK  %+v", 200, o.Payload)
}

func (o *TRAAPIStationOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailStation {
	return o.Payload
}

func (o *TRAAPIStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
