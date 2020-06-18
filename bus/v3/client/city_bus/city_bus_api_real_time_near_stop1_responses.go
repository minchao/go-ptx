// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v3/models"
)

// CityBusAPIRealTimeNearStop1Reader is a Reader for the CityBusAPIRealTimeNearStop1 structure.
type CityBusAPIRealTimeNearStop1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeNearStop1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeNearStop1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeNearStop1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRealTimeNearStop1OK creates a CityBusAPIRealTimeNearStop1OK with default headers values
func NewCityBusAPIRealTimeNearStop1OK() *CityBusAPIRealTimeNearStop1OK {
	return &CityBusAPIRealTimeNearStop1OK{}
}

/*CityBusAPIRealTimeNearStop1OK handles this case with default header values.

OK
*/
type CityBusAPIRealTimeNearStop1OK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusA2Data
}

func (o *CityBusAPIRealTimeNearStop1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/RealTimeNearStop/City/{City}/{RouteName}][%d] cityBusApiRealTimeNearStop1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRealTimeNearStop1OK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusA2Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeNearStop1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusA2Data)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeNearStop1Status299 creates a CityBusAPIRealTimeNearStop1Status299 with default headers values
func NewCityBusAPIRealTimeNearStop1Status299() *CityBusAPIRealTimeNearStop1Status299 {
	return &CityBusAPIRealTimeNearStop1Status299{}
}

/*CityBusAPIRealTimeNearStop1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeNearStop1Status299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeNearStop1Status299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/RealTimeNearStop/City/{City}/{RouteName}][%d] cityBusApiRealTimeNearStop1Status299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRealTimeNearStop1Status299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeNearStop1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
