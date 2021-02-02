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

// PTXServiceDTORailSpecificationV2MetroInformationSpot InformationSpot
//
// 詢問處位置資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.InformationSpot
type PTXServiceDTORailSpecificationV2MetroInformationSpot struct {

	// String
	//
	// 位置描述
	// Required: true
	Description *string `json:"Description"`

	// String
	//
	// 樓層
	// Required: true
	FloorLevel *string `json:"FloorLevel"`
}

// Validate validates this p t x service d t o rail specification v2 metro information spot
func (m *PTXServiceDTORailSpecificationV2MetroInformationSpot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloorLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroInformationSpot) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroInformationSpot) validateFloorLevel(formats strfmt.Registry) error {

	if err := validate.Required("FloorLevel", "body", m.FloorLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 metro information spot based on context it is used
func (m *PTXServiceDTORailSpecificationV2MetroInformationSpot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroInformationSpot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroInformationSpot) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroInformationSpot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
