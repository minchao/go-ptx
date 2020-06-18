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

// GeneralStationTimetableAPIControllerGetReader is a Reader for the GeneralStationTimetableAPIControllerGet structure.
type GeneralStationTimetableAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeneralStationTimetableAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeneralStationTimetableAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGeneralStationTimetableAPIControllerGetOK creates a GeneralStationTimetableAPIControllerGetOK with default headers values
func NewGeneralStationTimetableAPIControllerGetOK() *GeneralStationTimetableAPIControllerGetOK {
	return &GeneralStationTimetableAPIControllerGetOK{}
}

/*GeneralStationTimetableAPIControllerGetOK handles this case with default header values.

OK
*/
type GeneralStationTimetableAPIControllerGetOK struct {
	Payload *models.MOTCAPIRailModelsTRAGeneralStationWrapperServiceDTOVersion3RailTRAGeneralStationTimetableGeneralStationTimetable
}

func (o *GeneralStationTimetableAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/GeneralStationTimetable][%d] generalStationTimetableApiControllerGetOK  %+v", 200, o.Payload)
}

func (o *GeneralStationTimetableAPIControllerGetOK) GetPayload() *models.MOTCAPIRailModelsTRAGeneralStationWrapperServiceDTOVersion3RailTRAGeneralStationTimetableGeneralStationTimetable {
	return o.Payload
}

func (o *GeneralStationTimetableAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTRAGeneralStationWrapperServiceDTOVersion3RailTRAGeneralStationTimetableGeneralStationTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
