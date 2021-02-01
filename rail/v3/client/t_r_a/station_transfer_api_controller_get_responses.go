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

// StationTransferAPIControllerGetReader is a Reader for the StationTransferAPIControllerGet structure.
type StationTransferAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StationTransferAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStationTransferAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStationTransferAPIControllerGetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStationTransferAPIControllerGetOK creates a StationTransferAPIControllerGetOK with default headers values
func NewStationTransferAPIControllerGetOK() *StationTransferAPIControllerGetOK {
	return &StationTransferAPIControllerGetOK{}
}

/* StationTransferAPIControllerGetOK describes a response with status code 200, with default header values.

Success
*/
type StationTransferAPIControllerGetOK struct {
	Payload *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationTransferStationTransfer
}

func (o *StationTransferAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationTransfer][%d] stationTransferApiControllerGetOK  %+v", 200, o.Payload)
}
func (o *StationTransferAPIControllerGetOK) GetPayload() *models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationTransferStationTransfer {
	return o.Payload
}

func (o *StationTransferAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationTransferStationTransfer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStationTransferAPIControllerGetNotModified creates a StationTransferAPIControllerGetNotModified with default headers values
func NewStationTransferAPIControllerGetNotModified() *StationTransferAPIControllerGetNotModified {
	return &StationTransferAPIControllerGetNotModified{}
}

/* StationTransferAPIControllerGetNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type StationTransferAPIControllerGetNotModified struct {
}

func (o *StationTransferAPIControllerGetNotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationTransfer][%d] stationTransferApiControllerGetNotModified ", 304)
}

func (o *StationTransferAPIControllerGetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
