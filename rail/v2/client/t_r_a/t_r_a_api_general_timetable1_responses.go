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

// TRAAPIGeneralTimetable1Reader is a Reader for the TRAAPIGeneralTimetable1 structure.
type TRAAPIGeneralTimetable1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIGeneralTimetable1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIGeneralTimetable1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIGeneralTimetable1OK creates a TRAAPIGeneralTimetable1OK with default headers values
func NewTRAAPIGeneralTimetable1OK() *TRAAPIGeneralTimetable1OK {
	return &TRAAPIGeneralTimetable1OK{}
}

/*TRAAPIGeneralTimetable1OK handles this case with default header values.

Success
*/
type TRAAPIGeneralTimetable1OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailGeneralTimetable
}

func (o *TRAAPIGeneralTimetable1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable/TrainNo/{TrainNo}][%d] tRAApiGeneralTimetable1OK  %+v", 200, o.Payload)
}

func (o *TRAAPIGeneralTimetable1OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailGeneralTimetable {
	return o.Payload
}

func (o *TRAAPIGeneralTimetable1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
