// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOShipSpecificationV3Vessel Vessel
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.Vessel
type PTXServiceDTOShipSpecificationV3Vessel struct {

	// String
	//
	// 船舶代碼
	VesselID string `json:"VesselID,omitempty"`

	// String
	//
	// 船舶名稱
	VesselName string `json:"VesselName,omitempty"`
}

// Validate validates this p t x service d t o ship specification v3 vessel
func (m *PTXServiceDTOShipSpecificationV3Vessel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Vessel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Vessel) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3Vessel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
