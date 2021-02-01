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

// PTXServiceDTORailSpecificationV2MetroSubClassFare Fare
//
// 票價資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.SubClass.Fare
type PTXServiceDTORailSpecificationV2MetroSubClassFare struct {

	// String
	//
	// 市民所屬縣市代碼 (TPE: 臺北市, NWT: 新北市, OTHERS: 其他)
	CitizenCode string `json:"CitizenCode,omitempty"`

	// Int32
	//
	// 費率等級(1:成人(Adult), 2:學生(Student), 3:孩童(Child), 4:敬老(Senior), 5:愛心(Disabled), 6:愛心孩童(Disabled Child), 7:愛心優待/愛心陪伴, 8:團體(Group))
	// Required: true
	FareClass *int32 `json:"FareClass"`

	// Int32
	//
	// 票價
	// Required: true
	Price *int32 `json:"Price"`

	// String
	//
	// 販售方式 (1: 現場櫃台販售, 2: 現場機器販售, 3: 線上販售, 99: 其他)
	SaleType string `json:"SaleType,omitempty"`

	// Int32
	//
	// 票種(1:一般票(單程票), 2:來回票(悠遊卡/一卡通), 3:電子票証, 4:回數票, 5:定期票(30天期), 6：定期票(60天期))
	// Required: true
	TicketType *int32 `json:"TicketType"`
}

// Validate validates this p t x service d t o rail specification v2 metro sub class fare
func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) validateFareClass(formats strfmt.Registry) error {

	if err := validate.Required("FareClass", "body", m.FareClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("Price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) validateTicketType(formats strfmt.Registry) error {

	if err := validate.Required("TicketType", "body", m.TicketType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 metro sub class fare based on context it is used
func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroSubClassFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
