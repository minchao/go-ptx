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

// TaiwanTripBusAPIEstimatedTimeOfArrivalReader is a Reader for the TaiwanTripBusAPIEstimatedTimeOfArrival structure.
type TaiwanTripBusAPIEstimatedTimeOfArrivalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIEstimatedTimeOfArrivalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTaiwanTripBusAPIEstimatedTimeOfArrivalStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrivalOK creates a TaiwanTripBusAPIEstimatedTimeOfArrivalOK with default headers values
func NewTaiwanTripBusAPIEstimatedTimeOfArrivalOK() *TaiwanTripBusAPIEstimatedTimeOfArrivalOK {
	return &TaiwanTripBusAPIEstimatedTimeOfArrivalOK{}
}

/* TaiwanTripBusAPIEstimatedTimeOfArrivalOK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIEstimatedTimeOfArrivalOK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusN1EstimateTime
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalOK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/EstimatedTimeOfArrival/TaiwanTrip][%d] taiwanTripBusApiEstimatedTimeOfArrivalOK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalOK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusN1EstimateTime {
	return o.Payload
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaiwanTripBusAPIEstimatedTimeOfArrivalStatus299 creates a TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299 with default headers values
func NewTaiwanTripBusAPIEstimatedTimeOfArrivalStatus299() *TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299 {
	return &TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299{}
}

/* TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/EstimatedTimeOfArrival/TaiwanTrip][%d] taiwanTripBusApiEstimatedTimeOfArrivalStatus299  %+v", 299, o.Payload)
}
func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TaiwanTripBusAPIEstimatedTimeOfArrivalStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
