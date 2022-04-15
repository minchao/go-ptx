// Code generated by go-swagger; DO NOT EDIT.

package a_f_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// LiteTrainStationOfLine3288Reader is a Reader for the LiteTrainStationOfLine3288 structure.
type LiteTrainStationOfLine3288Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LiteTrainStationOfLine3288Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLiteTrainStationOfLine3288OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewLiteTrainStationOfLine3288Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLiteTrainStationOfLine3288NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLiteTrainStationOfLine3288OK creates a LiteTrainStationOfLine3288OK with default headers values
func NewLiteTrainStationOfLine3288OK() *LiteTrainStationOfLine3288OK {
	return &LiteTrainStationOfLine3288OK{}
}

/* LiteTrainStationOfLine3288OK describes a response with status code 200, with default header values.

Success
*/
type LiteTrainStationOfLine3288OK struct {
	Payload *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStationOfLineStationOfLine
}

func (o *LiteTrainStationOfLine3288OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/StationOfLine][%d] liteTrainStationOfLine3288OK  %+v", 200, o.Payload)
}
func (o *LiteTrainStationOfLine3288OK) GetPayload() *models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStationOfLineStationOfLine {
	return o.Payload
}

func (o *LiteTrainStationOfLine3288OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainStationOfLineStationOfLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainStationOfLine3288Status299 creates a LiteTrainStationOfLine3288Status299 with default headers values
func NewLiteTrainStationOfLine3288Status299() *LiteTrainStationOfLine3288Status299 {
	return &LiteTrainStationOfLine3288Status299{}
}

/* LiteTrainStationOfLine3288Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type LiteTrainStationOfLine3288Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *LiteTrainStationOfLine3288Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/StationOfLine][%d] liteTrainStationOfLine3288Status299  %+v", 299, o.Payload)
}
func (o *LiteTrainStationOfLine3288Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *LiteTrainStationOfLine3288Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLiteTrainStationOfLine3288NotModified creates a LiteTrainStationOfLine3288NotModified with default headers values
func NewLiteTrainStationOfLine3288NotModified() *LiteTrainStationOfLine3288NotModified {
	return &LiteTrainStationOfLine3288NotModified{}
}

/* LiteTrainStationOfLine3288NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type LiteTrainStationOfLine3288NotModified struct {
}

func (o *LiteTrainStationOfLine3288NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/AFR/StationOfLine][%d] liteTrainStationOfLine3288NotModified ", 304)
}

func (o *LiteTrainStationOfLine3288NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}