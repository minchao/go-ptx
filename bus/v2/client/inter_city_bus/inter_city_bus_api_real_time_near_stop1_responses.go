// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIRealTimeNearStop1Reader is a Reader for the InterCityBusAPIRealTimeNearStop1 structure.
type InterCityBusAPIRealTimeNearStop1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRealTimeNearStop1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRealTimeNearStop1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRealTimeNearStop1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIRealTimeNearStop1OK creates a InterCityBusAPIRealTimeNearStop1OK with default headers values
func NewInterCityBusAPIRealTimeNearStop1OK() *InterCityBusAPIRealTimeNearStop1OK {
	return &InterCityBusAPIRealTimeNearStop1OK{}
}

/*InterCityBusAPIRealTimeNearStop1OK handles this case with default header values.

OK
*/
type InterCityBusAPIRealTimeNearStop1OK struct {
	Payload []*models.ServiceDTOVersion2BusBusA2Data
}

func (o *InterCityBusAPIRealTimeNearStop1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/InterCity/{RouteName}][%d] interCityBusApiRealTimeNearStop1OK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIRealTimeNearStop1OK) GetPayload() []*models.ServiceDTOVersion2BusBusA2Data {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeNearStop1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeNearStop1Status299 creates a InterCityBusAPIRealTimeNearStop1Status299 with default headers values
func NewInterCityBusAPIRealTimeNearStop1Status299() *InterCityBusAPIRealTimeNearStop1Status299 {
	return &InterCityBusAPIRealTimeNearStop1Status299{}
}

/*InterCityBusAPIRealTimeNearStop1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRealTimeNearStop1Status299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *InterCityBusAPIRealTimeNearStop1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/InterCity/{RouteName}][%d] interCityBusApiRealTimeNearStop1Status299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIRealTimeNearStop1Status299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeNearStop1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
