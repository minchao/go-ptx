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

// PTXServiceDTOShipSpecificationV3ShipDailySchedule ShipDailySchedule
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.ShipDailySchedule
type PTXServiceDTOShipSpecificationV3ShipDailySchedule struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// String
	//
	// 航行日期
	// Required: true
	Date *string `json:"Date" xml:"String"`

	// Int32
	//
	// 方向性描述 : [0:'去程',1:'返程',2:'迴圈']
	Direction int64 `json:"Direction,omitempty"`

	// Array
	//
	// 發車班距資訊
	Frequencies []*PTXServiceDTOShipSpecificationV3Frequency "json:\"Frequencies\" xml:\"List`1\""

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"String"`

	// String
	//
	// 航線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// NameType
	//
	// 航線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// Array
	//
	// 營運定期班表資訊
	Timetables []*PTXServiceDTOShipSpecificationV3Timetable "json:\"Timetables\" xml:\"List`1\""

	// Int32
	//
	// [平臺]資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o ship specification v3 ship daily schedule
func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimetables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("Date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateFrequencies(formats strfmt.Registry) error {
	if swag.IsZero(m.Frequencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Frequencies); i++ {
		if swag.IsZero(m.Frequencies[i]) { // not required
			continue
		}

		if m.Frequencies[i] != nil {
			if err := m.Frequencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Frequencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Frequencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateTimetables(formats strfmt.Registry) error {
	if swag.IsZero(m.Timetables) { // not required
		return nil
	}

	for i := 0; i < len(m.Timetables); i++ {
		if swag.IsZero(m.Timetables[i]) { // not required
			continue
		}

		if m.Timetables[i] != nil {
			if err := m.Timetables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Timetables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Timetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 ship daily schedule based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrequencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimetables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) contextValidateFrequencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Frequencies); i++ {

		if m.Frequencies[i] != nil {
			if err := m.Frequencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Frequencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Frequencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) contextValidateTimetables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Timetables); i++ {

		if m.Timetables[i] != nil {
			if err := m.Timetables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Timetables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Timetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3ShipDailySchedule) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3ShipDailySchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}