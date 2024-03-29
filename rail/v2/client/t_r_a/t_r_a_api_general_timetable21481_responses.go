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

// TRAAPIGeneralTimetable21481Reader is a Reader for the TRAAPIGeneralTimetable21481 structure.
type TRAAPIGeneralTimetable21481Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIGeneralTimetable21481Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIGeneralTimetable21481OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTRAAPIGeneralTimetable21481Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTRAAPIGeneralTimetable21481NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIGeneralTimetable21481OK creates a TRAAPIGeneralTimetable21481OK with default headers values
func NewTRAAPIGeneralTimetable21481OK() *TRAAPIGeneralTimetable21481OK {
	return &TRAAPIGeneralTimetable21481OK{}
}

/* TRAAPIGeneralTimetable21481OK describes a response with status code 200, with default header values.

Success
*/
type TRAAPIGeneralTimetable21481OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailGeneralTimetable
}

func (o *TRAAPIGeneralTimetable21481OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable/TrainNo/{TrainNo}][%d] tRAApiGeneralTimetable21481OK  %+v", 200, o.Payload)
}
func (o *TRAAPIGeneralTimetable21481OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailGeneralTimetable {
	return o.Payload
}

func (o *TRAAPIGeneralTimetable21481OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIGeneralTimetable21481Status299 creates a TRAAPIGeneralTimetable21481Status299 with default headers values
func NewTRAAPIGeneralTimetable21481Status299() *TRAAPIGeneralTimetable21481Status299 {
	return &TRAAPIGeneralTimetable21481Status299{}
}

/* TRAAPIGeneralTimetable21481Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TRAAPIGeneralTimetable21481Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TRAAPIGeneralTimetable21481Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable/TrainNo/{TrainNo}][%d] tRAApiGeneralTimetable21481Status299  %+v", 299, o.Payload)
}
func (o *TRAAPIGeneralTimetable21481Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TRAAPIGeneralTimetable21481Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIGeneralTimetable21481NotModified creates a TRAAPIGeneralTimetable21481NotModified with default headers values
func NewTRAAPIGeneralTimetable21481NotModified() *TRAAPIGeneralTimetable21481NotModified {
	return &TRAAPIGeneralTimetable21481NotModified{}
}

/* TRAAPIGeneralTimetable21481NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TRAAPIGeneralTimetable21481NotModified struct {
}

func (o *TRAAPIGeneralTimetable21481NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable/TrainNo/{TrainNo}][%d] tRAApiGeneralTimetable21481NotModified ", 304)
}

func (o *TRAAPIGeneralTimetable21481NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
