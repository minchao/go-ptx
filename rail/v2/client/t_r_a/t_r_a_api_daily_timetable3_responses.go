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

// TRAAPIDailyTimetable3Reader is a Reader for the TRAAPIDailyTimetable3 structure.
type TRAAPIDailyTimetable3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIDailyTimetable3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIDailyTimetable3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIDailyTimetable3OK creates a TRAAPIDailyTimetable3OK with default headers values
func NewTRAAPIDailyTimetable3OK() *TRAAPIDailyTimetable3OK {
	return &TRAAPIDailyTimetable3OK{}
}

/*TRAAPIDailyTimetable3OK handles this case with default header values.

Success
*/
type TRAAPIDailyTimetable3OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailDailyTimetable
}

func (o *TRAAPIDailyTimetable3OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/TrainNo/{TrainNo}/TrainDate/{TrainDate}][%d] tRAApiDailyTimetable3OK  %+v", 200, o.Payload)
}

func (o *TRAAPIDailyTimetable3OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailDailyTimetable {
	return o.Payload
}

func (o *TRAAPIDailyTimetable3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
