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

// StationExitAPIControllerGetReader is a Reader for the StationExitAPIControllerGet structure.
type StationExitAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StationExitAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStationExitAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStationExitAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStationExitAPIControllerGetOK creates a StationExitAPIControllerGetOK with default headers values
func NewStationExitAPIControllerGetOK() *StationExitAPIControllerGetOK {
	return &StationExitAPIControllerGetOK{}
}

/* StationExitAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type StationExitAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationExitStationExit
}

func (o *StationExitAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationExit][%d] stationExitApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *StationExitAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationExitStationExit {
	return o.Payload
}

func (o *StationExitAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationExitStationExit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationExitAPIControllerGetNotModified creates a StationExitAPIControllerGetNotModified with default headers values
func NewStationExitAPIControllerGetNotModified() *StationExitAPIControllerGetNotModified {
	return &StationExitAPIControllerGetNotModified{}
}

/* StationExitAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type StationExitAPIControllerGetNotModified struct {
}

func (o *StationExitAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationExit][%d] stationExitApiControllerGetNotModified ", 304)
}

func (o *StationExitAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
