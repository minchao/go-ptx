// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIRealTimeByFrequencyUDP1Reader is a Reader for the InterCityBusAPIRealTimeByFrequencyUDP1 structure.
type InterCityBusAPIRealTimeByFrequencyUDP1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRealTimeByFrequencyUDP1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRealTimeByFrequencyUdp1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRealTimeByFrequencyUdp1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIRealTimeByFrequencyUdp1OK creates a InterCityBusAPIRealTimeByFrequencyUdp1OK with default headers values
func NewInterCityBusAPIRealTimeByFrequencyUdp1OK() *InterCityBusAPIRealTimeByFrequencyUdp1OK {
	return &InterCityBusAPIRealTimeByFrequencyUdp1OK{}
}

/*InterCityBusAPIRealTimeByFrequencyUdp1OK handles this case with default header values.

OK
*/
type InterCityBusAPIRealTimeByFrequencyUdp1OK struct {
	Payload []*models.ServiceDTOVersion2BusBusA1Data
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/InterCity/{RouteName}][%d] interCityBusApiRealTimeByFrequencyUdp1OK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp1OK) GetPayload() []*models.ServiceDTOVersion2BusBusA1Data {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeByFrequencyUdp1Status299 creates a InterCityBusAPIRealTimeByFrequencyUdp1Status299 with default headers values
func NewInterCityBusAPIRealTimeByFrequencyUdp1Status299() *InterCityBusAPIRealTimeByFrequencyUdp1Status299 {
	return &InterCityBusAPIRealTimeByFrequencyUdp1Status299{}
}

/*InterCityBusAPIRealTimeByFrequencyUdp1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRealTimeByFrequencyUdp1Status299 struct {
	Payload *models.ServiceDTOVersion3BaseHealth
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/Streaming/InterCity/{RouteName}][%d] interCityBusApiRealTimeByFrequencyUdp1Status299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp1Status299) GetPayload() *models.ServiceDTOVersion3BaseHealth {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeByFrequencyUdp1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
