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

// PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable TRAGeneralTrainTimetableList
//
// swagger:model PTX.API.Rail.Model.TRAGeneralTrainWrapper[PTX.Service.DTO.Rail.Specification.V3.TRA.GeneralTrainTimetable]
type PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

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

	// String
	//
	// 來源版號
	SrcVersion string `json:"SrcVersion,omitempty" xml:"SrcVersion,omitempty"`

	// String
	//
	// 定期性站別時刻表名稱
	TimetableName string `json:"TimetableName,omitempty" xml:"TimetableName,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	TrainTimetables []*PTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable "json:\"TrainTimetables\" xml:\"List`1\""

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// String
	//
	// 時刻表適用情形說明
	ValidityDesciption string `json:"ValidityDesciption,omitempty" xml:"ValidityDesciption,omitempty"`
}

// Validate validates this p t x API rail model t r a general train wrapper p t x service d t o rail specification v3 t r a general train timetable
func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTimetables(formats); err != nil {
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

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateEffectiveDate(formats strfmt.Registry) error {

	if err := validate.Required("EffectiveDate", "body", m.EffectiveDate); err != nil {
		return err
	}

	if err := validate.FormatOf("EffectiveDate", "body", "date-time", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateExpireDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ExpireDate", "body", "date-time", m.ExpireDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateTrainTimetables(formats strfmt.Registry) error {

	if err := validate.Required("TrainTimetables", "body", m.TrainTimetables); err != nil {
		return err
	}

	for i := 0; i < len(m.TrainTimetables); i++ {
		if swag.IsZero(m.TrainTimetables[i]) { // not required
			continue
		}

		if m.TrainTimetables[i] != nil {
			if err := m.TrainTimetables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TrainTimetables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TrainTimetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model t r a general train wrapper p t x service d t o rail specification v3 t r a general train timetable based on the context it is used
func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrainTimetables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) contextValidateTrainTimetables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TrainTimetables); i++ {

		if m.TrainTimetables[i] != nil {
			if err := m.TrainTimetables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TrainTimetables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TrainTimetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelTRAGeneralTrainWrapperPTXServiceDTORailSpecificationV3TRAGeneralTrainTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
