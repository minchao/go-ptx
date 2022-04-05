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

// TaiwanTripBusAPIStopOfRoute2264Reader is a Reader for the TaiwanTripBusAPIStopOfRoute2264 structure.
type TaiwanTripBusAPIStopOfRoute2264Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIStopOfRoute2264Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIStopOfRoute2264OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTaiwanTripBusAPIStopOfRoute2264Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIStopOfRoute2264OK creates a TaiwanTripBusAPIStopOfRoute2264OK with default headers values
func NewTaiwanTripBusAPIStopOfRoute2264OK() *TaiwanTripBusAPIStopOfRoute2264OK {
	return &TaiwanTripBusAPIStopOfRoute2264OK{}
}

/* TaiwanTripBusAPIStopOfRoute2264OK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIStopOfRoute2264OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusStopOfRoute
}

func (o *TaiwanTripBusAPIStopOfRoute2264OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/StopOfRoute/TaiwanTrip][%d] taiwanTripBusApiStopOfRoute2264OK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIStopOfRoute2264OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusStopOfRoute {
	return o.Payload
}

func (o *TaiwanTripBusAPIStopOfRoute2264OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaiwanTripBusAPIStopOfRoute2264Status299 creates a TaiwanTripBusAPIStopOfRoute2264Status299 with default headers values
func NewTaiwanTripBusAPIStopOfRoute2264Status299() *TaiwanTripBusAPIStopOfRoute2264Status299 {
	return &TaiwanTripBusAPIStopOfRoute2264Status299{}
}

/* TaiwanTripBusAPIStopOfRoute2264Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TaiwanTripBusAPIStopOfRoute2264Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TaiwanTripBusAPIStopOfRoute2264Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/StopOfRoute/TaiwanTrip][%d] taiwanTripBusApiStopOfRoute2264Status299  %+v", 299, o.Payload)
}
func (o *TaiwanTripBusAPIStopOfRoute2264Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TaiwanTripBusAPIStopOfRoute2264Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
