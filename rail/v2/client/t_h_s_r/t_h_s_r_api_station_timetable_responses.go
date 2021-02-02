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
	case 304:
		result := NewTHSRAPIStationTimetableNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIStationTimetableOK creates a THSRAPIStationTimetableOK with default headers values
func NewTHSRAPIStationTimetableOK() *THSRAPIStationTimetableOK {
	return &THSRAPIStationTimetableOK{}
}

/* THSRAPIStationTimetableOK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIStationTimetableOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailStationTimetable
}

func (o *THSRAPIStationTimetableOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tHSRApiStationTimetableOK  %+v", 200, o.Payload)
}
func (o *THSRAPIStationTimetableOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailStationTimetable {
	return o.Payload
}

func (o *THSRAPIStationTimetableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIStationTimetableNotModified creates a THSRAPIStationTimetableNotModified with default headers values
func NewTHSRAPIStationTimetableNotModified() *THSRAPIStationTimetableNotModified {
	return &THSRAPIStationTimetableNotModified{}
}

/* THSRAPIStationTimetableNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIStationTimetableNotModified struct {
}

func (o *THSRAPIStationTimetableNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tHSRApiStationTimetableNotModified ", 304)
}

func (o *THSRAPIStationTimetableNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
