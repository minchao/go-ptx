// Code generated by go-swagger; DO NOT EDIT.

package tourism

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/tourism/v2/models"
)

// TaiwanTripBusAPIRealTimeNearStop22611Reader is a Reader for the TaiwanTripBusAPIRealTimeNearStop22611 structure.
type TaiwanTripBusAPIRealTimeNearStop22611Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIRealTimeNearStop22611Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIRealTimeNearStop22611OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTaiwanTripBusAPIRealTimeNearStop22611Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIRealTimeNearStop22611OK creates a TaiwanTripBusAPIRealTimeNearStop22611OK with default headers values
func NewTaiwanTripBusAPIRealTimeNearStop22611OK() *TaiwanTripBusAPIRealTimeNearStop22611OK {
	return &TaiwanTripBusAPIRealTimeNearStop22611OK{}
}

/* TaiwanTripBusAPIRealTimeNearStop22611OK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIRealTimeNearStop22611OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusA2Data
}

func (o *TaiwanTripBusAPIRealTimeNearStop22611OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/RealTimeNearStop/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiRealTimeNearStop22611OK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIRealTimeNearStop22611OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusA2Data {
	return o.Payload
}

func (o *TaiwanTripBusAPIRealTimeNearStop22611OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaiwanTripBusAPIRealTimeNearStop22611Status299 creates a TaiwanTripBusAPIRealTimeNearStop22611Status299 with default headers values
func NewTaiwanTripBusAPIRealTimeNearStop22611Status299() *TaiwanTripBusAPIRealTimeNearStop22611Status299 {
	return &TaiwanTripBusAPIRealTimeNearStop22611Status299{}
}

/* TaiwanTripBusAPIRealTimeNearStop22611Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TaiwanTripBusAPIRealTimeNearStop22611Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TaiwanTripBusAPIRealTimeNearStop22611Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/RealTimeNearStop/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiRealTimeNearStop22611Status299  %+v", 299, o.Payload)
}
func (o *TaiwanTripBusAPIRealTimeNearStop22611Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TaiwanTripBusAPIRealTimeNearStop22611Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}