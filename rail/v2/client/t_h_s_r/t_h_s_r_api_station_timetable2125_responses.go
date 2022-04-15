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

// THSRAPIStationTimetable2125Reader is a Reader for the THSRAPIStationTimetable2125 structure.
type THSRAPIStationTimetable2125Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIStationTimetable2125Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIStationTimetable2125OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPIStationTimetable2125Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIStationTimetable2125NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIStationTimetable2125OK creates a THSRAPIStationTimetable2125OK with default headers values
func NewTHSRAPIStationTimetable2125OK() *THSRAPIStationTimetable2125OK {
	return &THSRAPIStationTimetable2125OK{}
}

/* THSRAPIStationTimetable2125OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIStationTimetable2125OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailStationTimetable
}

func (o *THSRAPIStationTimetable2125OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tHSRApiStationTimetable2125OK  %+v", 200, o.Payload)
}
func (o *THSRAPIStationTimetable2125OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailStationTimetable {
	return o.Payload
}

func (o *THSRAPIStationTimetable2125OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIStationTimetable2125Status299 creates a THSRAPIStationTimetable2125Status299 with default headers values
func NewTHSRAPIStationTimetable2125Status299() *THSRAPIStationTimetable2125Status299 {
	return &THSRAPIStationTimetable2125Status299{}
}

/* THSRAPIStationTimetable2125Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPIStationTimetable2125Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPIStationTimetable2125Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tHSRApiStationTimetable2125Status299  %+v", 299, o.Payload)
}
func (o *THSRAPIStationTimetable2125Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPIStationTimetable2125Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIStationTimetable2125NotModified creates a THSRAPIStationTimetable2125NotModified with default headers values
func NewTHSRAPIStationTimetable2125NotModified() *THSRAPIStationTimetable2125NotModified {
	return &THSRAPIStationTimetable2125NotModified{}
}

/* THSRAPIStationTimetable2125NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIStationTimetable2125NotModified struct {
}

func (o *THSRAPIStationTimetable2125NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTimetable/Station/{StationID}/{TrainDate}][%d] tHSRApiStationTimetable2125NotModified ", 304)
}

func (o *THSRAPIStationTimetable2125NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}