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

// PTXServiceDTOShipSpecificationV3ShipGeneralSchedule ShipGeneralSchedule
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.ShipGeneralSchedule
type PTXServiceDTOShipSpecificationV3ShipGeneralSchedule struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

	// Int32
	//
	// 方向性描述 : [0:'去程',1:'返程',2:'迴圈']
	// Required: true
	Direction *int64 `json:"Direction"`

	// String
	//
	// 有效起始日期
	// Required: true
	EffectiveDate *string `json:"EffectiveDate" xml:"EffectiveDate"`

	// String
	//
	// 有效終止日期
	ExpireDate string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`

	// Array
	//
	// 發車班距資訊
	Frequencies []*PTXServiceDTOShipSpecificationV3GeneralFrequency "json:\"Frequencies\" xml:\"List`1\""

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"OperatorID"`

	// String
	//
	// 航線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// NameType
	//
	// 航線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// DateTime
	//
	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Array
	//
	// 營運定期班表資訊
	Timetables []*PTXServiceDTOShipSpecificationV3GeneralTimetable "json:\"Timetables\" xml:\"List`1\""

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o ship specification v3 ship general schedule
func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
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

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimetables(formats); err != nil {
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

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateEffectiveDate(formats strfmt.Registry) error {

	if err := validate.Required("EffectiveDate", "body", m.EffectiveDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateFrequencies(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateTimetables(formats strfmt.Registry) error {
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

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 ship general schedule based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) contextValidateFrequencies(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) contextValidateTimetables(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3ShipGeneralSchedule) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3ShipGeneralSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
