// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV2PointType PointType
//
// 座標資料型別
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.PointType
type PTXServiceDTORailSpecificationV2PointType struct {

	// String
	//
	// 地理空間編碼
	GeoHash string `json:"GeoHash,omitempty" xml:"String,omitempty"`

	// 位置緯度(WGS84)
	PositionLat float64 `json:"PositionLat,omitempty"`

	// 位置經度(WGS84)
	PositionLon float64 `json:"PositionLon,omitempty"`
}

// Validate validates this p t x service d t o rail specification v2 point type
func (m *PTXServiceDTORailSpecificationV2PointType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 point type based on context it is used
func (m *PTXServiceDTORailSpecificationV2PointType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2PointType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2PointType) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2PointType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
