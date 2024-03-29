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

// PTXServiceDTOBusSpecificationV2TravelTime TravelTime
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.TravelTime
type PTXServiceDTOBusSpecificationV2TravelTime struct {

	// String
	//
	// 起始站位代碼
	// Required: true
	FromStationID *string `json:"FromStationID" xml:"FromStationID"`

	// String
	//
	// 起始站牌代碼
	// Required: true
	FromStopID *string `json:"FromStopID" xml:"FromStopID"`

	// Int32
	//
	// 站間預估行駛時間 [當RunTime值為-1時，代表該區間無提供旅行時間資料。]
	// Required: true
	RunTime *int32 `json:"RunTime"`

	// String
	//
	// 終點站位代碼
	// Required: true
	ToStationID *string `json:"ToStationID" xml:"ToStationID"`

	// String
	//
	// 終點站牌代碼
	// Required: true
	ToStopID *string `json:"ToStopID" xml:"ToStopID"`
}

// Validate validates this p t x service d t o bus specification v2 travel time
func (m *PTXServiceDTOBusSpecificationV2TravelTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStationID(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV2TravelTime) validateFromStationID(formats strfmt.Registry) error {

	if err := validate.Required("FromStationID", "body", m.FromStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2TravelTime) validateFromStopID(formats strfmt.Registry) error {

	if err := validate.Required("FromStopID", "body", m.FromStopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2TravelTime) validateRunTime(formats strfmt.Registry) error {

	if err := validate.Required("RunTime", "body", m.RunTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2TravelTime) validateToStationID(formats strfmt.Registry) error {

	if err := validate.Required("ToStationID", "body", m.ToStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2TravelTime) validateToStopID(formats strfmt.Registry) error {

	if err := validate.Required("ToStopID", "body", m.ToStopID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v2 travel time based on context it is used
func (m *PTXServiceDTOBusSpecificationV2TravelTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2TravelTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2TravelTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2TravelTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
