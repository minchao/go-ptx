// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIShapeByStation2885Reader is a Reader for the CityBusAPIShapeByStation2885 structure.
type CityBusAPIShapeByStation2885Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIShapeByStation2885Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIShapeByStation2885OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIShapeByStation2885Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIShapeByStation2885NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIShapeByStation2885OK creates a CityBusAPIShapeByStation2885OK with default headers values
func NewCityBusAPIShapeByStation2885OK() *CityBusAPIShapeByStation2885OK {
	return &CityBusAPIShapeByStation2885OK{}
}

/* CityBusAPIShapeByStation2885OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIShapeByStation2885OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusShape
}

func (o *CityBusAPIShapeByStation2885OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}/PassThrough/Station/{StationID}][%d] cityBusApiShapeByStation2885OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIShapeByStation2885OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusShape {
	return o.Payload
}

func (o *CityBusAPIShapeByStation2885OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIShapeByStation2885Status299 creates a CityBusAPIShapeByStation2885Status299 with default headers values
func NewCityBusAPIShapeByStation2885Status299() *CityBusAPIShapeByStation2885Status299 {
	return &CityBusAPIShapeByStation2885Status299{}
}

/* CityBusAPIShapeByStation2885Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIShapeByStation2885Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIShapeByStation2885Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}/PassThrough/Station/{StationID}][%d] cityBusApiShapeByStation2885Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIShapeByStation2885Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIShapeByStation2885Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIShapeByStation2885NotModified creates a CityBusAPIShapeByStation2885NotModified with default headers values
func NewCityBusAPIShapeByStation2885NotModified() *CityBusAPIShapeByStation2885NotModified {
	return &CityBusAPIShapeByStation2885NotModified{}
}

/* CityBusAPIShapeByStation2885NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIShapeByStation2885NotModified struct {
}

func (o *CityBusAPIShapeByStation2885NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Shape/City/{City}/PassThrough/Station/{StationID}][%d] cityBusApiShapeByStation2885NotModified ", 304)
}

func (o *CityBusAPIShapeByStation2885NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}