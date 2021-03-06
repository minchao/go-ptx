// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v3/models"
)

// StationFacilityAPIControllerGetReader is a Reader for the StationFacilityAPIControllerGet structure.
type StationFacilityAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StationFacilityAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStationFacilityAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStationFacilityAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStationFacilityAPIControllerGetOK creates a StationFacilityAPIControllerGetOK with default headers values
func NewStationFacilityAPIControllerGetOK() *StationFacilityAPIControllerGetOK {
	return &StationFacilityAPIControllerGetOK{}
}

/* StationFacilityAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type StationFacilityAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility
}

func (o *StationFacilityAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationFacility][%d] stationFacilityApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *StationFacilityAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility {
	return o.Payload
}

func (o *StationFacilityAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationFacilityStationFacility)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationFacilityAPIControllerGetNotModified creates a StationFacilityAPIControllerGetNotModified with default headers values
func NewStationFacilityAPIControllerGetNotModified() *StationFacilityAPIControllerGetNotModified {
	return &StationFacilityAPIControllerGetNotModified{}
}

/* StationFacilityAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type StationFacilityAPIControllerGetNotModified struct {
}

func (o *StationFacilityAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationFacility][%d] stationFacilityApiControllerGetNotModified ", 304)
}

func (o *StationFacilityAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
