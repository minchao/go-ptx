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

// TRAAPIGeneralTimetableReader is a Reader for the TRAAPIGeneralTimetable structure.
type TRAAPIGeneralTimetableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIGeneralTimetableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIGeneralTimetableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTRAAPIGeneralTimetableStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTRAAPIGeneralTimetableNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIGeneralTimetableOK creates a TRAAPIGeneralTimetableOK with default headers values
func NewTRAAPIGeneralTimetableOK() *TRAAPIGeneralTimetableOK {
	return &TRAAPIGeneralTimetableOK{}
}

/* TRAAPIGeneralTimetableOK describes a response with status code 200, with default header values.

Success
*/
type TRAAPIGeneralTimetableOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailGeneralTimetable
}

func (o *TRAAPIGeneralTimetableOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable][%d] tRAApiGeneralTimetableOK  %+v", 200, o.Payload)
}
func (o *TRAAPIGeneralTimetableOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailGeneralTimetable {
	return o.Payload
}

func (o *TRAAPIGeneralTimetableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIGeneralTimetableStatus299 creates a TRAAPIGeneralTimetableStatus299 with default headers values
func NewTRAAPIGeneralTimetableStatus299() *TRAAPIGeneralTimetableStatus299 {
	return &TRAAPIGeneralTimetableStatus299{}
}

/* TRAAPIGeneralTimetableStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TRAAPIGeneralTimetableStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TRAAPIGeneralTimetableStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable][%d] tRAApiGeneralTimetableStatus299  %+v", 299, o.Payload)
}
func (o *TRAAPIGeneralTimetableStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TRAAPIGeneralTimetableStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIGeneralTimetableNotModified creates a TRAAPIGeneralTimetableNotModified with default headers values
func NewTRAAPIGeneralTimetableNotModified() *TRAAPIGeneralTimetableNotModified {
	return &TRAAPIGeneralTimetableNotModified{}
}

/* TRAAPIGeneralTimetableNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TRAAPIGeneralTimetableNotModified struct {
}

func (o *TRAAPIGeneralTimetableNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/GeneralTimetable][%d] tRAApiGeneralTimetableNotModified ", 304)
}

func (o *TRAAPIGeneralTimetableNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
