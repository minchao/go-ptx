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

// CityBusAPIScheduleReader is a Reader for the CityBusAPISchedule structure.
type CityBusAPIScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIScheduleStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIScheduleOK creates a CityBusAPIScheduleOK with default headers values
func NewCityBusAPIScheduleOK() *CityBusAPIScheduleOK {
	return &CityBusAPIScheduleOK{}
}

/*CityBusAPIScheduleOK handles this case with default header values.

OK
*/
type CityBusAPIScheduleOK struct {
	Payload []*models.ServiceDTOVersion2BusBusSchedule
}

func (o *CityBusAPIScheduleOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/City/{City}][%d] cityBusApiScheduleOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIScheduleOK) GetPayload() []*models.ServiceDTOVersion2BusBusSchedule {
	return o.Payload
}

func (o *CityBusAPIScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIScheduleStatus299 creates a CityBusAPIScheduleStatus299 with default headers values
func NewCityBusAPIScheduleStatus299() *CityBusAPIScheduleStatus299 {
	return &CityBusAPIScheduleStatus299{}
}

/*CityBusAPIScheduleStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIScheduleStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIScheduleStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/City/{City}][%d] cityBusApiScheduleStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIScheduleStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIScheduleStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
