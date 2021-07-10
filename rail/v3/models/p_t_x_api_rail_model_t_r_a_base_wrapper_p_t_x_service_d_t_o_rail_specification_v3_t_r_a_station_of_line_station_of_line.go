// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine TRAStationOfLineList
//
// swagger:model PTX.API.Rail.Model.TRABaseWrapper[PTX.Service.DTO.Rail.Specification.V3.TRA.StationOfLine.StationOfLine]
type PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	StationOfLines []*PTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine "json:\"StationOfLines\" xml:\"List`1\""

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API rail model t r a base wrapper p t x service d t o rail specification v3 t r a station of line station of line
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationOfLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
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

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateStationOfLines(formats strfmt.Registry) error {

	if err := validate.Required("StationOfLines", "body", m.StationOfLines); err != nil {
		return err
	}

	for i := 0; i < len(m.StationOfLines); i++ {
		if swag.IsZero(m.StationOfLines[i]) { // not required
			continue
		}

		if m.StationOfLines[i] != nil {
			if err := m.StationOfLines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationOfLines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model t r a base wrapper p t x service d t o rail specification v3 t r a station of line station of line based on the context it is used
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationOfLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) contextValidateStationOfLines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StationOfLines); i++ {

		if m.StationOfLines[i] != nil {
			if err := m.StationOfLines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationOfLines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRAStationOfLineStationOfLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
