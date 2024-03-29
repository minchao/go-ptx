// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV3TRANetworkMapNameType MapNameType
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.Network.MapNameType
type PTXServiceDTORailSpecificationV3TRANetworkMapNameType struct {

	// String
	//
	// 路網圖網址(英文版)
	En string `json:"En,omitempty" xml:"En,omitempty"`

	// String
	//
	// 路網圖名稱
	MapName string `json:"MapName,omitempty" xml:"MapName,omitempty"`

	// String
	//
	// 路網圖網址(中文版)
	ZhTw string `json:"Zh_tw,omitempty" xml:"Zh_tw,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a network map name type
func (m *PTXServiceDTORailSpecificationV3TRANetworkMapNameType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a network map name type based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRANetworkMapNameType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRANetworkMapNameType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRANetworkMapNameType) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRANetworkMapNameType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
