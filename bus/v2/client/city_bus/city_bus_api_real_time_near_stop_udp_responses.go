// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIRealTimeNearStopUDPReader is a Reader for the CityBusAPIRealTimeNearStopUDP structure.
type CityBusAPIRealTimeNearStopUDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeNearStopUDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeNearStopUDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeNearStopUDPStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRealTimeNearStopUDPOK creates a CityBusAPIRealTimeNearStopUDPOK with default headers values
func NewCityBusAPIRealTimeNearStopUDPOK() *CityBusAPIRealTimeNearStopUDPOK {
	return &CityBusAPIRealTimeNearStopUDPOK{}
}

/*CityBusAPIRealTimeNearStopUDPOK handles this case with default header values.

OK
*/
type CityBusAPIRealTimeNearStopUDPOK struct {
	Payload []*models.ServiceDTOVersion2BusBusA2Data
}

func (o *CityBusAPIRealTimeNearStopUDPOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/City/{City}][%d] cityBusApiRealTimeNearStopUdpOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRealTimeNearStopUDPOK) GetPayload() []*models.ServiceDTOVersion2BusBusA2Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeNearStopUDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeNearStopUDPStatus299 creates a CityBusAPIRealTimeNearStopUDPStatus299 with default headers values
func NewCityBusAPIRealTimeNearStopUDPStatus299() *CityBusAPIRealTimeNearStopUDPStatus299 {
	return &CityBusAPIRealTimeNearStopUDPStatus299{}
}

/*CityBusAPIRealTimeNearStopUDPStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeNearStopUDPStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeNearStopUDPStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/Streaming/City/{City}][%d] cityBusApiRealTimeNearStopUdpStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRealTimeNearStopUDPStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeNearStopUDPStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
