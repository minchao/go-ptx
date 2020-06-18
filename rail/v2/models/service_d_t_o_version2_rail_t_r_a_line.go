// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailTRALine Line
//
// 路線基本資料
//
// swagger:model Service.DTO.Version2.Rail.TRA.Line
type ServiceDTOVersion2RailTRALine struct {

	// 是否位於支線
	// Required: true
	IsBranch *bool `json:"IsBranch"`

	// 路線顏色
	LineColor string `json:"LineColor,omitempty"`

	// 路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// 路線英文名稱
	// Required: true
	LineNameEn *string `json:"LineNameEn"`

	// 路線中文名稱
	// Required: true
	LineNameZh *string `json:"LineNameZh"`

	// 路線編號
	// Required: true
	LineNo *string `json:"LineNo"`

	// 路線區間英文名稱
	// Required: true
	LineSectionNameEn *string `json:"LineSectionNameEn"`

	// 路線區間中文名稱
	// Required: true
	LineSectionNameZh *string `json:"LineSectionNameZh"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 rail t r a line
func (m *ServiceDTOVersion2RailTRALine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNameEn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNameZh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSectionNameEn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSectionNameZh(formats); err != nil {
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

func (m *ServiceDTOVersion2RailTRALine) validateIsBranch(formats strfmt.Registry) error {

	if err := validate.Required("IsBranch", "body", m.IsBranch); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateLineNameEn(formats strfmt.Registry) error {

	if err := validate.Required("LineNameEn", "body", m.LineNameEn); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateLineNameZh(formats strfmt.Registry) error {

	if err := validate.Required("LineNameZh", "body", m.LineNameZh); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateLineSectionNameEn(formats strfmt.Registry) error {

	if err := validate.Required("LineSectionNameEn", "body", m.LineSectionNameEn); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateLineSectionNameZh(formats strfmt.Registry) error {

	if err := validate.Required("LineSectionNameZh", "body", m.LineSectionNameZh); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRALine) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRALine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRALine) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTRALine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
