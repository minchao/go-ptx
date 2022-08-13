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

// PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable TimeTable
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusDailyTimeTable+TimeTable
type PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable struct {

	// Array
	//
	// 公車停靠時間資料
	// Required: true
	StopTimes []*PTXServiceDTOBusSpecificationV2BusDailyTimeTableStopTime "json:\"StopTimes\" xml:\"List`1\""

	// String
	//
	// 班次代碼
	TripID string `json:"TripID,omitempty" xml:"TripID,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 bus daily time table time table
func (m *PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopTimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable) validateStopTimes(formats strfmt.Registry) error {

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

// ContextValidate validate this p t x service d t o bus specification v2 bus daily time table time table based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable) contextValidateStopTimes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusDailyTimeTableTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
