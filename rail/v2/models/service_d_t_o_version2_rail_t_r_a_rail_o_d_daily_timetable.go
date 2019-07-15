// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailTRARailODDailyTimetable RailODDailyTimetable
//
// 台鐵起訖站間到離站時刻資料型別
// swagger:model Service.DTO.Version2.Rail.TRA.RailODDailyTimetable
type ServiceDTOVersion2RailTRARailODDailyTimetable struct {

	// RailDailyTrainInfo
	//
	// 車次資料
	// Required: true
	DailyTrainInfo *ServiceDTOVersion2RailTRATimeInfoRailDailyTrainInfo `json:"DailyTrainInfo"`

	// RailStopTime
	//
	// 迄站停靠時間資料
	// Required: true
	DestinationStopTime *ServiceDTOVersion2RailTRARailStopTime `json:"DestinationStopTime"`

	// RailStopTime
	//
	// 起站停靠時間資料
	// Required: true
	OriginStopTime *ServiceDTOVersion2RailTRARailStopTime `json:"OriginStopTime"`

	// 行駛日期(格式: yyyy:MM:dd)
	// Required: true
	TrainDate *string `json:"TrainDate"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 rail t r a rail o d daily timetable
func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDailyTrainInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStopTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStopTime(formats); err != nil {
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

func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) validateDailyTrainInfo(formats strfmt.Registry) error {

	if err := validate.Required("DailyTrainInfo", "body", m.DailyTrainInfo); err != nil {
		return err
	}

	if m.DailyTrainInfo != nil {
		if err := m.DailyTrainInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DailyTrainInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) validateDestinationStopTime(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStopTime", "body", m.DestinationStopTime); err != nil {
		return err
	}

	if m.DestinationStopTime != nil {
		if err := m.DestinationStopTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DestinationStopTime")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) validateOriginStopTime(formats strfmt.Registry) error {

	if err := validate.Required("OriginStopTime", "body", m.OriginStopTime); err != nil {
		return err
	}

	if m.OriginStopTime != nil {
		if err := m.OriginStopTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginStopTime")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) validateTrainDate(formats strfmt.Registry) error {

	if err := validate.Required("TrainDate", "body", m.TrainDate); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailODDailyTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTRARailODDailyTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
