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

// CityBusAPIShape1Reader is a Reader for the CityBusAPIShape1 structure.
type CityBusAPIShape1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIShape1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIShape1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIShape1Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIShape1OK creates a CityBusAPIShape1OK with default headers values
func NewCityBusAPIShape1OK() *CityBusAPIShape1OK {
	return &CityBusAPIShape1OK{}
}

/*CityBusAPIShape1OK handles this case with default header values.

OK
*/
type CityBusAPIShape1OK struct {
	Payload []*models.ServiceDTOVersion2BusBusShape
}

func (o *CityBusAPIShape1OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}/{RouteName}][%d] cityBusApiShape1OK  %+v", 200, o.Payload)
}

func (o *CityBusAPIShape1OK) GetPayload() []*models.ServiceDTOVersion2BusBusShape {
	return o.Payload
}

func (o *CityBusAPIShape1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIShape1Status299 creates a CityBusAPIShape1Status299 with default headers values
func NewCityBusAPIShape1Status299() *CityBusAPIShape1Status299 {
	return &CityBusAPIShape1Status299{}
}

/*CityBusAPIShape1Status299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIShape1Status299 struct {
	Payload *models.ServiceDTOVersion3BaseDisplayHealth
}

func (o *CityBusAPIShape1Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}/{RouteName}][%d] cityBusApiShape1Status299  %+v", 299, o.Payload)
}

func (o *CityBusAPIShape1Status299) GetPayload() *models.ServiceDTOVersion3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIShape1Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
