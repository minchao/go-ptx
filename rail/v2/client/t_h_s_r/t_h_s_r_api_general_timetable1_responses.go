// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIGeneralTimetable1Reader is a Reader for the THSRAPIGeneralTimetable1 structure.
type THSRAPIGeneralTimetable1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIGeneralTimetable1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIGeneralTimetable1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIGeneralTimetable1OK creates a THSRAPIGeneralTimetable1OK with default headers values
func NewTHSRAPIGeneralTimetable1OK() *THSRAPIGeneralTimetable1OK {
	return &THSRAPIGeneralTimetable1OK{}
}

/*THSRAPIGeneralTimetable1OK handles this case with default header values.

Success
*/
type THSRAPIGeneralTimetable1OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailGeneralTimetable
}

func (o *THSRAPIGeneralTimetable1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/GeneralTimetable/TrainNo/{TrainNo}][%d] tHSRApiGeneralTimetable1OK  %+v", 200, o.Payload)
}

func (o *THSRAPIGeneralTimetable1OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailGeneralTimetable {
	return o.Payload
}

func (o *THSRAPIGeneralTimetable1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
