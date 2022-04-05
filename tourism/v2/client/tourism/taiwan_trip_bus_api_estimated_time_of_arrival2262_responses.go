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

// TaiwanTripBusAPIEstimatedTimeOfArrival2262Reader is a Reader for the TaiwanTripBusAPIEstimatedTimeOfArrival2262 structure.
type TaiwanTripBusAPIEstimatedTimeOfArrival2262Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIEstimatedTimeOfArrival2262OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTaiwanTripBusAPIEstimatedTimeOfArrival2262Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrival2262OK creates a TaiwanTripBusAPIEstimatedTimeOfArrival2262OK with default headers values
func NewTaiwanTripBusAPIEstimatedTimeOfArrival2262OK() *TaiwanTripBusAPIEstimatedTimeOfArrival2262OK {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival2262OK{}
}

/* TaiwanTripBusAPIEstimatedTimeOfArrival2262OK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIEstimatedTimeOfArrival2262OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusN1EstimateTime
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/EstimatedTimeOfArrival/TaiwanTrip][%d] taiwanTripBusApiEstimatedTimeOfArrival2262OK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusN1EstimateTime {
	return o.Payload
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrival2262Status299 creates a TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299 with default headers values
func NewTaiwanTripBusAPIEstimatedTimeOfArrival2262Status299() *TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299 {
	return &TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299{}
}

/* TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/EstimatedTimeOfArrival/TaiwanTrip][%d] taiwanTripBusApiEstimatedTimeOfArrival2262Status299  %+v", 299, o.Payload)
}
func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrival2262Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
