// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v3/models"
)

// TrainTypeAPIControllerGetReader is a Reader for the TrainTypeAPIControllerGet structure.
type TrainTypeAPIControllerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrainTypeAPIControllerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTrainTypeAPIControllerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTrainTypeAPIControllerGetOK creates a TrainTypeAPIControllerGetOK with default headers values
func NewTrainTypeAPIControllerGetOK() *TrainTypeAPIControllerGetOK {
	return &TrainTypeAPIControllerGetOK{}
}

/*TrainTypeAPIControllerGetOK handles this case with default header values.

OK
*/
type TrainTypeAPIControllerGetOK struct {
	Payload *models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRATrainType
}

func (o *TrainTypeAPIControllerGetOK) Error() string {
	return fmt.Sprintf("[GET /v3/Rail/TRA/TrainType][%d] trainTypeApiControllerGetOK  %+v", 200, o.Payload)
}

func (o *TrainTypeAPIControllerGetOK) GetPayload() *models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRATrainType {
	return o.Payload
}

func (o *TrainTypeAPIControllerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MOTCAPIRailModelsTRABaseWrapperServiceDTOVersion3RailTRATrainType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
