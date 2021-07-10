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

// PTXServiceDTOTourismSpecificationV2BusSchedule BusSchedule
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusSchedule
type PTXServiceDTOTourismSpecificationV2BusSchedule struct {

	// Int32
	//
	// 去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// Array
	//
	// 發車班距
	Frequencys []*PTXServiceDTOTourismSpecificationV2BusFrequency "json:\"Frequencys\" xml:\"List`1\""

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關代碼} + {RouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"String"`

	// String
	//
	// 地區既用中之子路線代碼(為原資料內碼)
	// Required: true
	SubRouteID *string `json:"SubRouteID" xml:"String"`

	// NameType
	//
	// 子路線名稱
	// Required: true
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"SubRouteName" xml:"NameType"`

	// String
	//
	// 子路線唯一識別代碼，規則為 {業管機關代碼} + {SubRouteID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	SubRouteUID *string `json:"SubRouteUID" xml:"String"`

	// NameType
	//
	// 台灣好行路線名稱
	// Required: true
	TaiwanTripName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"TaiwanTripName" xml:"NameType"`

	// Array
	//
	// 預定班表
	Timetables []*PTXServiceDTOTourismSpecificationV2BusTimetable "json:\"Timetables\" xml:\"List`1\""

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o tourism specification v2 bus schedule
func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaiwanTripName(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateFrequencys(formats strfmt.Registry) error {
	if swag.IsZero(m.Frequencys) { // not required
		return nil
	}

	for i := 0; i < len(m.Frequencys); i++ {
		if swag.IsZero(m.Frequencys[i]) { // not required
			continue
		}

		if m.Frequencys[i] != nil {
			if err := m.Frequencys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Frequencys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateSubRouteID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteID", "body", m.SubRouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateSubRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateSubRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteUID", "body", m.SubRouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateTaiwanTripName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateTimetables(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 bus schedule based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrequencys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaiwanTripName(ctx, formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) contextValidateFrequencys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Frequencys); i++ {

		if m.Frequencys[i] != nil {
			if err := m.Frequencys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Frequencys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) contextValidateTaiwanTripName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) contextValidateTimetables(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusSchedule) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
