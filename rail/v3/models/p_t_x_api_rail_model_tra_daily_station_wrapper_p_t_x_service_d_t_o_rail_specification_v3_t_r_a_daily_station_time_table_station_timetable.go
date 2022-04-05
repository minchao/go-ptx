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

// PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable TRADailyStationTimeTableList
//
// swagger:model PTX.API.Rail.Model.TraDailyStationWrapper[PTX.Service.DTO.Rail.Specification.V3.TRA.DailyStationTimeTable.StationTimetable]
type PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable struct {

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
	StationTimetables []*PTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable "json:\"StationTimetables\" xml:\"List`1\""

	// String
	//
	// 營運日說明(yyyy-MM-dd)
	// Required: true
	TrainDate *string `json:"TrainDate" xml:"String"`

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

// Validate validates this p t x API rail model tra daily station wrapper p t x service d t o rail specification v3 t r a daily station time table station timetable
func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStationTimetables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainDate(formats); err != nil {
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

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateStationTimetables(formats strfmt.Registry) error {

	if err := validate.Required("StationTimetables", "body", m.StationTimetables); err != nil {
		return err
	}

	for i := 0; i < len(m.StationTimetables); i++ {
		if swag.IsZero(m.StationTimetables[i]) { // not required
			continue
		}

		if m.StationTimetables[i] != nil {
			if err := m.StationTimetables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationTimetables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StationTimetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateTrainDate(formats strfmt.Registry) error {

	if err := validate.Required("TrainDate", "body", m.TrainDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model tra daily station wrapper p t x service d t o rail specification v3 t r a daily station time table station timetable based on the context it is used
func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationTimetables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) contextValidateStationTimetables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StationTimetables); i++ {

		if m.StationTimetables[i] != nil {
			if err := m.StationTimetables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationTimetables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StationTimetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelTraDailyStationWrapperPTXServiceDTORailSpecificationV3TRADailyStationTimeTableStationTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
