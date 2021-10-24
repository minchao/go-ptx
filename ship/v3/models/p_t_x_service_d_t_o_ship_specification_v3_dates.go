// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOShipSpecificationV3Dates Dates
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.Dates
type PTXServiceDTOShipSpecificationV3Dates struct {

	// String
	//
	// 特殊日期
	Date string `json:"Date,omitempty" xml:"String,omitempty"`
}

// Validate validates this p t x service d t o ship specification v3 dates
func (m *PTXServiceDTOShipSpecificationV3Dates) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o ship specification v3 dates based on context it is used
func (m *PTXServiceDTOShipSpecificationV3Dates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Dates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Dates) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3Dates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
