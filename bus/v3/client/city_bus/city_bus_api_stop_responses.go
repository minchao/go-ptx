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

// CityBusAPIStopReader is a Reader for the CityBusAPIStop structure.
type CityBusAPIStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIStopStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIStopOK creates a CityBusAPIStopOK with default headers values
func NewCityBusAPIStopOK() *CityBusAPIStopOK {
	return &CityBusAPIStopOK{}
}

/*CityBusAPIStopOK handles this case with default header values.

OK
*/
type CityBusAPIStopOK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusStop
}

func (o *CityBusAPIStopOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Stop/City/{City}][%d] cityBusApiStopOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIStopOK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusStop {
	return o.Payload
}

func (o *CityBusAPIStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusStop)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIStopStatus299 creates a CityBusAPIStopStatus299 with default headers values
func NewCityBusAPIStopStatus299() *CityBusAPIStopStatus299 {
	return &CityBusAPIStopStatus299{}
}

/*CityBusAPIStopStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIStopStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIStopStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Stop/City/{City}][%d] cityBusApiStopStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIStopStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIStopStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
