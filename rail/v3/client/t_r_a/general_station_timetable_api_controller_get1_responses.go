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
	case 304:
		result := NewGeneralStationTimetableAPIControllerGet1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeneralStationTimetableAPIControllerGet1OK creates a GeneralStationTimetableAPIControllerGet1OK with default headers values
func NewGeneralStationTimetableAPIControllerGet1OK() *GeneralStationTimetableAPIControllerGet1OK {
	return &GeneralStationTimetableAPIControllerGet1OK{}
}

/* GeneralStationTimetableAPIControllerGet1OK describes a response with status code 200, with default header values.

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

// NewGeneralStationTimetableAPIControllerGet1NotModified creates a GeneralStationTimetableAPIControllerGet1NotModified with default headers values
func NewGeneralStationTimetableAPIControllerGet1NotModified() *GeneralStationTimetableAPIControllerGet1NotModified {
	return &GeneralStationTimetableAPIControllerGet1NotModified{}
}

/* GeneralStationTimetableAPIControllerGet1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type GeneralStationTimetableAPIControllerGet1NotModified struct {
}

func (o *GeneralStationTimetableAPIControllerGet1NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/GeneralStationTimetable/Station/{StationID}][%d] generalStationTimetableApiControllerGet1NotModified ", 304)
}

func (o *GeneralStationTimetableAPIControllerGet1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
