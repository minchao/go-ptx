// Code generated by go-swagger; DO NOT EDIT.

package t_h_s_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// THSRAPIDailyTrainInfoReader is a Reader for the THSRAPIDailyTrainInfo structure.
type THSRAPIDailyTrainInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIDailyTrainInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIDailyTrainInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIDailyTrainInfoOK creates a THSRAPIDailyTrainInfoOK with default headers values
func NewTHSRAPIDailyTrainInfoOK() *THSRAPIDailyTrainInfoOK {
	return &THSRAPIDailyTrainInfoOK{}
}

/*THSRAPIDailyTrainInfoOK handles this case with default header values.

OK
*/
type THSRAPIDailyTrainInfoOK struct {
	Payload []*models.ServiceDTOVersion2RailTHSRRailDailyTrainInfo
}

func (o *THSRAPIDailyTrainInfoOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/DailyTrainInfo/Today][%d] tHSRApiDailyTrainInfoOK  %+v", 200, o.Payload)
}

func (o *THSRAPIDailyTrainInfoOK) GetPayload() []*models.ServiceDTOVersion2RailTHSRRailDailyTrainInfo {
	return o.Payload
}

func (o *THSRAPIDailyTrainInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
