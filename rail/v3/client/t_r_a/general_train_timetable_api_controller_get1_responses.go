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

// GeneralTrainTimetableAPIControllerGet1Reader is a Reader for the GeneralTrainTimetableAPIControllerGet1 structure.
type GeneralTrainTimetableAPIControllerGet1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeneralTrainTimetableAPIControllerGet1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeneralTrainTimetableAPIControllerGet1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGeneralTrainTimetableAPIControllerGet1NotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeneralTrainTimetableAPIControllerGet1OK creates a GeneralTrainTimetableAPIControllerGet1OK with default headers values
func NewGeneralTrainTimetableAPIControllerGet1OK() *GeneralTrainTimetableAPIControllerGet1OK {
	return &GeneralTrainTimetableAPIControllerGet1OK{}
}

/* GeneralTrainTimetableAPIControllerGet1OK describes a response with status code 200, with default header values.

Success
*/
type GeneralTrainTimetableAPIControllerGet1OK struct {
	Payload *models.PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable
}

func (o *GeneralTrainTimetableAPIControllerGet1OK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/GeneralTrainTimetable/TrainNo/{TrainNo}][%d] generalTrainTimetableApiControllerGet1OK  %+v", 200, o.Payload)
}
func (o *GeneralTrainTimetableAPIControllerGet1OK) GetPayload() *models.PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable {
	return o.Payload
}

func (o *GeneralTrainTimetableAPIControllerGet1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneralTrainTimetableAPIControllerGet1NotModified creates a GeneralTrainTimetableAPIControllerGet1NotModified with default headers values
func NewGeneralTrainTimetableAPIControllerGet1NotModified() *GeneralTrainTimetableAPIControllerGet1NotModified {
	return &GeneralTrainTimetableAPIControllerGet1NotModified{}
}

/* GeneralTrainTimetableAPIControllerGet1NotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type GeneralTrainTimetableAPIControllerGet1NotModified struct {
}

func (o *GeneralTrainTimetableAPIControllerGet1NotModified) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/GeneralTrainTimetable/TrainNo/{TrainNo}][%d] generalTrainTimetableApiControllerGet1NotModified ", 304)
}

func (o *GeneralTrainTimetableAPIControllerGet1NotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
