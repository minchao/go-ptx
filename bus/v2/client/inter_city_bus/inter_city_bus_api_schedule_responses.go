// Code generated by go-swagger; DO NOT EDIT.

package inter_city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/v2/models"
)

// InterCityBusAPIScheduleReader is a Reader for the InterCityBusAPISchedule structure.
type InterCityBusAPIScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIScheduleStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIScheduleOK creates a InterCityBusAPIScheduleOK with default headers values
func NewInterCityBusAPIScheduleOK() *InterCityBusAPIScheduleOK {
	return &InterCityBusAPIScheduleOK{}
}

/*InterCityBusAPIScheduleOK handles this case with default header values.

OK
*/
type InterCityBusAPIScheduleOK struct {
	Payload []*models.ServiceDTOVersion2BusBusSchedule
}

func (o *InterCityBusAPIScheduleOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiScheduleOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIScheduleOK) GetPayload() []*models.ServiceDTOVersion2BusBusSchedule {
	return o.Payload
}

func (o *InterCityBusAPIScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIScheduleStatus299 creates a InterCityBusAPIScheduleStatus299 with default headers values
func NewInterCityBusAPIScheduleStatus299() *InterCityBusAPIScheduleStatus299 {
	return &InterCityBusAPIScheduleStatus299{}
}

/*InterCityBusAPIScheduleStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIScheduleStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseHealth
}

func (o *InterCityBusAPIScheduleStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Schedule/InterCity][%d] interCityBusApiScheduleStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIScheduleStatus299) GetPayload() *models.ServiceDTOVersion3BaseHealth {
	return o.Payload
}

func (o *InterCityBusAPIScheduleStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
