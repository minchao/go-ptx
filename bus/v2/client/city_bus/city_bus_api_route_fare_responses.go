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

// CityBusAPIRouteFareReader is a Reader for the CityBusAPIRouteFare structure.
type CityBusAPIRouteFareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteFareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteFareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRouteFareStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRouteFareOK creates a CityBusAPIRouteFareOK with default headers values
func NewCityBusAPIRouteFareOK() *CityBusAPIRouteFareOK {
	return &CityBusAPIRouteFareOK{}
}

/*CityBusAPIRouteFareOK handles this case with default header values.

Success
*/
type CityBusAPIRouteFareOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRouteFare
}

func (o *CityBusAPIRouteFareOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}][%d] cityBusApiRouteFareOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRouteFareOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRouteFare {
	return o.Payload
}

func (o *CityBusAPIRouteFareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteFareStatus299 creates a CityBusAPIRouteFareStatus299 with default headers values
func NewCityBusAPIRouteFareStatus299() *CityBusAPIRouteFareStatus299 {
	return &CityBusAPIRouteFareStatus299{}
}

/*CityBusAPIRouteFareStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRouteFareStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRouteFareStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}][%d] cityBusApiRouteFareStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRouteFareStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRouteFareStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
