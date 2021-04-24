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

// PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay SpecialDay
//
// 特定日期
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.SubClass.SpecialDay
type PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay struct {

	// String
	//
	// 描述
	// Required: true
	Description *string `json:"Description" xml:"String"`

	// DateTime
	//
	// 結束時間
	// Required: true
	EndDate *string `json:"EndDate"`

	// DateTime
	//
	// 開始時間
	// Required: true
	SaterDate *string `json:"SaterDate"`
}

// Validate validates this p t x service d t o rail specification v2 metro sub class special day
func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaterDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("EndDate", "body", m.EndDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) validateSaterDate(formats strfmt.Registry) error {

	if err := validate.Required("SaterDate", "body", m.SaterDate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 metro sub class special day based on context it is used
func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
