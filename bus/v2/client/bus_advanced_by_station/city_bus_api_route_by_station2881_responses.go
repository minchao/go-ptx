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

// CityBusAPIRouteByStation2881Reader is a Reader for the CityBusAPIRouteByStation2881 structure.
type CityBusAPIRouteByStation2881Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteByStation2881Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteByStation2881OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRouteByStation2881Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRouteByStation2881NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRouteByStation2881OK creates a CityBusAPIRouteByStation2881OK with default headers values
func NewCityBusAPIRouteByStation2881OK() *CityBusAPIRouteByStation2881OK {
	return &CityBusAPIRouteByStation2881OK{}
}

/* CityBusAPIRouteByStation2881OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRouteByStation2881OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *CityBusAPIRouteByStation2881OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}/PassThrough/Station/{StationID}][%d] cityBusApiRouteByStation2881OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRouteByStation2881OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *CityBusAPIRouteByStation2881OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteByStation2881Status299 creates a CityBusAPIRouteByStation2881Status299 with default headers values
func NewCityBusAPIRouteByStation2881Status299() *CityBusAPIRouteByStation2881Status299 {
	return &CityBusAPIRouteByStation2881Status299{}
}

/* CityBusAPIRouteByStation2881Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRouteByStation2881Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRouteByStation2881Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}/PassThrough/Station/{StationID}][%d] cityBusApiRouteByStation2881Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRouteByStation2881Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRouteByStation2881Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteByStation2881NotModified creates a CityBusAPIRouteByStation2881NotModified with default headers values
func NewCityBusAPIRouteByStation2881NotModified() *CityBusAPIRouteByStation2881NotModified {
	return &CityBusAPIRouteByStation2881NotModified{}
}

/* CityBusAPIRouteByStation2881NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRouteByStation2881NotModified struct {
}

func (o *CityBusAPIRouteByStation2881NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/City/{City}/PassThrough/Station/{StationID}][%d] cityBusApiRouteByStation2881NotModified ", 304)
}

func (o *CityBusAPIRouteByStation2881NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
