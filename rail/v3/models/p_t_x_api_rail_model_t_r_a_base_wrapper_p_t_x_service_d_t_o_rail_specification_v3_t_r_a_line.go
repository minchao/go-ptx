// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine TRALineList
//
// swagger:model PTX.API.Rail.Model.TRABaseWrapper[PTX.Service.DTO.Rail.Specification.V3.TRA.Line]
type PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	Lines []*PTXServiceDTORailSpecificationV3TRALine `json:"Lines"`

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this p t x API rail model t r a base wrapper p t x service d t o rail specification v3 t r a line
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
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

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) validateLines(formats strfmt.Registry) error {

	if err := validate.Required("Lines", "body", m.Lines); err != nil {
		return err
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelTRABaseWrapperPTXServiceDTORailSpecificationV3TRALine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}