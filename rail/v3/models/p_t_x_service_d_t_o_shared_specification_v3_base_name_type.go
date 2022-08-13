// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOSharedSpecificationV3BaseNameType NameType
//
// swagger:model PTX.Service.DTO.Shared.Specification.V3.Base.NameType
type PTXServiceDTOSharedSpecificationV3BaseNameType struct {

	// String
	//
	// 英文名稱
	En string `json:"En,omitempty" xml:"En,omitempty"`

	// String
	//
	// 中文繁體名稱
	ZhTw string `json:"Zh_tw,omitempty" xml:"Zh_tw,omitempty"`
}

// Validate validates this p t x service d t o shared specification v3 base name type
func (m *PTXServiceDTOSharedSpecificationV3BaseNameType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o shared specification v3 base name type based on context it is used
func (m *PTXServiceDTOSharedSpecificationV3BaseNameType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV3BaseNameType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV3BaseNameType) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOSharedSpecificationV3BaseNameType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
