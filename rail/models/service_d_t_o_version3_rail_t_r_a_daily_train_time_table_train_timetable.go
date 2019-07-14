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

// ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable TrainTimetable
// swagger:model Service.DTO.Version3.Rail.TRA.DailyTrainTimeTable.TrainTimetable
type ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable struct {

	// 停靠站資料
	// Required: true
	StopTimes []*ServiceDTOVersion3RailTRADailyTrainTimeTableStopTime `json:"StopTimes"`

	// TrainInfo
	//
	// 車次資料
	// Required: true
	TrainInfo *ServiceDTOVersion3RailTRADailyTrainTimeTableTrainInfo `json:"TrainInfo"`
}

// Validate validates this service d t o version3 rail t r a daily train time table train timetable
func (m *ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable) validateStopTimes(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable) validateTrainInfo(formats strfmt.Registry) error {

	if err := validate.Required("TrainInfo", "body", m.TrainInfo); err != nil {
		return err
	}

	if m.TrainInfo != nil {
		if err := m.TrainInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TrainInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRADailyTrainTimeTableTrainTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
