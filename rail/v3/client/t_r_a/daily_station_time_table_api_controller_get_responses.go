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

// DailyStationTimeTableAPIControllerGetReader is a Reader for the DailyStationTimeTableAPIControllerGet structure.
type DailyStationTimeTableAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyStationTimeTableAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyStationTimeTableAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDailyStationTimeTableAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDailyStationTimeTableAPIControllerGetOK creates a DailyStationTimeTableAPIControllerGetOK with default headers values
func NewDailyStationTimeTableAPIControllerGetOK() *DailyStationTimeTableAPIControllerGetOK {
	return &DailyStationTimeTableAPIControllerGetOK{}
}

/* DailyStationTimeTableAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type DailyStationTimeTableAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable
}

func (o *DailyStationTimeTableAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/Today][%d] dailyStationTimeTableApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *DailyStationTimeTableAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable {
	return o.Payload
}

func (o *DailyStationTimeTableAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyStationTimeTableAPIControllerGetNotModified creates a DailyStationTimeTableAPIControllerGetNotModified with default headers values
func NewDailyStationTimeTableAPIControllerGetNotModified() *DailyStationTimeTableAPIControllerGetNotModified {
	return &DailyStationTimeTableAPIControllerGetNotModified{}
}

/* DailyStationTimeTableAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type DailyStationTimeTableAPIControllerGetNotModified struct {
}

func (o *DailyStationTimeTableAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/Today][%d] dailyStationTimeTableApiControllerGetNotModified ", 304)
}

func (o *DailyStationTimeTableAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
