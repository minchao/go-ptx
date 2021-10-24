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

// BusAPIRealTimeByFrequencyByOperatorReader is a Reader for the BusAPIRealTimeByFrequencyByOperator structure.
type BusAPIRealTimeByFrequencyByOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BusAPIRealTimeByFrequencyByOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBusAPIRealTimeByFrequencyByOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewBusAPIRealTimeByFrequencyByOperatorStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBusAPIRealTimeByFrequencyByOperatorNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBusAPIRealTimeByFrequencyByOperatorOK creates a BusAPIRealTimeByFrequencyByOperatorOK with default headers values
func NewBusAPIRealTimeByFrequencyByOperatorOK() *BusAPIRealTimeByFrequencyByOperatorOK {
	return &BusAPIRealTimeByFrequencyByOperatorOK{}
}

/* BusAPIRealTimeByFrequencyByOperatorOK describes a response with status code 200, with default header values.

Success
*/
type BusAPIRealTimeByFrequencyByOperatorOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA1Data
}

func (o *BusAPIRealTimeByFrequencyByOperatorOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/OperatorNo/{OperatorNo}][%d] busApiRealTimeByFrequencyByOperatorOK  %+v", 200, o.Payload)
}
func (o *BusAPIRealTimeByFrequencyByOperatorOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA1Data {
	return o.Payload
}

func (o *BusAPIRealTimeByFrequencyByOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRealTimeByFrequencyByOperatorStatus299 creates a BusAPIRealTimeByFrequencyByOperatorStatus299 with default headers values
func NewBusAPIRealTimeByFrequencyByOperatorStatus299() *BusAPIRealTimeByFrequencyByOperatorStatus299 {
	return &BusAPIRealTimeByFrequencyByOperatorStatus299{}
}

/* BusAPIRealTimeByFrequencyByOperatorStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type BusAPIRealTimeByFrequencyByOperatorStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *BusAPIRealTimeByFrequencyByOperatorStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/OperatorNo/{OperatorNo}][%d] busApiRealTimeByFrequencyByOperatorStatus299  %+v", 299, o.Payload)
}
func (o *BusAPIRealTimeByFrequencyByOperatorStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *BusAPIRealTimeByFrequencyByOperatorStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRealTimeByFrequencyByOperatorNotModified creates a BusAPIRealTimeByFrequencyByOperatorNotModified with default headers values
func NewBusAPIRealTimeByFrequencyByOperatorNotModified() *BusAPIRealTimeByFrequencyByOperatorNotModified {
	return &BusAPIRealTimeByFrequencyByOperatorNotModified{}
}

/* BusAPIRealTimeByFrequencyByOperatorNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BusAPIRealTimeByFrequencyByOperatorNotModified struct {
}

func (o *BusAPIRealTimeByFrequencyByOperatorNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/OperatorNo/{OperatorNo}][%d] busApiRealTimeByFrequencyByOperatorNotModified ", 304)
}

func (o *BusAPIRealTimeByFrequencyByOperatorNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
