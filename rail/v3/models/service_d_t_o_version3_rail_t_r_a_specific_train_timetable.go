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

// ServiceDTOVersion3RailTRASpecificTrainTimetable SpecificTrainTimetable
//
// 台鐵特殊車次時刻表資料型別
//
// swagger:model Service.DTO.Version3.Rail.TRA.SpecificTrainTimetable
type ServiceDTOVersion3RailTRASpecificTrainTimetable struct {

	// SpecialDay
	//
	// 營運日型態
	// Required: true
	SpecialDay *ServiceDTOVersion3RailTRACommonSpecialDay `json:"SpecialDay"`

	// 停靠時間資料
	// Required: true
	StopTimes []*ServiceDTOVersion3RailTRACommonStopTime `json:"StopTimes"`

	// TrainInfo
	//
	// 定期車次資料
	// Required: true
	TrainInfo *ServiceDTOVersion3RailTRACommonTrainInfo `json:"TrainInfo"`
}

// Validate validates this service d t o version3 rail t r a specific train timetable
func (m *ServiceDTOVersion3RailTRASpecificTrainTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpecialDay(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ServiceDTOVersion3RailTRASpecificTrainTimetable) validateSpecialDay(formats strfmt.Registry) error {

	if err := validate.Required("SpecialDay", "body", m.SpecialDay); err != nil {
		return err
	}

	if m.SpecialDay != nil {
		if err := m.SpecialDay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SpecialDay")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRASpecificTrainTimetable) validateStopTimes(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3RailTRASpecificTrainTimetable) validateTrainInfo(formats strfmt.Registry) error {

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
func (m *ServiceDTOVersion3RailTRASpecificTrainTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRASpecificTrainTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRASpecificTrainTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
