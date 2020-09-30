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

// CityBusAPIRealTimeByFrequencyReader is a Reader for the CityBusAPIRealTimeByFrequency structure.
type CityBusAPIRealTimeByFrequencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRealTimeByFrequencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRealTimeByFrequencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRealTimeByFrequencyStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRealTimeByFrequencyOK creates a CityBusAPIRealTimeByFrequencyOK with default headers values
func NewCityBusAPIRealTimeByFrequencyOK() *CityBusAPIRealTimeByFrequencyOK {
	return &CityBusAPIRealTimeByFrequencyOK{}
}

/*CityBusAPIRealTimeByFrequencyOK handles this case with default header values.

Success
*/
type CityBusAPIRealTimeByFrequencyOK struct {
	Payload *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data
}

func (o *CityBusAPIRealTimeByFrequencyOK) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/RealTimeByFrequency/City/{City}][%d] cityBusApiRealTimeByFrequencyOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRealTimeByFrequencyOK) GetPayload() *models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3A1Data)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRealTimeByFrequencyStatus299 creates a CityBusAPIRealTimeByFrequencyStatus299 with default headers values
func NewCityBusAPIRealTimeByFrequencyStatus299() *CityBusAPIRealTimeByFrequencyStatus299 {
	return &CityBusAPIRealTimeByFrequencyStatus299{}
}

/*CityBusAPIRealTimeByFrequencyStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRealTimeByFrequencyStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRealTimeByFrequencyStatus299) Error() string {
	return fmt.Sprintf("[GET /v3/Bus/RealTimeByFrequency/City/{City}][%d] cityBusApiRealTimeByFrequencyStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRealTimeByFrequencyStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRealTimeByFrequencyStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}