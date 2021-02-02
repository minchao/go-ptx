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

// PTXServiceDTORailSpecificationV2MetroLine Line
//
// 捷運路線資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.Line
type PTXServiceDTORailSpecificationV2MetroLine struct {

	// Boolean
	//
	// 是否位於支線
	// Required: true
	IsBranch *bool `json:"IsBranch"`

	// String
	//
	// 路線顏色
	// Required: true
	LineColor *string `json:"LineColor"`

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// NameType
	//
	// 路線名稱
	// Required: true
	LineName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"LineName"`

	// String
	//
	// 路線編號
	LineNo string `json:"LineNo,omitempty"`

	// NameType
	//
	// 路線區間名稱
	// Required: true
	LineSectionName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"LineSectionName"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 metro line
func (m *PTXServiceDTORailSpecificationV2MetroLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSectionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateIsBranch(formats strfmt.Registry) error {

	if err := validate.Required("IsBranch", "body", m.IsBranch); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateLineColor(formats strfmt.Registry) error {

	if err := validate.Required("LineColor", "body", m.LineColor); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateLineSectionName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro line based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroLine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV2MetroLine) contextValidateLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLine) contextValidateLineSectionName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroLine) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
