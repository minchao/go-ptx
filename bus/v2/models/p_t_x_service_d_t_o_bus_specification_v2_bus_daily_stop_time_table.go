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

// PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable BusDailyStopTimeTable
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusDailyStopTimeTable
type PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable struct {

	// DateTime
	//
	// 適用日期
	// Required: true
	// Format: date-time
	BusDate *strfmt.DateTime `json:"BusDate"`

	// String
	//
	// 迄點站站牌代碼,機關定義站牌代號(為原資料內碼)
	// Required: true
	DestinationStopID *string `json:"DestinationStopID" xml:"String"`

	// NameType
	//
	// 迄點站站牌名稱
	DestinationStopName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStopName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 營運業者簡碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"String"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID" xml:"String"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// NameType
	//
	// 路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"String"`

	// Array
	//
	// 公車站牌停靠資料
	// Required: true
	Stops []*PTXServiceDTOBusSpecificationV2BusDailyStopTimeTableStop "json:\"Stops\" xml:\"List`1\""

	// String
	//
	// 地區既用中之附屬路線代碼(為原資料內碼)
	// Required: true
	SubRouteID *string `json:"SubRouteID" xml:"String"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"SubRouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	SubRouteUID *string `json:"SubRouteUID" xml:"String"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bus specification v2 bus daily stop time table
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStopName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
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

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStops(formats); err != nil {
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

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateBusDate(formats strfmt.Registry) error {

	if err := validate.Required("BusDate", "body", m.BusDate); err != nil {
		return err
	}

	if err := validate.FormatOf("BusDate", "body", "date-time", m.BusDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateDestinationStopID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStopID", "body", m.DestinationStopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateDestinationStopName(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationStopName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateStops(formats strfmt.Registry) error {

	if err := validate.Required("Stops", "body", m.Stops); err != nil {
		return err
	}

	for i := 0; i < len(m.Stops); i++ {
		if swag.IsZero(m.Stops[i]) { // not required
			continue
		}

		if m.Stops[i] != nil {
			if err := m.Stops[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stops" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateSubRouteID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteID", "body", m.SubRouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateSubRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateSubRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteUID", "body", m.SubRouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus daily stop time table based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) contextValidateDestinationStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) contextValidateStops(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stops); i++ {

		if m.Stops[i] != nil {
			if err := m.Stops[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stops" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusDailyStopTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}