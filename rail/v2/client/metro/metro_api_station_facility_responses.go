// Code generated by go-swagger; DO NOT EDIT.

package metro

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/models"
)

// MetroAPIStationFacilityReader is a Reader for the MetroAPIStationFacility structure.
type MetroAPIStationFacilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPIStationFacilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPIStationFacilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMetroAPIStationFacilityNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPIStationFacilityOK creates a MetroAPIStationFacilityOK with default headers values
func NewMetroAPIStationFacilityOK() *MetroAPIStationFacilityOK {
	return &MetroAPIStationFacilityOK{}
}

/* MetroAPIStationFacilityOK describes a response with status code 200, with default header values.

Success
*/
type MetroAPIStationFacilityOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroStationFacility
}

func (o *MetroAPIStationFacilityOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/StationFacility/{Operator}][%d] metroApiStationFacilityOK  %+v", 200, o.Payload)
}
func (o *MetroAPIStationFacilityOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroStationFacility {
	return o.Payload
}

func (o *MetroAPIStationFacilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPIStationFacilityNotModified creates a MetroAPIStationFacilityNotModified with default headers values
func NewMetroAPIStationFacilityNotModified() *MetroAPIStationFacilityNotModified {
	return &MetroAPIStationFacilityNotModified{}
}

/* MetroAPIStationFacilityNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type MetroAPIStationFacilityNotModified struct {
}

func (o *MetroAPIStationFacilityNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/StationFacility/{Operator}][%d] metroApiStationFacilityNotModified ", 304)
}

func (o *MetroAPIStationFacilityNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
