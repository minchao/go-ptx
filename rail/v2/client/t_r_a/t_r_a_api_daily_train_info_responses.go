// Code generated by go-swagger; DO NOT EDIT.

package t_r_a

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/minchao/go-ptx/rail/v2/models"
)

// TRAAPIDailyTrainInfoReader is a Reader for the TRAAPIDailyTrainInfo structure.
type TRAAPIDailyTrainInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TRAAPIDailyTrainInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTRAAPIDailyTrainInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTRAAPIDailyTrainInfoOK creates a TRAAPIDailyTrainInfoOK with default headers values
func NewTRAAPIDailyTrainInfoOK() *TRAAPIDailyTrainInfoOK {
	return &TRAAPIDailyTrainInfoOK{}
}

/*TRAAPIDailyTrainInfoOK handles this case with default header values.

OK
*/
type TRAAPIDailyTrainInfoOK struct {
	Payload []*models.ServiceDTOVersion2RailTRARailDailyTrainInfo
}

func (o *TRAAPIDailyTrainInfoOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/TRA/DailyTrainInfo/Today][%d] tRAApiDailyTrainInfoOK  %+v", 200, o.Payload)
}

func (o *TRAAPIDailyTrainInfoOK) GetPayload() []*models.ServiceDTOVersion2RailTRARailDailyTrainInfo {
	return o.Payload
}

func (o *TRAAPIDailyTrainInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
