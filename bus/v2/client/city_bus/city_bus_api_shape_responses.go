// Code generated by go-swagger; DO NOT EDIT.

package city_bus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIShapeReader is a Reader for the CityBusAPIShape structure.
type CityBusAPIShapeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIShapeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIShapeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIShapeStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCityBusAPIShapeOK creates a CityBusAPIShapeOK with default headers values
func NewCityBusAPIShapeOK() *CityBusAPIShapeOK {
	return &CityBusAPIShapeOK{}
}

/*CityBusAPIShapeOK handles this case with default header values.

OK
*/
type CityBusAPIShapeOK struct {
	Payload []*models.ServiceDTOVersion2BusBusShape
}

func (o *CityBusAPIShapeOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}][%d] cityBusApiShapeOK  %+v", 200, o.Payload)
}

func (o *CityBusAPIShapeOK) GetPayload() []*models.ServiceDTOVersion2BusBusShape {
	return o.Payload
}

func (o *CityBusAPIShapeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIShapeStatus299 creates a CityBusAPIShapeStatus299 with default headers values
func NewCityBusAPIShapeStatus299() *CityBusAPIShapeStatus299 {
	return &CityBusAPIShapeStatus299{}
}

/*CityBusAPIShapeStatus299 handles this case with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIShapeStatus299 struct {
	Payload *models.ServiceDTOVersion3BaseHealth
}

func (o *CityBusAPIShapeStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}][%d] cityBusApiShapeStatus299  %+v", 299, o.Payload)
}

func (o *CityBusAPIShapeStatus299) GetPayload() *models.ServiceDTOVersion3BaseHealth {
	return o.Payload
}

func (o *CityBusAPIShapeStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDTOVersion3BaseHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
