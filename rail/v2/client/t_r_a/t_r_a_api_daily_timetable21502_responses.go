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

// TRAAPIDailyTimetable21502Reader is a Reader for the TRAAPIDailyTimetable21502 structure.
type TRAAPIDailyTimetable21502Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIDailyTimetable21502Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIDailyTimetable21502OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTRAAPIDailyTimetable21502Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTRAAPIDailyTimetable21502NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIDailyTimetable21502OK creates a TRAAPIDailyTimetable21502OK with default headers values
func NewTRAAPIDailyTimetable21502OK() *TRAAPIDailyTimetable21502OK {
	return &TRAAPIDailyTimetable21502OK{}
}

/* TRAAPIDailyTimetable21502OK describes a response with status code 200, with default header values.

Success
*/
type TRAAPIDailyTimetable21502OK struct {
	Payload *models.PTXServiceDTORailSpecificationV2TRATrainDateList
}

func (o *TRAAPIDailyTimetable21502OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/TrainDates][%d] tRAApiDailyTimetable21502OK  %+v", 200, o.Payload)
}
func (o *TRAAPIDailyTimetable21502OK) GetPayload() *models.PTXServiceDTORailSpecificationV2TRATrainDateList {
	return o.Payload
}

func (o *TRAAPIDailyTimetable21502OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTORailSpecificationV2TRATrainDateList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIDailyTimetable21502Status299 creates a TRAAPIDailyTimetable21502Status299 with default headers values
func NewTRAAPIDailyTimetable21502Status299() *TRAAPIDailyTimetable21502Status299 {
	return &TRAAPIDailyTimetable21502Status299{}
}

/* TRAAPIDailyTimetable21502Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TRAAPIDailyTimetable21502Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TRAAPIDailyTimetable21502Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/TrainDates][%d] tRAApiDailyTimetable21502Status299  %+v", 299, o.Payload)
}
func (o *TRAAPIDailyTimetable21502Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TRAAPIDailyTimetable21502Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIDailyTimetable21502NotModified creates a TRAAPIDailyTimetable21502NotModified with default headers values
func NewTRAAPIDailyTimetable21502NotModified() *TRAAPIDailyTimetable21502NotModified {
	return &TRAAPIDailyTimetable21502NotModified{}
}

/* TRAAPIDailyTimetable21502NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TRAAPIDailyTimetable21502NotModified struct {
}

func (o *TRAAPIDailyTimetable21502NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/TrainDates][%d] tRAApiDailyTimetable21502NotModified ", 304)
}

func (o *TRAAPIDailyTimetable21502NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
