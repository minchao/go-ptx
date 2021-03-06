// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV2BusRouteFareFare Fare
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusRouteFare+Fare
type PTXServiceDTOBusSpecificationV2BusRouteFareFare struct {

	// Int32
	//
	// 費率等級 : [1:'成人',2:'學生',3:'孩童',4:'敬老',5:'愛心',6:'愛心孩童',7:'愛心優待或愛心陪伴',8:'團體',9:'軍警',10:'由各運業者自行定義的半票']
	// Required: true
	FareClass *int64 `json:"FareClass"`

	// Int32
	//
	// 計費價格(新台幣)，其中-1表示不提供售票服務
	// Required: true
	Price *int32 `json:"Price"`

	// Int32
	//
	// 票種類型 : [1:'一般票',2:'來回票',3:'電子票證',4:'回數票',5:'定期票30天期',6:'定期票60天期',7:'早鳥票',8:'定期票90天期']
	// Required: true
	TicketType *int64 `json:"TicketType"`
}

// Validate validates this p t x service d t o bus specification v2 bus route fare fare
func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFareClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) validateFareClass(formats strfmt.Registry) error {

	if err := validate.Required("FareClass", "body", m.FareClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("Price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) validateTicketType(formats strfmt.Registry) error {

	if err := validate.Required("TicketType", "body", m.TicketType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v2 bus route fare fare based on context it is used
func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusRouteFareFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusRouteFareFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
