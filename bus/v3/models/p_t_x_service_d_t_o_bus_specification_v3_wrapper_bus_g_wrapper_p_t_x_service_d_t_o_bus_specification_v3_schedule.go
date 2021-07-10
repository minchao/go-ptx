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

// PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule BusScheduleList
//
// 定期型 wrapper (有生效的起迄時間資訊)
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Wrapper.BusGWrapper[PTX.Service.DTO.Bus.Specification.V3.Schedule]
type PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// DateTime
	//
	// 有效起始日期
	// Required: true
	// Format: date-time
	EffectiveDate *strfmt.DateTime `json:"EffectiveDate"`

	// 有效終止日期
	// Format: date-time
	ExpireDate strfmt.DateTime `json:"ExpireDate,omitempty"`

	// String
	//
	// 公車定期營運班表名稱
	ScheduleName string `json:"ScheduleName,omitempty" xml:"String,omitempty"`

	// Array
	//
	// 資料列表
	// Required: true
	Schedules []*PTXServiceDTOBusSpecificationV3Schedule "json:\"Schedules\" xml:\"List`1\""

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Int32
	//
	// [平臺]資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// String
	//
	// 時刻表適用情形說明
	ValidityDesciption string `json:"ValidityDesciption,omitempty" xml:"String,omitempty"`
}

// Validate validates this p t x service d t o bus specification v3 wrapper bus g wrapper p t x service d t o bus specification v3 schedule
func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedules(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateEffectiveDate(formats strfmt.Registry) error {

	if err := validate.Required("EffectiveDate", "body", m.EffectiveDate); err != nil {
		return err
	}

	if err := validate.FormatOf("EffectiveDate", "body", "date-time", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateExpireDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ExpireDate", "body", "date-time", m.ExpireDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateSchedules(formats strfmt.Registry) error {

	if err := validate.Required("Schedules", "body", m.Schedules); err != nil {
		return err
	}

	for i := 0; i < len(m.Schedules); i++ {
		if swag.IsZero(m.Schedules[i]) { // not required
			continue
		}

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 wrapper bus g wrapper p t x service d t o bus specification v3 schedule based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) contextValidateSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Schedules); i++ {

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3WrapperBusGWrapperPTXServiceDTOBusSpecificationV3Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
