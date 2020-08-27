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

// CityBusAPIAlertReader is a Reader for the CityBusAPIAlert structure.
type CityBusAPIAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIAlertStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIAlertOK creates a CityBusAPIAlertOK with default headers values
func NewCityBusAPIAlertOK() *CityBusAPIAlertOK {
	return &CityBusAPIAlertOK{}
}

/*CityBusAPIAlertOK handles this case with default header values.

OK
*/
type CityBusAPIAlertOK struct {
	Payload *models.MOTCAPIBusDALBusServiceDTOVersion3BusAlert
}

func (o *CityBusAPIAlertOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Alert/City/{City}][%d] cityBusApiAlertOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIAlertOK) GetPayload() *models.MOTCAPIBusDALBusServiceDTOVersion3BusAlert {
	return o.Payload
}

func (o *CityBusAPIAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIBusDALBusServiceDTOVersion3BusAlert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIAlertStatus299 creates a CityBusAPIAlertStatus299 with default headers values
func NewCityBusAPIAlertStatus299() *CityBusAPIAlertStatus299 {
	return &CityBusAPIAlertStatus299{}
}

/*CityBusAPIAlertStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIAlertStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIAlertStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/Alert/City/{City}][%d] cityBusApiAlertStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIAlertStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIAlertStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
