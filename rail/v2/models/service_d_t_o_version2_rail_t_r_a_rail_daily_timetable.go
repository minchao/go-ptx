// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailTRARailDailyTimetable RailDailyTimetable
//
// 台鐵到離站時刻資料型別
// swagger:model Service.DTO.Version2.Rail.TRA.RailDailyTimetable
type ServiceDTOVersion2RailTRARailDailyTimetable struct {

	// RailDailyTrainInfo
	//
	// 車次資料
	// Required: true
	DailyTrainInfo *ServiceDTOVersion2RailTRATimeInfoRailDailyTrainInfo `json:"DailyTrainInfo"`

	// 停靠時間資料
	// Required: true
	StopTimes []*ServiceDTOVersion2RailTRARailStopTime `json:"StopTimes"`

	// 行駛日期(格式: yyyy-MM-dd)
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

// Validate validates this service d t o version2 rail t r a rail daily timetable
func (m *ServiceDTOVersion2RailTRARailDailyTimetable) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion2RailTRARailDailyTimetable) validateDailyTrainInfo(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2RailTRARailDailyTimetable) validateStopTimes(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTimetable) validateTrainDate(formats strfmt.Registry) error {

	if err := validate.Required("TrainDate", "body", m.TrainDate); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailDailyTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailDailyTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTRARailDailyTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
