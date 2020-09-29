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

// CityBusAPIFirstLastTripInfoReader is a Reader for the CityBusAPIFirstLastTripInfo structure.
type CityBusAPIFirstLastTripInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIFirstLastTripInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIFirstLastTripInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIFirstLastTripInfoStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIFirstLastTripInfoOK creates a CityBusAPIFirstLastTripInfoOK with default headers values
func NewCityBusAPIFirstLastTripInfoOK() *CityBusAPIFirstLastTripInfoOK {
	return &CityBusAPIFirstLastTripInfoOK{}
}

/*CityBusAPIFirstLastTripInfoOK handles this case with default header values.

Success
*/
type CityBusAPIFirstLastTripInfoOK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3FirstLastTripInfo
}

func (o *CityBusAPIFirstLastTripInfoOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfoOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIFirstLastTripInfoOK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3FirstLastTripInfo {
	return o.Payload
}

func (o *CityBusAPIFirstLastTripInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusVWrapperPTXServiceDTOBusSpecificationV3FirstLastTripInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIFirstLastTripInfoStatus299 creates a CityBusAPIFirstLastTripInfoStatus299 with default headers values
func NewCityBusAPIFirstLastTripInfoStatus299() *CityBusAPIFirstLastTripInfoStatus299 {
	return &CityBusAPIFirstLastTripInfoStatus299{}
}

/*CityBusAPIFirstLastTripInfoStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIFirstLastTripInfoStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIFirstLastTripInfoStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/FirstLastTripInfo/City/{City}][%d] cityBusApiFirstLastTripInfoStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIFirstLastTripInfoStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIFirstLastTripInfoStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
