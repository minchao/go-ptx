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

// PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine Line
//
// 捷運路網資料-捷運路線資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.NetworkSubClass.Line
type PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine struct {

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// String
	//
	// 路線編號
	LineNo string `json:"LineNo,omitempty"`
}

// Validate validates this p t x service d t o rail specification v2 metro network sub class line
func (m *PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 metro network sub class line based on context it is used
func (m *PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroNetworkSubClassLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
