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

// PTXServiceDTORailSpecificationV3LiteTrainODFareFare Fare
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.LiteTrain.ODFare.Fare
type PTXServiceDTORailSpecificationV3LiteTrainODFareFare struct {

	// Int32
	//
	// 艙等 = ['1: 標準座車廂', '2: 商務座車廂', '3: 自由座車廂']
	// Required: true
	CabinClass *int32 `json:"CabinClass"`

	// Int32
	//
	// 費率等級 = ['1: 成人', '2: 學生', '3: 孩童', '4: 敬老', '5: 愛心', '6: 愛心孩童', '7: 愛心優待/愛心陪伴', '8: 團體', '9: 軍警']
	// Required: true
	FareClass *int32 `json:"FareClass"`

	// Int32
	//
	// 計費價格(新台幣)
	// Required: true
	Price *int32 `json:"Price"`

	// Int32
	//
	// 票種類型 = ['1: 一般票', '2: 來回票', '3: 電子票證', '4: 回數票', '5: 定期票(30天期)', '6: 定期票(60天期)', '7: 早鳥票']
	// Required: true
	TicketType *int32 `json:"TicketType"`
}

// Validate validates this p t x service d t o rail specification v3 lite train o d fare fare
func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCabinClass(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) validateCabinClass(formats strfmt.Registry) error {

	if err := validate.Required("CabinClass", "body", m.CabinClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) validateFareClass(formats strfmt.Registry) error {

	if err := validate.Required("FareClass", "body", m.FareClass); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("Price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) validateTicketType(formats strfmt.Registry) error {

	if err := validate.Required("TicketType", "body", m.TicketType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 lite train o d fare fare based on context it is used
func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainODFareFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3LiteTrainODFareFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
