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

// THSRAPIStationExitReader is a Reader for the THSRAPIStationExit structure.
type THSRAPIStationExitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *THSRAPIStationExitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTHSRAPIStationExitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTHSRAPIStationExitOK creates a THSRAPIStationExitOK with default headers values
func NewTHSRAPIStationExitOK() *THSRAPIStationExitOK {
	return &THSRAPIStationExitOK{}
}

/*THSRAPIStationExitOK handles this case with default header values.

Success
*/
type THSRAPIStationExitOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2THSRStationExit
}

func (o *THSRAPIStationExitOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/THSR/StationExit][%d] tHSRApiStationExitOK  %+v", 200, o.Payload)
}

func (o *THSRAPIStationExitOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2THSRStationExit {
	return o.Payload
}

func (o *THSRAPIStationExitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
