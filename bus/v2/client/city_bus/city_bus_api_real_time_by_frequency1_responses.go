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

// CityBusAPIRealTimeByFrequency1Reader is a Reader for the CityBusAPIRealTimeByFrequency1 structure.
type CityBusAPIRealTimeByFrequency1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeByFrequency1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeByFrequency1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeByFrequency1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRealTimeByFrequency1OK creates a CityBusAPIRealTimeByFrequency1OK with default headers values
func NewCityBusAPIRealTimeByFrequency1OK() *CityBusAPIRealTimeByFrequency1OK {
	return &CityBusAPIRealTimeByFrequency1OK{}
}

/*CityBusAPIRealTimeByFrequency1OK handles this case with default header values.

Success
*/
type CityBusAPIRealTimeByFrequency1OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusA1Data
}

func (o *CityBusAPIRealTimeByFrequency1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/City/{City}/{RouteName}][%d] cityBusApiRealTimeByFrequency1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRealTimeByFrequency1OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusA1Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequency1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeByFrequency1Status299 creates a CityBusAPIRealTimeByFrequency1Status299 with default headers values
func NewCityBusAPIRealTimeByFrequency1Status299() *CityBusAPIRealTimeByFrequency1Status299 {
	return &CityBusAPIRealTimeByFrequency1Status299{}
}

/*CityBusAPIRealTimeByFrequency1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeByFrequency1Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeByFrequency1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RealTimeByFrequency/City/{City}/{RouteName}][%d] cityBusApiRealTimeByFrequency1Status299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRealTimeByFrequency1Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequency1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
