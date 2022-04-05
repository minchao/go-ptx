// Code generated by go-swagger; DO NOT EDIT.

package bus_advanced_by_operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/bus/v2/models"
)

// CityBusAPIRouteFareByOperator2865Reader is a Reader for the CityBusAPIRouteFareByOperator2865 structure.
type CityBusAPIRouteFareByOperator2865Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CityBusAPIRouteFareByOperator2865Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCityBusAPIRouteFareByOperator2865OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 299:
		result := NewCityBusAPIRouteFareByOperator2865Status299()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewCityBusAPIRouteFareByOperator2865NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCityBusAPIRouteFareByOperator2865OK creates a CityBusAPIRouteFareByOperator2865OK with default headers values
func NewCityBusAPIRouteFareByOperator2865OK() *CityBusAPIRouteFareByOperator2865OK {
	return &CityBusAPIRouteFareByOperator2865OK{}
}

/* CityBusAPIRouteFareByOperator2865OK describes a response with status code 200, with default header values.

Success
*/
type CityBusAPIRouteFareByOperator2865OK struct {
	Payload []*models.PTXServiceDTOBusSpecificationV2BusRouteFare
}

func (o *CityBusAPIRouteFareByOperator2865OK) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}/OperatorNo/{OperatorNo}][%d] cityBusApiRouteFareByOperator2865OK  %+v", 200, o.Payload)
}
func (o *CityBusAPIRouteFareByOperator2865OK) GetPayload() []*models.PTXServiceDTOBusSpecificationV2BusRouteFare {
	return o.Payload
}

func (o *CityBusAPIRouteFareByOperator2865OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteFareByOperator2865Status299 creates a CityBusAPIRouteFareByOperator2865Status299 with default headers values
func NewCityBusAPIRouteFareByOperator2865Status299() *CityBusAPIRouteFareByOperator2865Status299 {
	return &CityBusAPIRouteFareByOperator2865Status299{}
}

/* CityBusAPIRouteFareByOperator2865Status299 describes a response with status code 299, with default header values.

加入參數'?health=true'即可查詢此API服務的健康狀態
*/
type CityBusAPIRouteFareByOperator2865Status299 struct {
	Payload *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth
}

func (o *CityBusAPIRouteFareByOperator2865Status299) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}/OperatorNo/{OperatorNo}][%d] cityBusApiRouteFareByOperator2865Status299  %+v", 299, o.Payload)
}
func (o *CityBusAPIRouteFareByOperator2865Status299) GetPayload() *models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth {
	return o.Payload
}

func (o *CityBusAPIRouteFareByOperator2865Status299) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXServiceDTOSharedSpecificationV3BaseDisplayHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCityBusAPIRouteFareByOperator2865NotModified creates a CityBusAPIRouteFareByOperator2865NotModified with default headers values
func NewCityBusAPIRouteFareByOperator2865NotModified() *CityBusAPIRouteFareByOperator2865NotModified {
	return &CityBusAPIRouteFareByOperator2865NotModified{}
}

/* CityBusAPIRouteFareByOperator2865NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type CityBusAPIRouteFareByOperator2865NotModified struct {
}

func (o *CityBusAPIRouteFareByOperator2865NotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Bus/RouteFare/City/{City}/OperatorNo/{OperatorNo}][%d] cityBusApiRouteFareByOperator2865NotModified ", 304)
}

func (o *CityBusAPIRouteFareByOperator2865NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
