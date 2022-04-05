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

// THSRAPIDailyTimetable21241Reader is a Reader for the THSRAPIDailyTimetable21241 structure.
type THSRAPIDailyTimetable21241Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTimetable21241Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTimetable21241OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPIDailyTimetable21241Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIDailyTimetable21241NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIDailyTimetable21241OK creates a THSRAPIDailyTimetable21241OK with default headers values
func NewTHSRAPIDailyTimetable21241OK() *THSRAPIDailyTimetable21241OK {
	return &THSRAPIDailyTimetable21241OK{}
}

/* THSRAPIDailyTimetable21241OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIDailyTimetable21241OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTimetable
}

func (o *THSRAPIDailyTimetable21241OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Today/TrainNo/{TrainNo}][%d] tHSRApiDailyTimetable21241OK  %+v", 200, o.Payload)
}
func (o *THSRAPIDailyTimetable21241OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTimetable {
	return o.Payload
}

func (o *THSRAPIDailyTimetable21241OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTimetable21241Status299 creates a THSRAPIDailyTimetable21241Status299 with default headers values
func NewTHSRAPIDailyTimetable21241Status299() *THSRAPIDailyTimetable21241Status299 {
	return &THSRAPIDailyTimetable21241Status299{}
}

/* THSRAPIDailyTimetable21241Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPIDailyTimetable21241Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPIDailyTimetable21241Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Today/TrainNo/{TrainNo}][%d] tHSRApiDailyTimetable21241Status299  %+v", 299, o.Payload)
}
func (o *THSRAPIDailyTimetable21241Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPIDailyTimetable21241Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTimetable21241NotModified creates a THSRAPIDailyTimetable21241NotModified with default headers values
func NewTHSRAPIDailyTimetable21241NotModified() *THSRAPIDailyTimetable21241NotModified {
	return &THSRAPIDailyTimetable21241NotModified{}
}

/* THSRAPIDailyTimetable21241NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIDailyTimetable21241NotModified struct {
}

func (o *THSRAPIDailyTimetable21241NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Today/TrainNo/{TrainNo}][%d] tHSRApiDailyTimetable21241NotModified ", 304)
}

func (o *THSRAPIDailyTimetable21241NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
