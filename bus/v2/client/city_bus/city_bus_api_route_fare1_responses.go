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

// CityBusAPIRouteFare1Reader is a Reader for the CityBusAPIRouteFare1 structure.
type CityBusAPIRouteFare1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteFare1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteFare1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRouteFare1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIRouteFare1OK creates a CityBusAPIRouteFare1OK with default headers values
func NewCityBusAPIRouteFare1OK() *CityBusAPIRouteFare1OK {
	return &CityBusAPIRouteFare1OK{}
}

/*CityBusAPIRouteFare1OK handles this case with default header values.

OK
*/
type CityBusAPIRouteFare1OK struct {
	Payload []*models.ServiceDTOVersion2BusBusRouteFare
}

func (o *CityBusAPIRouteFare1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}/{RouteName}][%d] cityBusApiRouteFare1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIRouteFare1OK) GetPayload() []*models.ServiceDTOVersion2BusBusRouteFare {
	return o.Payload
}

func (o *CityBusAPIRouteFare1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteFare1Status299 creates a CityBusAPIRouteFare1Status299 with default headers values
func NewCityBusAPIRouteFare1Status299() *CityBusAPIRouteFare1Status299 {
	return &CityBusAPIRouteFare1Status299{}
}

/*CityBusAPIRouteFare1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRouteFare1Status299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIRouteFare1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}/{RouteName}][%d] cityBusApiRouteFare1Status299  %+v", 299, o.Payload)
}

func (o *CityBusAPIRouteFare1Status299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRouteFare1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
