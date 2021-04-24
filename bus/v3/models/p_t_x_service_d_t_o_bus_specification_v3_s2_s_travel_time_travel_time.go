// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime TravelTime
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.S2STravelTime+TravelTime
type PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime struct {

	// Single
	//
	// 站間距離
	// Required: true
	Distance *float32 `json:"Distance"`

	// String
	//
	// 起站站牌代碼
	// Required: true
	FromStopID *string `json:"FromStopID" xml:"String"`

	// Int32
	//
	// 站間預估行駛時間
	// Required: true
	RunTime *int32 `json:"RunTime"`

	// Int32
	//
	// 站間序號
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// Int32
	//
	// 起站停靠時間
	// Required: true
	StopTime *int32 `json:"StopTime"`

	// String
	//
	// 迄站站牌代碼
	// Required: true
	ToStopID *string `json:"ToStopID" xml:"String"`
}

// Validate validates this p t x service d t o bus specification v3 s2 s travel time travel time
func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("Distance", "body", m.Distance); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) validateFromStopID(formats strfmt.Registry) error {

	if err := validate.Required("FromStopID", "body", m.FromStopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) validateRunTime(formats strfmt.Registry) error {

	if err := validate.Required("RunTime", "body", m.RunTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) validateStopTime(formats strfmt.Registry) error {

	if err := validate.Required("StopTime", "body", m.StopTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) validateToStopID(formats strfmt.Registry) error {

	if err := validate.Required("ToStopID", "body", m.ToStopID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v3 s2 s travel time travel time based on context it is used
func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3S2STravelTimeTravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
