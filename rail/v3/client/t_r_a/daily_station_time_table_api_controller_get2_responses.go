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

// DailyStationTimeTableAPIControllerGet2Reader is a Reader for the DailyStationTimeTableAPIControllerGet2 structure.
type DailyStationTimeTableAPIControllerGet2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DailyStationTimeTableAPIControllerGet2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDailyStationTimeTableAPIControllerGet2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewDailyStationTimeTableAPIControllerGet2Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDailyStationTimeTableAPIControllerGet2NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDailyStationTimeTableAPIControllerGet2OK creates a DailyStationTimeTableAPIControllerGet2OK with default headers values
func NewDailyStationTimeTableAPIControllerGet2OK() *DailyStationTimeTableAPIControllerGet2OK {
	return &DailyStationTimeTableAPIControllerGet2OK{}
}

/* DailyStationTimeTableAPIControllerGet2OK describes a response with status code 200, with default header values.

Success
*/
type DailyStationTimeTableAPIControllerGet2OK struct {
	Payload *models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable
}

func (o *DailyStationTimeTableAPIControllerGet2OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/TrainDate/{TrainDate}][%d] dailyStationTimeTableApiControllerGet2OK  %+v", 200, o.Payload)
}
func (o *DailyStationTimeTableAPIControllerGet2OK) GetPayload() *models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable {
	return o.Payload
}

func (o *DailyStationTimeTableAPIControllerGet2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyStationTimeTableAPIControllerGet2Status299 creates a DailyStationTimeTableAPIControllerGet2Status299 with default headers values
func NewDailyStationTimeTableAPIControllerGet2Status299() *DailyStationTimeTableAPIControllerGet2Status299 {
	return &DailyStationTimeTableAPIControllerGet2Status299{}
}

/* DailyStationTimeTableAPIControllerGet2Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type DailyStationTimeTableAPIControllerGet2Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *DailyStationTimeTableAPIControllerGet2Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/TrainDate/{TrainDate}][%d] dailyStationTimeTableApiControllerGet2Status299  %+v", 299, o.Payload)
}
func (o *DailyStationTimeTableAPIControllerGet2Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *DailyStationTimeTableAPIControllerGet2Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDailyStationTimeTableAPIControllerGet2NotModified creates a DailyStationTimeTableAPIControllerGet2NotModified with default headers values
func NewDailyStationTimeTableAPIControllerGet2NotModified() *DailyStationTimeTableAPIControllerGet2NotModified {
	return &DailyStationTimeTableAPIControllerGet2NotModified{}
}

/* DailyStationTimeTableAPIControllerGet2NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type DailyStationTimeTableAPIControllerGet2NotModified struct {
}

func (o *DailyStationTimeTableAPIControllerGet2NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/DailyStationTimetable/TrainDate/{TrainDate}][%d] dailyStationTimeTableApiControllerGet2NotModified ", 304)
}

func (o *DailyStationTimeTableAPIControllerGet2NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
