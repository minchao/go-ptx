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

// InterCityBusAPIRouteReader is a Reader for the InterCityBusAPIRoute structure.
type InterCityBusAPIRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRouteStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIRouteOK creates a InterCityBusAPIRouteOK with default headers values
func NewInterCityBusAPIRouteOK() *InterCityBusAPIRouteOK {
	return &InterCityBusAPIRouteOK{}
}

/*InterCityBusAPIRouteOK handles this case with default header values.

Success
*/
type InterCityBusAPIRouteOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *InterCityBusAPIRouteOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity][%d] interCityBusApiRouteOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIRouteOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *InterCityBusAPIRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteStatus299 creates a InterCityBusAPIRouteStatus299 with default headers values
func NewInterCityBusAPIRouteStatus299() *InterCityBusAPIRouteStatus299 {
	return &InterCityBusAPIRouteStatus299{}
}

/*InterCityBusAPIRouteStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRouteStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIRouteStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/InterCity][%d] interCityBusApiRouteStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIRouteStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRouteStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
