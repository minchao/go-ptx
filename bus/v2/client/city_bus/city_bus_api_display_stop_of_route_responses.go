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

// CityBusAPIDisplayStopOfRouteReader is a Reader for the CityBusAPIDisplayStopOfRoute structure.
type CityBusAPIDisplayStopOfRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIDisplayStopOfRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIDisplayStopOfRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIDisplayStopOfRouteStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIDisplayStopOfRouteOK creates a CityBusAPIDisplayStopOfRouteOK with default headers values
func NewCityBusAPIDisplayStopOfRouteOK() *CityBusAPIDisplayStopOfRouteOK {
	return &CityBusAPIDisplayStopOfRouteOK{}
}

/*CityBusAPIDisplayStopOfRouteOK handles this case with default header values.

Success
*/
type CityBusAPIDisplayStopOfRouteOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusDisplayStopOfRoute
}

func (o *CityBusAPIDisplayStopOfRouteOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DisplayStopOfRoute/City/{City}][%d] cityBusApiDisplayStopOfRouteOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIDisplayStopOfRouteOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusDisplayStopOfRoute {
	return o.Payload
}

func (o *CityBusAPIDisplayStopOfRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIDisplayStopOfRouteStatus299 creates a CityBusAPIDisplayStopOfRouteStatus299 with default headers values
func NewCityBusAPIDisplayStopOfRouteStatus299() *CityBusAPIDisplayStopOfRouteStatus299 {
	return &CityBusAPIDisplayStopOfRouteStatus299{}
}

/*CityBusAPIDisplayStopOfRouteStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIDisplayStopOfRouteStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIDisplayStopOfRouteStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DisplayStopOfRoute/City/{City}][%d] cityBusApiDisplayStopOfRouteStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIDisplayStopOfRouteStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIDisplayStopOfRouteStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
