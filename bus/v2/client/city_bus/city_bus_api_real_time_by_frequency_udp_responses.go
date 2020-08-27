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

// CityBusAPIRealTimeByFrequencyUDPReader is a Reader for the CityBusAPIRealTimeByFrequencyUDP structure.
type CityBusAPIRealTimeByFrequencyUDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeByFrequencyUDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeByFrequencyUDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeByFrequencyUDPStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRealTimeByFrequencyUDPOK creates a CityBusAPIRealTimeByFrequencyUDPOK with default headers values
func NewCityBusAPIRealTimeByFrequencyUDPOK() *CityBusAPIRealTimeByFrequencyUDPOK {
	return &CityBusAPIRealTimeByFrequencyUDPOK{}
}

/*CityBusAPIRealTimeByFrequencyUDPOK handles this case with default header values.

OK
*/
type CityBusAPIRealTimeByFrequencyUDPOK struct {
	Payload []*models.ServiceDTOVersion2BusBusA1Data
}

func (o *CityBusAPIRealTimeByFrequencyUDPOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/City/{City}][%d] cityBusApiRealTimeByFrequencyUdpOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRealTimeByFrequencyUDPOK) GetPayload() []*models.ServiceDTOVersion2BusBusA1Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequencyUDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeByFrequencyUDPStatus299 creates a CityBusAPIRealTimeByFrequencyUDPStatus299 with default headers values
func NewCityBusAPIRealTimeByFrequencyUDPStatus299() *CityBusAPIRealTimeByFrequencyUDPStatus299 {
	return &CityBusAPIRealTimeByFrequencyUDPStatus299{}
}

/*CityBusAPIRealTimeByFrequencyUDPStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeByFrequencyUDPStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeByFrequencyUDPStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/City/{City}][%d] cityBusApiRealTimeByFrequencyUdpStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRealTimeByFrequencyUDPStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequencyUDPStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
