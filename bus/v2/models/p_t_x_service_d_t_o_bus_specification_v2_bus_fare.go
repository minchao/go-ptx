// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV2BusFare BusFare
//
// 票票價種類及費率說明
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusFare
type PTXServiceDTOBusSpecificationV2BusFare struct {

	// Array
	//
	// 優惠時段
	DiscountPeriods []*PTXServiceDTOBusSpecificationV2BusDiscountPeriods `json:"DiscountPeriods"`

	// integer
	//
	// 費率等級 : [1:'成人',2:'學生',3:'孩童',4:'敬老',5:'愛心',6:'愛心孩童',7:'愛心優待或愛心陪伴',8:'團體',9:'軍警',10:'由各運業者自行定義的半票']
	// Required: true
	FareClass *int32 `json:"FareClass"`

	// String
	//
	// 票價名稱
	FareName string `json:"FareName,omitempty"`

	// Int32
	//
	// 計費價格(新台幣)，其中-1表示不提供售票服務
	// Required: true
	Price *int32 `json:"Price"`

	// integer
	//
	// 票種類型 : [1:'一般票',2:'來回票',3:'電子票證',4:'回數票',5:'定期票30天期',6:'定期票60天期',7:'早鳥票',8:'定期票90天期']
	// Required: true
	TicketType *int32 `json:"TicketType"`
}

// Validate validates this p t x service d t o bus specification v2 bus fare
func (m *PTXServiceDTOBusSpecificationV2BusFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountPeriods(formats); err != nil {
		res = append(res, err)
	}

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

func (m *PTXServiceDTOBusSpecificationV2BusFare) validateDiscountPeriods(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountPeriods) { // not required
		return nil
	}

	for i := 0; i < len(m.DiscountPeriods); i++ {
		if swag.IsZero(m.DiscountPeriods[i]) { // not required
			continue
		}

		if m.DiscountPeriods[i] != nil {
			if err := m.DiscountPeriods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DiscountPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusFare) validateFareClass(formats strfmt.Registry) error {

	if err := validate.Required("FareClass", "body", m.FareClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusFare) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("Price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusFare) validateTicketType(formats strfmt.Registry) error {

	if err := validate.Required("TicketType", "body", m.TicketType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}