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

// TaiwanTripBusAPIStopOfRoute22641Reader is a Reader for the TaiwanTripBusAPIStopOfRoute22641 structure.
type TaiwanTripBusAPIStopOfRoute22641Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaiwanTripBusAPIStopOfRoute22641Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaiwanTripBusAPIStopOfRoute22641OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewTaiwanTripBusAPIStopOfRoute22641Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaiwanTripBusAPIStopOfRoute22641OK creates a TaiwanTripBusAPIStopOfRoute22641OK with default headers values
func NewTaiwanTripBusAPIStopOfRoute22641OK() *TaiwanTripBusAPIStopOfRoute22641OK {
	return &TaiwanTripBusAPIStopOfRoute22641OK{}
}

/* TaiwanTripBusAPIStopOfRoute22641OK describes a response with status code 200, with default header values.

Success
*/
type TaiwanTripBusAPIStopOfRoute22641OK struct {
	Payload []*models.PTXServiceDTOTourismSpecificationV2BusStopOfRoute
}

func (o *TaiwanTripBusAPIStopOfRoute22641OK) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/StopOfRoute/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiStopOfRoute22641OK  %+v", 200, o.Payload)
}
func (o *TaiwanTripBusAPIStopOfRoute22641OK) GetPayload() []*models.PTXServiceDTOTourismSpecificationV2BusStopOfRoute {
	return o.Payload
}

func (o *TaiwanTripBusAPIStopOfRoute22641OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaiwanTripBusAPIStopOfRoute22641Status299 creates a TaiwanTripBusAPIStopOfRoute22641Status299 with default headers values
func NewTaiwanTripBusAPIStopOfRoute22641Status299() *TaiwanTripBusAPIStopOfRoute22641Status299 {
	return &TaiwanTripBusAPIStopOfRoute22641Status299{}
}

/* TaiwanTripBusAPIStopOfRoute22641Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type TaiwanTripBusAPIStopOfRoute22641Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *TaiwanTripBusAPIStopOfRoute22641Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Tourism/Bus/StopOfRoute/TaiwanTrip/{TaiwanTripName}][%d] taiwanTripBusApiStopOfRoute22641Status299  %+v", 299, o.Payload)
}
func (o *TaiwanTripBusAPIStopOfRoute22641Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *TaiwanTripBusAPIStopOfRoute22641Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
