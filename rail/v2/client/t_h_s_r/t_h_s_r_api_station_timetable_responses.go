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

// THSRAPIStationTimetableReader is a Reader for the THSRAPIStationTimetable structure.
type THSRAPIStationTimetableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIStationTimetableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIStationTimetableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIStationTimetableOK creates a THSRAPIStationTimetableOK with default headers values
func NewTHSRAPIStationTimetableOK() *THSRAPIStationTimetableOK {
	return &THSRAPIStationTimetableOK{}
}

/*THSRAPIStationTimetableOK handles this case with default header values.

OK
*/
type THSRAPIStationTimetableOK struct {
	Payload []*models.ServiceDTOVersion2RailTHSRRailStationTimetable
}

func (o *THSRAPIStationTimetableOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tHSRApiStationTimetableOK  %+v", 200, o.Payload)
}

func (o *THSRAPIStationTimetableOK) GetPayload() []*models.ServiceDTOVersion2RailTHSRRailStationTimetable {
	return o.Payload
}

func (o *THSRAPIStationTimetableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
