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

// THSRAPIDailyTrainInfo2123Reader is a Reader for the THSRAPIDailyTrainInfo2123 structure.
type THSRAPIDailyTrainInfo2123Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTrainInfo2123Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTrainInfo2123OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTHSRAPIDailyTrainInfo2123Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTHSRAPIDailyTrainInfo2123NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIDailyTrainInfo2123OK creates a THSRAPIDailyTrainInfo2123OK with default headers values
func NewTHSRAPIDailyTrainInfo2123OK() *THSRAPIDailyTrainInfo2123OK {
	return &THSRAPIDailyTrainInfo2123OK{}
}

/* THSRAPIDailyTrainInfo2123OK describes a response with status code 200, with default header values.

Success
*/
type THSRAPIDailyTrainInfo2123OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTrainInfo
}

func (o *THSRAPIDailyTrainInfo2123OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/Today][%d] tHSRApiDailyTrainInfo2123OK  %+v", 200, o.Payload)
}
func (o *THSRAPIDailyTrainInfo2123OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailDailyTrainInfo {
	return o.Payload
}

func (o *THSRAPIDailyTrainInfo2123OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTrainInfo2123Status299 creates a THSRAPIDailyTrainInfo2123Status299 with default headers values
func NewTHSRAPIDailyTrainInfo2123Status299() *THSRAPIDailyTrainInfo2123Status299 {
	return &THSRAPIDailyTrainInfo2123Status299{}
}

/* THSRAPIDailyTrainInfo2123Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type THSRAPIDailyTrainInfo2123Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *THSRAPIDailyTrainInfo2123Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/Today][%d] tHSRApiDailyTrainInfo2123Status299  %+v", 299, o.Payload)
}
func (o *THSRAPIDailyTrainInfo2123Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *THSRAPIDailyTrainInfo2123Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTHSRAPIDailyTrainInfo2123NotModified creates a THSRAPIDailyTrainInfo2123NotModified with default headers values
func NewTHSRAPIDailyTrainInfo2123NotModified() *THSRAPIDailyTrainInfo2123NotModified {
	return &THSRAPIDailyTrainInfo2123NotModified{}
}

/* THSRAPIDailyTrainInfo2123NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type THSRAPIDailyTrainInfo2123NotModified struct {
}

func (o *THSRAPIDailyTrainInfo2123NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/Today][%d] tHSRApiDailyTrainInfo2123NotModified ", 304)
}

func (o *THSRAPIDailyTrainInfo2123NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}