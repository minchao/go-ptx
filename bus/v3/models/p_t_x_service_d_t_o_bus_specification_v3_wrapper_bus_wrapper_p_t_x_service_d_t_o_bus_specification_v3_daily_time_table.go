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

// PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable BusDailyTimeTableList
//
// 基本 wrapper
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Wrapper.BusWrapper[PTX.Service.DTO.Bus.Specification.V3.DailyTimeTable]
type PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Array
	//
	// 資料列表
	// Required: true
	DailyTimeTables []*PTXServiceDTOBusSpecificationV3DailyTimeTable "json:\"DailyTimeTables\" xml:\"List`1\""

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
}

// Validate validates this p t x service d t o bus specification v3 wrapper bus wrapper p t x service d t o bus specification v3 daily time table
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDailyTimeTables(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) validateDailyTimeTables(formats strfmt.Registry) error {

	if err := validate.Required("DailyTimeTables", "body", m.DailyTimeTables); err != nil {
		return err
	}

	for i := 0; i < len(m.DailyTimeTables); i++ {
		if swag.IsZero(m.DailyTimeTables[i]) { // not required
			continue
		}

		if m.DailyTimeTables[i] != nil {
			if err := m.DailyTimeTables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DailyTimeTables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DailyTimeTables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 wrapper bus wrapper p t x service d t o bus specification v3 daily time table based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDailyTimeTables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) contextValidateDailyTimeTables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DailyTimeTables); i++ {

		if m.DailyTimeTables[i] != nil {
			if err := m.DailyTimeTables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DailyTimeTables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DailyTimeTables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3WrapperBusWrapperPTXServiceDTOBusSpecificationV3DailyTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
