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

// CityBusAPIEstimatedTimeOfArrival1Reader is a Reader for the CityBusAPIEstimatedTimeOfArrival1 structure.
type CityBusAPIEstimatedTimeOfArrival1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIEstimatedTimeOfArrival1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIEstimatedTimeOfArrival1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIEstimatedTimeOfArrival1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIEstimatedTimeOfArrival1OK creates a CityBusAPIEstimatedTimeOfArrival1OK with default headers values
func NewCityBusAPIEstimatedTimeOfArrival1OK() *CityBusAPIEstimatedTimeOfArrival1OK {
	return &CityBusAPIEstimatedTimeOfArrival1OK{}
}

/*CityBusAPIEstimatedTimeOfArrival1OK handles this case with default header values.

Success
*/
type CityBusAPIEstimatedTimeOfArrival1OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusN1EstimateTime
}

func (o *CityBusAPIEstimatedTimeOfArrival1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/City/{City}/{RouteName}][%d] cityBusApiEstimatedTimeOfArrival1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIEstimatedTimeOfArrival1OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusN1EstimateTime {
	return o.Payload
}

func (o *CityBusAPIEstimatedTimeOfArrival1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIEstimatedTimeOfArrival1Status299 creates a CityBusAPIEstimatedTimeOfArrival1Status299 with default headers values
func NewCityBusAPIEstimatedTimeOfArrival1Status299() *CityBusAPIEstimatedTimeOfArrival1Status299 {
	return &CityBusAPIEstimatedTimeOfArrival1Status299{}
}

/*CityBusAPIEstimatedTimeOfArrival1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIEstimatedTimeOfArrival1Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIEstimatedTimeOfArrival1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/EstimatedTimeOfArrival/City/{City}/{RouteName}][%d] cityBusApiEstimatedTimeOfArrival1Status299  %+v", 299, o.Payload)
}

func (o *CityBusAPIEstimatedTimeOfArrival1Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIEstimatedTimeOfArrival1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
