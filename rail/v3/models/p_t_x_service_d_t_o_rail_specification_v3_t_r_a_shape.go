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

// PTXServiceDTORailSpecificationV3TRAShape Shape
//
// 臺鐵線型資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.Shape
type PTXServiceDTORailSpecificationV3TRAShape struct {

	// String
	//
	// well-known text，為路線軌跡資料
	// Required: true
	Geometry *string `json:"Geometry" xml:"String"`

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

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v3 t r a shape
func (m *PTXServiceDTORailSpecificationV3TRAShape) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeometry(formats); err != nil {
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

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAShape) validateGeometry(formats strfmt.Registry) error {

	if err := validate.Required("Geometry", "body", m.Geometry); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAShape) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAShape) validateLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAShape) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAShape) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a shape based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRAShape) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAShape) contextValidateLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAShape) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAShape) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAShape
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
