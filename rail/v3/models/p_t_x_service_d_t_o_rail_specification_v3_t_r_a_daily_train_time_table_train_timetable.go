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

// PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable TrainTimetable
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.DailyTrainTimeTable.TrainTimetable
type PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable struct {

	// Array
	//
	// 停靠站資料
	// Required: true
	StopTimes []*PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableStopTime "json:\"StopTimes\" xml:\"List`1\""

	// TrainInfo
	//
	// 車次資料
	// Required: true
	TrainInfo struct {
		PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainInfo
	} `json:"TrainInfo" xml:"TrainInfo"`
}

// Validate validates this p t x service d t o rail specification v3 t r a daily train time table train timetable
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) validateStopTimes(formats strfmt.Registry) error {

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

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) validateTrainInfo(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a daily train time table train timetable based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrainInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) contextValidateStopTimes(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) contextValidateTrainInfo(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableTrainTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
