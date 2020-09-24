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

// GeneralStationTimetableAPIControllerGet1Reader is a Reader for the GeneralStationTimetableAPIControllerGet1 structure.
type GeneralStationTimetableAPIControllerGet1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeneralStationTimetableAPIControllerGet1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeneralStationTimetableAPIControllerGet1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeneralStationTimetableAPIControllerGet1OK creates a GeneralStationTimetableAPIControllerGet1OK with default headers values
func NewGeneralStationTimetableAPIControllerGet1OK() *GeneralStationTimetableAPIControllerGet1OK {
	return &GeneralStationTimetableAPIControllerGet1OK{}
}

/*GeneralStationTimetableAPIControllerGet1OK handles this case with default header values.

Success
*/
type GeneralStationTimetableAPIControllerGet1OK struct {
	Payload *models.PTXAPIRailModelTRAGeneralStationWrapperPTXServiceDTORailSpecificationV3TRAGeneralStationTimetableGeneralStationTimetable
}

func (o *GeneralStationTimetableAPIControllerGet1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/GeneralStationTimetable/Station/{StationID}][%d] generalStationTimetableApiControllerGet1OK  %+v", 200, o.Payload)
}

func (o *GeneralStationTimetableAPIControllerGet1OK) GetPayload() *models.PTXAPIRailModelTRAGeneralStationWrapperPTXServiceDTORailSpecificationV3TRAGeneralStationTimetableGeneralStationTimetable {
	return o.Payload
}

func (o *GeneralStationTimetableAPIControllerGet1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRAGeneralStationWrapperPTXServiceDTORailSpecificationV3TRAGeneralStationTimetableGeneralStationTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
