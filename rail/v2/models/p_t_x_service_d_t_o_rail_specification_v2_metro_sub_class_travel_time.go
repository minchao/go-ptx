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

// PTXServiceDTORailSpecificationV2MetroSubClassTravelTime TravelTime
//
// 站間運行時間資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.SubClass.TravelTime
type PTXServiceDTORailSpecificationV2MetroSubClassTravelTime struct {

	// String
	//
	// 起站車站代號
	// Required: true
	FromStationID *string `json:"FromStationID" xml:"FromStationID"`

	// NameType
	//
	// 起站車站名稱
	// Required: true
	FromStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"FromStationName" xml:"NameType"`

	// Int32
	//
	// 站間行駛時間(秒)
	// Required: true
	RunTime *int32 `json:"RunTime"`

	// Int32
	//
	// 站間序號
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// 起站停靠時間(秒)
	StopTime int32 `json:"StopTime,omitempty"`

	// String
	//
	// 迄站車站代號
	// Required: true
	ToStationID *string `json:"ToStationID" xml:"ToStationID"`

	// NameType
	//
	// 迄站車站名稱
	// Required: true
	ToStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"ToStationName" xml:"NameType"`
}

// Validate validates this p t x service d t o rail specification v2 metro sub class travel time
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) validateFromStationID(formats strfmt.Registry) error {

	if err := validate.Required("FromStationID", "body", m.FromStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) validateFromStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) validateRunTime(formats strfmt.Registry) error {

	if err := validate.Required("RunTime", "body", m.RunTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) validateToStationID(formats strfmt.Registry) error {

	if err := validate.Required("ToStationID", "body", m.ToStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) validateToStationName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro sub class travel time based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFromStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) contextValidateFromStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) contextValidateToStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroSubClassTravelTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroSubClassTravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
