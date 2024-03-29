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

// TRAAPIStationTimetable2151Reader is a Reader for the TRAAPIStationTimetable2151 structure.
type TRAAPIStationTimetable2151Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIStationTimetable2151Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIStationTimetable2151OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTRAAPIStationTimetable2151Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTRAAPIStationTimetable2151NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIStationTimetable2151OK creates a TRAAPIStationTimetable2151OK with default headers values
func NewTRAAPIStationTimetable2151OK() *TRAAPIStationTimetable2151OK {
	return &TRAAPIStationTimetable2151OK{}
}

/* TRAAPIStationTimetable2151OK describes a response with status code 200, with default header values.

Success
*/
type TRAAPIStationTimetable2151OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailStationTimetable
}

func (o *TRAAPIStationTimetable2151OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tRAApiStationTimetable2151OK  %+v", 200, o.Payload)
}
func (o *TRAAPIStationTimetable2151OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailStationTimetable {
	return o.Payload
}

func (o *TRAAPIStationTimetable2151OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIStationTimetable2151Status299 creates a TRAAPIStationTimetable2151Status299 with default headers values
func NewTRAAPIStationTimetable2151Status299() *TRAAPIStationTimetable2151Status299 {
	return &TRAAPIStationTimetable2151Status299{}
}

/* TRAAPIStationTimetable2151Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TRAAPIStationTimetable2151Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TRAAPIStationTimetable2151Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tRAApiStationTimetable2151Status299  %+v", 299, o.Payload)
}
func (o *TRAAPIStationTimetable2151Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TRAAPIStationTimetable2151Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIStationTimetable2151NotModified creates a TRAAPIStationTimetable2151NotModified with default headers values
func NewTRAAPIStationTimetable2151NotModified() *TRAAPIStationTimetable2151NotModified {
	return &TRAAPIStationTimetable2151NotModified{}
}

/* TRAAPIStationTimetable2151NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TRAAPIStationTimetable2151NotModified struct {
}

func (o *TRAAPIStationTimetable2151NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tRAApiStationTimetable2151NotModified ", 304)
}

func (o *TRAAPIStationTimetable2151NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
