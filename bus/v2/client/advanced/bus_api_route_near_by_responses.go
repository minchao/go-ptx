// Code generated by go-swagger; DO NOT EDIT.

package advanced

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// BusAPIRouteNearByReader is a Reader for the BusAPIRouteNearBy structure.
type BusAPIRouteNearByReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BusAPIRouteNearByReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBusAPIRouteNearByOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewBusAPIRouteNearByStatus299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewBusAPIRouteNearByNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBusAPIRouteNearByOK creates a BusAPIRouteNearByOK with default headers values
func NewBusAPIRouteNearByOK() *BusAPIRouteNearByOK {
	return &BusAPIRouteNearByOK{}
}

/* BusAPIRouteNearByOK describes a response with status code 200, with default header values.

Success
*/
type BusAPIRouteNearByOK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRoute
}

func (o *BusAPIRouteNearByOK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/NearBy][%d] busApiRouteNearByOK  %+v", 200, o.Payload)
}
func (o *BusAPIRouteNearByOK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRoute {
	return o.Payload
}

func (o *BusAPIRouteNearByOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRouteNearByStatus299 creates a BusAPIRouteNearByStatus299 with default headers values
func NewBusAPIRouteNearByStatus299() *BusAPIRouteNearByStatus299 {
	return &BusAPIRouteNearByStatus299{}
}

/* BusAPIRouteNearByStatus299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type BusAPIRouteNearByStatus299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *BusAPIRouteNearByStatus299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/NearBy][%d] busApiRouteNearByStatus299  %+v", 299, o.Payload)
}
func (o *BusAPIRouteNearByStatus299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *BusAPIRouteNearByStatus299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBusAPIRouteNearByNotModified creates a BusAPIRouteNearByNotModified with default headers values
func NewBusAPIRouteNearByNotModified() *BusAPIRouteNearByNotModified {
	return &BusAPIRouteNearByNotModified{}
}

/* BusAPIRouteNearByNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type BusAPIRouteNearByNotModified struct {
}

func (o *BusAPIRouteNearByNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/Route/NearBy][%d] busApiRouteNearByNotModified ", 304)
}

func (o *BusAPIRouteNearByNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}