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

// PTXServiceDTORailSpecificationV2TRALineSimple Line_Simple
//
// 路線基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.Line_Simple
type PTXServiceDTORailSpecificationV2TRALineSimple struct {

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// String
	//
	// 路線編號
	// Required: true
	LineNo *string `json:"LineNo" xml:"String"`
}

// Validate validates this p t x service d t o rail specification v2 t r a line simple
func (m *PTXServiceDTORailSpecificationV2TRALineSimple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRALineSimple) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRALineSimple) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 t r a line simple based on context it is used
func (m *PTXServiceDTORailSpecificationV2TRALineSimple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRALineSimple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRALineSimple) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRALineSimple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
