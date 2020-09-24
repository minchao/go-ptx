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

// InterCityBusAPIStationReader is a Reader for the InterCityBusAPIStation structure.
type InterCityBusAPIStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStationStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIStationOK creates a InterCityBusAPIStationOK with default headers values
func NewInterCityBusAPIStationOK() *InterCityBusAPIStationOK {
	return &InterCityBusAPIStationOK{}
}

/*InterCityBusAPIStationOK handles this case with default header values.

Success
*/
type InterCityBusAPIStationOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusStation
}

func (o *InterCityBusAPIStationOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Station/InterCity][%d] interCityBusApiStationOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIStationOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusStation {
	return o.Payload
}

func (o *InterCityBusAPIStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStationStatus299 creates a InterCityBusAPIStationStatus299 with default headers values
func NewInterCityBusAPIStationStatus299() *InterCityBusAPIStationStatus299 {
	return &InterCityBusAPIStationStatus299{}
}

/*InterCityBusAPIStationStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStationStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIStationStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Station/InterCity][%d] interCityBusApiStationStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIStationStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStationStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
