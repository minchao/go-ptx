// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOSharedSpecificationV2BaseFare Fare
//
// 票價資料型別
//
// swagger:model PTX.Service.DTO.Shared.Specification.V2.Base.Fare
type PTXServiceDTOSharedSpecificationV2BaseFare struct {

	// 收費價格(新台幣)
	Price float64 `json:"Price,omitempty"`

	// String
	//
	// 票種名稱
	TicketType string `json:"TicketType,omitempty" xml:"TicketType,omitempty"`
}

// Validate validates this p t x service d t o shared specification v2 base fare
func (m *PTXServiceDTOSharedSpecificationV2BaseFare) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o shared specification v2 base fare based on context it is used
func (m *PTXServiceDTOSharedSpecificationV2BaseFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV2BaseFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV2BaseFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOSharedSpecificationV2BaseFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
