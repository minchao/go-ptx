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

// MetroAPILiveBoardReader is a Reader for the MetroAPILiveBoard structure.
type MetroAPILiveBoardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroAPILiveBoardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroAPILiveBoardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMetroAPILiveBoardNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetroAPILiveBoardOK creates a MetroAPILiveBoardOK with default headers values
func NewMetroAPILiveBoardOK() *MetroAPILiveBoardOK {
	return &MetroAPILiveBoardOK{}
}

/* MetroAPILiveBoardOK describes a response with status code 200, with default header values.

Success
*/
type MetroAPILiveBoardOK struct {
	Payload []*models.PTXServiceDTORailSpecificationV2MetroLiveBoard
}

func (o *MetroAPILiveBoardOK) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/LiveBoard/{Operator}][%d] metroApiLiveBoardOK  %+v", 200, o.Payload)
}
func (o *MetroAPILiveBoardOK) GetPayload() []*models.PTXServiceDTORailSpecificationV2MetroLiveBoard {
	return o.Payload
}

func (o *MetroAPILiveBoardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroAPILiveBoardNotModified creates a MetroAPILiveBoardNotModified with default headers values
func NewMetroAPILiveBoardNotModified() *MetroAPILiveBoardNotModified {
	return &MetroAPILiveBoardNotModified{}
}

/* MetroAPILiveBoardNotModified describes a response with status code 304, with default header values.

服務端會在Response加上Last-Modified header，表示最近的更新時間。客戶端能利用此時間，於Request加上If-Modified-Since header，若沒有更新，服務端會回應304 StatusCode且空值Content
*/
type MetroAPILiveBoardNotModified struct {
}

func (o *MetroAPILiveBoardNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/Rail/Metro/LiveBoard/{Operator}][%d] metroApiLiveBoardNotModified ", 304)
}

func (o *MetroAPILiveBoardNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
