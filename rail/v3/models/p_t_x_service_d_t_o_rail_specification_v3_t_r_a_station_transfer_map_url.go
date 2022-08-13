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

// PTXServiceDTORailSpecificationV3TRAStationTransferMapURL MapURL
//
// 轉乘地圖簡圖連結資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.StationTransfer.MapURL
type PTXServiceDTORailSpecificationV3TRAStationTransferMapURL struct {

	// String
	//
	// 樓層
	FloorLevel string `json:"FloorLevel,omitempty" xml:"FloorLevel,omitempty"`

	// String
	//
	// 地圖名稱
	// Required: true
	MapName *string `json:"MapName" xml:"MapName"`

	// String
	//
	// 地圖簡圖連結
	// Required: true
	MapURL *string `json:"MapURL" xml:"MapUrl"`
}

// Validate validates this p t x service d t o rail specification v3 t r a station transfer map URL
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferMapURL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMapName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferMapURL) validateMapName(formats strfmt.Registry) error {

	if err := validate.Required("MapName", "body", m.MapName); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAStationTransferMapURL) validateMapURL(formats strfmt.Registry) error {

	if err := validate.Required("MapURL", "body", m.MapURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a station transfer map URL based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferMapURL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferMapURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAStationTransferMapURL) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAStationTransferMapURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
