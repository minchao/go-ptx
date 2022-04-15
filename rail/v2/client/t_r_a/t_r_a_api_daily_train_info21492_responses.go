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

// TRAAPIDailyTrainInfo21492Reader is a Reader for the TRAAPIDailyTrainInfo21492 structure.
type TRAAPIDailyTrainInfo21492Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIDailyTrainInfo21492Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIDailyTrainInfo21492OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTRAAPIDailyTrainInfo21492Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewTRAAPIDailyTrainInfo21492NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTRAAPIDailyTrainInfo21492OK creates a TRAAPIDailyTrainInfo21492OK with default headers values
func NewTRAAPIDailyTrainInfo21492OK() *TRAAPIDailyTrainInfo21492OK {
	return &TRAAPIDailyTrainInfo21492OK{}
}

/* TRAAPIDailyTrainInfo21492OK describes a response with status code 200, with default header values.

Success
*/
type TRAAPIDailyTrainInfo21492OK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2TRARailDailyTrainInfo
}

func (o *TRAAPIDailyTrainInfo21492OK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTrainInfo/TrainDate/{TrainDate}][%d] tRAApiDailyTrainInfo21492OK  %+v", 200, o.Payload)
}
func (o *TRAAPIDailyTrainInfo21492OK) GetPayload() []*models.PTXServiceDTORailSpecificationV2TRARailDailyTrainInfo {
	return o.Payload
}

func (o *TRAAPIDailyTrainInfo21492OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIDailyTrainInfo21492Status299 creates a TRAAPIDailyTrainInfo21492Status299 with default headers values
func NewTRAAPIDailyTrainInfo21492Status299() *TRAAPIDailyTrainInfo21492Status299 {
	return &TRAAPIDailyTrainInfo21492Status299{}
}

/* TRAAPIDailyTrainInfo21492Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TRAAPIDailyTrainInfo21492Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TRAAPIDailyTrainInfo21492Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTrainInfo/TrainDate/{TrainDate}][%d] tRAApiDailyTrainInfo21492Status299  %+v", 299, o.Payload)
}
func (o *TRAAPIDailyTrainInfo21492Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TRAAPIDailyTrainInfo21492Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTRAAPIDailyTrainInfo21492NotModified creates a TRAAPIDailyTrainInfo21492NotModified with default headers values
func NewTRAAPIDailyTrainInfo21492NotModified() *TRAAPIDailyTrainInfo21492NotModified {
	return &TRAAPIDailyTrainInfo21492NotModified{}
}

/* TRAAPIDailyTrainInfo21492NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type TRAAPIDailyTrainInfo21492NotModified struct {
}

func (o *TRAAPIDailyTrainInfo21492NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTrainInfo/TrainDate/{TrainDate}][%d] tRAApiDailyTrainInfo21492NotModified ", 304)
}

func (o *TRAAPIDailyTrainInfo21492NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}