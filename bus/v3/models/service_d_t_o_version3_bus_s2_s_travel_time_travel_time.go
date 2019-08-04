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

// ServiceDTOVersion3BusS2STravelTimeTravelTime TravelTime
// swagger:model Service.DTO.Version3.Bus.S2STravelTime.TravelTime
type ServiceDTOVersion3BusS2STravelTimeTravelTime struct {

	// 站間距離
	// Required: true
	Distance *float32 `json:"Distance"`

	// 起站站牌代碼
	// Required: true
	FromStopID *string `json:"FromStopID"`

	// 站間預估行駛時間
	// Required: true
	RunTime *int32 `json:"RunTime"`

	// 站間序號
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// 起站停靠時間
	// Required: true
	StopTime *int32 `json:"StopTime"`

	// 迄站站牌代碼
	// Required: true
	ToStopID *string `json:"ToStopID"`
}

// Validate validates this service d t o version3 bus s2 s travel time travel time
func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStopID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("Distance", "body", m.Distance); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) validateFromStopID(formats strfmt.Registry) error {

	if err := validate.Required("FromStopID", "body", m.FromStopID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) validateRunTime(formats strfmt.Registry) error {

	if err := validate.Required("RunTime", "body", m.RunTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) validateStopTime(formats strfmt.Registry) error {

	if err := validate.Required("StopTime", "body", m.StopTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) validateToStopID(formats strfmt.Registry) error {

	if err := validate.Required("ToStopID", "body", m.ToStopID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusS2STravelTimeTravelTime) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusS2STravelTimeTravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
