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

// THSRAPIDailyTimetable4Reader is a Reader for the THSRAPIDailyTimetable4 structure.
type THSRAPIDailyTimetable4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTimetable4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTimetable4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPIDailyTimetable4Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIDailyTimetable4NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIDailyTimetable4OK creates a THSRAPIDailyTimetable4OK with default headers values
func NewTHSRAPIDailyTimetable4OK() *THSRAPIDailyTimetable4OK {
	return &THSRAPIDailyTimetable4OK{}
}

/* THSRAPIDailyTimetable4OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIDailyTimetable4OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTimetable
}

func (o *THSRAPIDailyTimetable4OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/TrainNo/{TrainNo}/TrainDate/{TrainDate}][%d] tHSRApiDailyTimetable4OK  %+v", 200, o.Payload)
}
func (o *THSRAPIDailyTimetable4OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTimetable {
	return o.Payload
}

func (o *THSRAPIDailyTimetable4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTimetable4Status299 creates a THSRAPIDailyTimetable4Status299 with default headers values
func NewTHSRAPIDailyTimetable4Status299() *THSRAPIDailyTimetable4Status299 {
	return &THSRAPIDailyTimetable4Status299{}
}

/* THSRAPIDailyTimetable4Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPIDailyTimetable4Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPIDailyTimetable4Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/TrainNo/{TrainNo}/TrainDate/{TrainDate}][%d] tHSRApiDailyTimetable4Status299  %+v", 299, o.Payload)
}
func (o *THSRAPIDailyTimetable4Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPIDailyTimetable4Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTimetable4NotModified creates a THSRAPIDailyTimetable4NotModified with default headers values
func NewTHSRAPIDailyTimetable4NotModified() *THSRAPIDailyTimetable4NotModified {
	return &THSRAPIDailyTimetable4NotModified{}
}

/* THSRAPIDailyTimetable4NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIDailyTimetable4NotModified struct {
}

func (o *THSRAPIDailyTimetable4NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/TrainNo/{TrainNo}/TrainDate/{TrainDate}][%d] tHSRApiDailyTimetable4NotModified ", 304)
}

func (o *THSRAPIDailyTimetable4NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
