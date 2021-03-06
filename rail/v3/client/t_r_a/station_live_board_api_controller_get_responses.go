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

// StationLiveBoardAPIControllerGetReader is a Reader for the StationLiveBoardAPIControllerGet structure.
type StationLiveBoardAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StationLiveBoardAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStationLiveBoardAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStationLiveBoardAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStationLiveBoardAPIControllerGetOK creates a StationLiveBoardAPIControllerGetOK with default headers values
func NewStationLiveBoardAPIControllerGetOK() *StationLiveBoardAPIControllerGetOK {
	return &StationLiveBoardAPIControllerGetOK{}
}

/* StationLiveBoardAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type StationLiveBoardAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard
}

func (o *StationLiveBoardAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationLiveBoard][%d] stationLiveBoardApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *StationLiveBoardAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard {
	return o.Payload
}

func (o *StationLiveBoardAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationLiveBoardAPIControllerGetNotModified creates a StationLiveBoardAPIControllerGetNotModified with default headers values
func NewStationLiveBoardAPIControllerGetNotModified() *StationLiveBoardAPIControllerGetNotModified {
	return &StationLiveBoardAPIControllerGetNotModified{}
}

/* StationLiveBoardAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type StationLiveBoardAPIControllerGetNotModified struct {
}

func (o *StationLiveBoardAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationLiveBoard][%d] stationLiveBoardApiControllerGetNotModified ", 304)
}

func (o *StationLiveBoardAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
