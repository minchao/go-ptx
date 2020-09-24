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

// THSRAPIGeneralTimetableReader is a Reader for the THSRAPIGeneralTimetable structure.
type THSRAPIGeneralTimetableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIGeneralTimetableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIGeneralTimetableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIGeneralTimetableOK creates a THSRAPIGeneralTimetableOK with default headers values
func NewTHSRAPIGeneralTimetableOK() *THSRAPIGeneralTimetableOK {
	return &THSRAPIGeneralTimetableOK{}
}

/*THSRAPIGeneralTimetableOK handles this case with default header values.

Success
*/
type THSRAPIGeneralTimetableOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRRailGeneralTimetable
}

func (o *THSRAPIGeneralTimetableOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/GeneralTimetable][%d] tHSRApiGeneralTimetableOK  %+v", 200, o.Payload)
}

func (o *THSRAPIGeneralTimetableOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRRailGeneralTimetable {
	return o.Payload
}

func (o *THSRAPIGeneralTimetableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
