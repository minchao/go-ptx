// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PTXServiceDTOShipSpecificationV3Operators Operators
// swagger:model PTX.Service.DTO.Ship.Specification.V3.Operators
type PTXServiceDTOShipSpecificationV3Operators struct {

	// String
	//
	// 營運業者代碼
	OperatorID string `json:"OperatorID,omitempty"`
}

// Validate validates this p t x service d t o ship specification v3 operators
func (m *PTXServiceDTOShipSpecificationV3Operators) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Operators) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Operators) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3Operators
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
