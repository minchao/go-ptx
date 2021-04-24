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

// PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong SectionAmong
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.DailyTrainTimeTable.SectionAmong
type PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong struct {

	// String
	//
	// 迄站車站代碼
	// Required: true
	EndStationID *string `json:"EndStationID" xml:"String"`

	// String
	//
	// 起站車站代碼
	// Required: true
	StartStationID *string `json:"StartStationID" xml:"String"`
}

// Validate validates this p t x service d t o rail specification v3 t r a daily train time table section among
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartStationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong) validateEndStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndStationID", "body", m.EndStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong) validateStartStationID(formats strfmt.Registry) error {

	if err := validate.Required("StartStationID", "body", m.StartStationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a daily train time table section among based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRADailyTrainTimeTableSectionAmong
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
