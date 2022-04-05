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

// PTXServiceDTORailSpecificationV2TRARailDailyTimetable RailDailyTimetable
//
// 台鐵到離站時刻資料型別
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.RailDailyTimetable
type PTXServiceDTORailSpecificationV2TRARailDailyTimetable struct {

	// RailDailyTrainInfo
	//
	// 車次資料
	// Required: true
	DailyTrainInfo struct {
		PTXServiceDTORailSpecificationV2TRATimeInfoRailDailyTrainInfo
	} `json:"DailyTrainInfo" xml:"RailDailyTrainInfo"`

	// Array
	//
	// 停靠時間資料
	// Required: true
	StopTimes []*PTXServiceDTORailSpecificationV2TRARailStopTime "json:\"StopTimes\" xml:\"List`1\""

	// String
	//
	// 行駛日期(格式: yyyy-MM-dd)
	// Required: true
	TrainDate *string `json:"TrainDate" xml:"String"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 t r a rail daily timetable
func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDailyTrainInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainDate(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) validateDailyTrainInfo(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) validateStopTimes(formats strfmt.Registry) error {

	if err := validate.Required("StopTimes", "body", m.StopTimes); err != nil {
		return err
	}

	for i := 0; i < len(m.StopTimes); i++ {
		if swag.IsZero(m.StopTimes[i]) { // not required
			continue
		}

		if m.StopTimes[i] != nil {
			if err := m.StopTimes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopTimes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StopTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) validateTrainDate(formats strfmt.Registry) error {

	if err := validate.Required("TrainDate", "body", m.TrainDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t r a rail daily timetable based on the context it is used
func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDailyTrainInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) contextValidateDailyTrainInfo(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) contextValidateStopTimes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StopTimes); i++ {

		if m.StopTimes[i] != nil {
			if err := m.StopTimes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopTimes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StopTimes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRARailDailyTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRARailDailyTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
