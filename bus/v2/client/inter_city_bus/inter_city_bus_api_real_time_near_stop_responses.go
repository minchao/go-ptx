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

// InterCityBusAPIRealTimeNearStopReader is a Reader for the InterCityBusAPIRealTimeNearStop structure.
type InterCityBusAPIRealTimeNearStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRealTimeNearStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRealTimeNearStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRealTimeNearStopStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIRealTimeNearStopOK creates a InterCityBusAPIRealTimeNearStopOK with default headers values
func NewInterCityBusAPIRealTimeNearStopOK() *InterCityBusAPIRealTimeNearStopOK {
	return &InterCityBusAPIRealTimeNearStopOK{}
}

/*InterCityBusAPIRealTimeNearStopOK handles this case with default header values.

OK
*/
type InterCityBusAPIRealTimeNearStopOK struct {
	Payload []*models.ServiceDTOVersion2BusBusA2Data
}

func (o *InterCityBusAPIRealTimeNearStopOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/InterCity][%d] interCityBusApiRealTimeNearStopOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIRealTimeNearStopOK) GetPayload() []*models.ServiceDTOVersion2BusBusA2Data {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeNearStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRealTimeNearStopStatus299 creates a InterCityBusAPIRealTimeNearStopStatus299 with default headers values
func NewInterCityBusAPIRealTimeNearStopStatus299() *InterCityBusAPIRealTimeNearStopStatus299 {
	return &InterCityBusAPIRealTimeNearStopStatus299{}
}

/*InterCityBusAPIRealTimeNearStopStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRealTimeNearStopStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseHealth
}

func (o *InterCityBusAPIRealTimeNearStopStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeNearStop/InterCity][%d] interCityBusApiRealTimeNearStopStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIRealTimeNearStopStatus299) GetPayload() *models.ServiceDTOVersion3BaseHealth {
	return o.Payload
}

func (o *InterCityBusAPIRealTimeNearStopStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
