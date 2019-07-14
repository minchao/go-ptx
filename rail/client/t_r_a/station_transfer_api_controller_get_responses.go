// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/models"
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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStationTransferAPIControllerGetOK creates a StationTransferAPIControllerGetOK with default headers values
func NewStationTransferAPIControllerGetOK() *StationTransferAPIControllerGetOK {
	return &StationTransferAPIControllerGetOK{}
}

/*StationTransferAPIControllerGetOK handles this case with default header values.

OK
*/
type StationTransferAPIControllerGetOK struct {
	Payload *models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStationTransferStationTransfer
}

func (o *StationTransferAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/StationTransfer][%d] stationTransferApiControllerGetOK  %+v", 200, o.Payload)
}

func (o *StationTransferAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRAStationTransferStationTransfer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
