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

// InterCityBusAPIStopReader is a Reader for the InterCityBusAPIStop structure.
type InterCityBusAPIStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterCityBusAPIStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterCityBusAPIStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewInterCityBusAPIStopStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInterCityBusAPIStopOK creates a InterCityBusAPIStopOK with default headers values
func NewInterCityBusAPIStopOK() *InterCityBusAPIStopOK {
	return &InterCityBusAPIStopOK{}
}

/*InterCityBusAPIStopOK handles this case with default header values.

OK
*/
type InterCityBusAPIStopOK struct {
	Payload []*models.ServiceDTOVersion2BusBusStop
}

func (o *InterCityBusAPIStopOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity][%d] interCityBusApiStopOK  %+v", 200, o.Payload)
}

func (o *InterCityBusAPIStopOK) GetPayload() []*models.ServiceDTOVersion2BusBusStop {
	return o.Payload
}

func (o *InterCityBusAPIStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterCityBusAPIStopStatus299 creates a InterCityBusAPIStopStatus299 with default headers values
func NewInterCityBusAPIStopStatus299() *InterCityBusAPIStopStatus299 {
	return &InterCityBusAPIStopStatus299{}
}

/*InterCityBusAPIStopStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type InterCityBusAPIStopStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *InterCityBusAPIStopStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Stop/InterCity][%d] interCityBusApiStopStatus299  %+v", 299, o.Payload)
}

func (o *InterCityBusAPIStopStatus299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *InterCityBusAPIStopStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
