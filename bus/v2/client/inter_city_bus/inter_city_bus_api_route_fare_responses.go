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

// InterCityBusAPIRouteFareReader is a Reader for the InterCityBusAPIRouteFare structure.
type InterCityBusAPIRouteFareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIRouteFareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIRouteFareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIRouteFareStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIRouteFareOK creates a InterCityBusAPIRouteFareOK with default headers values
func NewInterCityBusAPIRouteFareOK() *InterCityBusAPIRouteFareOK {
	return &InterCityBusAPIRouteFareOK{}
}

/*InterCityBusAPIRouteFareOK handles this case with default header values.

OK
*/
type InterCityBusAPIRouteFareOK struct {
	Payload []*models.ServiceDTOVersion2BusBusRouteFare
}

func (o *InterCityBusAPIRouteFareOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/InterCity][%d] interCityBusApiRouteFareOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIRouteFareOK) GetPayload() []*models.ServiceDTOVersion2BusBusRouteFare {
	return o.Payload
}

func (o *InterCityBusAPIRouteFareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIRouteFareStatus299 creates a InterCityBusAPIRouteFareStatus299 with default headers values
func NewInterCityBusAPIRouteFareStatus299() *InterCityBusAPIRouteFareStatus299 {
	return &InterCityBusAPIRouteFareStatus299{}
}

/*InterCityBusAPIRouteFareStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIRouteFareStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *InterCityBusAPIRouteFareStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/InterCity][%d] interCityBusApiRouteFareStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIRouteFareStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIRouteFareStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
