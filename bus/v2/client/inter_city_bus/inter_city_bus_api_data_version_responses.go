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

// InterCityBusAPIDataVersionReader is a Reader for the InterCityBusAPIDataVersion structure.
type InterCityBusAPIDataVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIDataVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIDataVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIDataVersionStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInterCityBusAPIDataVersionOK creates a InterCityBusAPIDataVersionOK with default headers values
func NewInterCityBusAPIDataVersionOK() *InterCityBusAPIDataVersionOK {
	return &InterCityBusAPIDataVersionOK{}
}

/*InterCityBusAPIDataVersionOK handles this case with default header values.

Success
*/
type InterCityBusAPIDataVersionOK struct {
	Payload *models.PTXServiceDTOBusSpecificationV2BusVersion
}

func (o *InterCityBusAPIDataVersionOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DataVersion/InterCity][%d] interCityBusApiDataVersionOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIDataVersionOK) GetPayload() *models.PTXServiceDTOBusSpecificationV2BusVersion {
	return o.Payload
}

func (o *InterCityBusAPIDataVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV2BusVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIDataVersionStatus299 creates a InterCityBusAPIDataVersionStatus299 with default headers values
func NewInterCityBusAPIDataVersionStatus299() *InterCityBusAPIDataVersionStatus299 {
	return &InterCityBusAPIDataVersionStatus299{}
}

/*InterCityBusAPIDataVersionStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIDataVersionStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *InterCityBusAPIDataVersionStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/DataVersion/InterCity][%d] interCityBusApiDataVersionStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIDataVersionStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIDataVersionStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
