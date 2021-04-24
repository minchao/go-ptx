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

// PTXServiceDTORailSpecificationV2MetroStationTimeTable StationTimeTable
//
// 站別時刻表資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.StationTimeTable
type PTXServiceDTORailSpecificationV2MetroStationTimeTable struct {

	// String
	//
	// 目的站車站代號
	// Required: true
	DestinationStaionID *string `json:"DestinationStaionID" xml:"String"`

	// NameType
	//
	// 目的站車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStationName" xml:"NameType"`

	// integer
	//
	// 營運路線方向描述 : [0:'去程',1:'返程']
	Direction string `json:"Direction,omitempty"`

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// String
	//
	// 營運路線代碼
	RouteID string `json:"RouteID,omitempty" xml:"String,omitempty"`

	// ServiceDay
	//
	// 服務日型態
	// Required: true
	ServiceDay struct {
		PTXServiceDTORailSpecificationV2MetroSubClassServiceDay
	} `json:"ServiceDay" xml:"ServiceDay"`

	// Array
	//
	// 特定日期
	SpecialDays []*PTXServiceDTORailSpecificationV2MetroSubClassSpecialDay "json:\"SpecialDays\" xml:\"List`1\""

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"String"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName" xml:"NameType"`

	// Array
	//
	// 車站發車時刻資訊
	// Required: true
	Timetables []*PTXServiceDTORailSpecificationV2MetroSubClassTimetable "json:\"Timetables\" xml:\"List`1\""

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 metro station time table
func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStaionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecialDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimetables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateDestinationStaionID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStaionID", "body", m.DestinationStaionID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateServiceDay(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateSpecialDays(formats strfmt.Registry) error {
	if swag.IsZero(m.SpecialDays) { // not required
		return nil
	}

	for i := 0; i < len(m.SpecialDays); i++ {
		if swag.IsZero(m.SpecialDays[i]) { // not required
			continue
		}

		if m.SpecialDays[i] != nil {
			if err := m.SpecialDays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateTimetables(formats strfmt.Registry) error {

	if err := validate.Required("Timetables", "body", m.Timetables); err != nil {
		return err
	}

	for i := 0; i < len(m.Timetables); i++ {
		if swag.IsZero(m.Timetables[i]) { // not required
			continue
		}

		if m.Timetables[i] != nil {
			if err := m.Timetables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Timetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro station time table based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpecialDays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationName(ctx, formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) contextValidateSpecialDays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SpecialDays); i++ {

		if m.SpecialDays[i] != nil {
			if err := m.SpecialDays[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) contextValidateTimetables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Timetables); i++ {

		if m.Timetables[i] != nil {
			if err := m.Timetables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Timetables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroStationTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroStationTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
