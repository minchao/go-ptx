// Code generated by go-swagger; DO NOT EDIT.

package advanced

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// BusAPIRealTimeByFrequencyNearByReader is a Reader for the BusAPIRealTimeByFrequencyNearBy structure.
type BusAPIRealTimeByFrequencyNearByReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BusAPIRealTimeByFrequencyNearByReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBusAPIRealTimeByFrequencyNearByOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewBusAPIRealTimeByFrequencyNearByStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBusAPIRealTimeByFrequencyNearByNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBusAPIRealTimeByFrequencyNearByOK creates a BusAPIRealTimeByFrequencyNearByOK with default headers values
func NewBusAPIRealTimeByFrequencyNearByOK() *BusAPIRealTimeByFrequencyNearByOK {
	return &BusAPIRealTimeByFrequencyNearByOK{}
}

/* BusAPIRealTimeByFrequencyNearByOK describes a response with status code 200, with default header values.

Success
*/
type BusAPIRealTimeByFrequencyNearByOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA1Data
}

func (o *BusAPIRealTimeByFrequencyNearByOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/NearBy][%d] busApiRealTimeByFrequencyNearByOK  %+v", 200, o.Payload)
}
func (o *BusAPIRealTimeByFrequencyNearByOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA1Data {
	return o.Payload
}

func (o *BusAPIRealTimeByFrequencyNearByOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRealTimeByFrequencyNearByStatus299 creates a BusAPIRealTimeByFrequencyNearByStatus299 with default headers values
func NewBusAPIRealTimeByFrequencyNearByStatus299() *BusAPIRealTimeByFrequencyNearByStatus299 {
	return &BusAPIRealTimeByFrequencyNearByStatus299{}
}

/* BusAPIRealTimeByFrequencyNearByStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type BusAPIRealTimeByFrequencyNearByStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *BusAPIRealTimeByFrequencyNearByStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/NearBy][%d] busApiRealTimeByFrequencyNearByStatus299  %+v", 299, o.Payload)
}
func (o *BusAPIRealTimeByFrequencyNearByStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *BusAPIRealTimeByFrequencyNearByStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRealTimeByFrequencyNearByNotModified creates a BusAPIRealTimeByFrequencyNearByNotModified with default headers values
func NewBusAPIRealTimeByFrequencyNearByNotModified() *BusAPIRealTimeByFrequencyNearByNotModified {
	return &BusAPIRealTimeByFrequencyNearByNotModified{}
}

/* BusAPIRealTimeByFrequencyNearByNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BusAPIRealTimeByFrequencyNearByNotModified struct {
}

func (o *BusAPIRealTimeByFrequencyNearByNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/NearBy][%d] busApiRealTimeByFrequencyNearByNotModified ", 304)
}

func (o *BusAPIRealTimeByFrequencyNearByNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
