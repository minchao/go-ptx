// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_near_by

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// BusAPIRealTimeNearStopNearBy2854Reader is a Reader for the BusAPIRealTimeNearStopNearBy2854 structure.
type BusAPIRealTimeNearStopNearBy2854Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BusAPIRealTimeNearStopNearBy2854Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBusAPIRealTimeNearStopNearBy2854OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewBusAPIRealTimeNearStopNearBy2854Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBusAPIRealTimeNearStopNearBy2854NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBusAPIRealTimeNearStopNearBy2854OK creates a BusAPIRealTimeNearStopNearBy2854OK with default headers values
func NewBusAPIRealTimeNearStopNearBy2854OK() *BusAPIRealTimeNearStopNearBy2854OK {
	return &BusAPIRealTimeNearStopNearBy2854OK{}
}

/* BusAPIRealTimeNearStopNearBy2854OK describes a response with status code 200, with default header values.

Success
*/
type BusAPIRealTimeNearStopNearBy2854OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA2Data
}

func (o *BusAPIRealTimeNearStopNearBy2854OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/NearBy][%d] busApiRealTimeNearStopNearBy2854OK  %+v", 200, o.Payload)
}
func (o *BusAPIRealTimeNearStopNearBy2854OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA2Data {
	return o.Payload
}

func (o *BusAPIRealTimeNearStopNearBy2854OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRealTimeNearStopNearBy2854Status299 creates a BusAPIRealTimeNearStopNearBy2854Status299 with default headers values
func NewBusAPIRealTimeNearStopNearBy2854Status299() *BusAPIRealTimeNearStopNearBy2854Status299 {
	return &BusAPIRealTimeNearStopNearBy2854Status299{}
}

/* BusAPIRealTimeNearStopNearBy2854Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type BusAPIRealTimeNearStopNearBy2854Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *BusAPIRealTimeNearStopNearBy2854Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/NearBy][%d] busApiRealTimeNearStopNearBy2854Status299  %+v", 299, o.Payload)
}
func (o *BusAPIRealTimeNearStopNearBy2854Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *BusAPIRealTimeNearStopNearBy2854Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRealTimeNearStopNearBy2854NotModified creates a BusAPIRealTimeNearStopNearBy2854NotModified with default headers values
func NewBusAPIRealTimeNearStopNearBy2854NotModified() *BusAPIRealTimeNearStopNearBy2854NotModified {
	return &BusAPIRealTimeNearStopNearBy2854NotModified{}
}

/* BusAPIRealTimeNearStopNearBy2854NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BusAPIRealTimeNearStopNearBy2854NotModified struct {
}

func (o *BusAPIRealTimeNearStopNearBy2854NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/NearBy][%d] busApiRealTimeNearStopNearBy2854NotModified ", 304)
}

func (o *BusAPIRealTimeNearStopNearBy2854NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}