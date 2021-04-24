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

// PTXServiceDTORailSpecificationV3TRALine Line
//
// 路線基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.Line
type PTXServiceDTORailSpecificationV3TRALine struct {

	// Boolean
	//
	// 是否位於支線
	// Required: true
	IsBranch *bool `json:"IsBranch"`

	// String
	//
	// 路線顏色
	LineColor string `json:"LineColor,omitempty" xml:"String,omitempty"`

	// String
	//
	// 路線群組
	LineGroup string `json:"LineGroup,omitempty" xml:"String,omitempty"`

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// NameType
	//
	// 路線名稱
	// Required: true
	LineName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"LineName" xml:"NameType"`

	// String
	//
	// 路線編號
	// Required: true
	LineNo *string `json:"LineNo" xml:"String"`

	// NameType
	//
	// 路線區間名稱
	// Required: true
	LineSectionName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"LineSectionName" xml:"NameType"`

	// String
	//
	// 路線群組
	LineURL string `json:"LineURL,omitempty" xml:"String,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a line
func (m *PTXServiceDTORailSpecificationV3TRALine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSectionName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) validateIsBranch(formats strfmt.Registry) error {

	if err := validate.Required("IsBranch", "body", m.IsBranch); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) validateLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) validateLineSectionName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a line based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRALine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineSectionName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) contextValidateLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALine) contextValidateLineSectionName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALine) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRALine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
