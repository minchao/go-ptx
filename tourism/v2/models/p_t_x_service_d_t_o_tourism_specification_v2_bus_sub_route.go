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

// PTXServiceDTOTourismSpecificationV2BusSubRoute BusSubRoute
//
// 附屬路線資料型別
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusSubRoute
type PTXServiceDTOTourismSpecificationV2BusSubRoute struct {

	// Int32
	//
	// 去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// String
	//
	// 平日第一班發車時間
	FirstBusTime string `json:"FirstBusTime,omitempty" xml:"FirstBusTime,omitempty"`

	// String
	//
	// 車頭描述
	Headsign string `json:"Headsign,omitempty" xml:"Headsign,omitempty"`

	// String
	//
	// 車頭英文描述
	HeadsignEn string `json:"HeadsignEn,omitempty" xml:"HeadsignEn,omitempty"`

	// String
	//
	// 假日去程第一班發車時間
	HolidayFirstBusTime string `json:"HolidayFirstBusTime,omitempty" xml:"HolidayFirstBusTime,omitempty"`

	// String
	//
	// 假日返程第一班發車時間
	HolidayLastBusTime string `json:"HolidayLastBusTime,omitempty" xml:"HolidayLastBusTime,omitempty"`

	// String
	//
	// 平日返程第一班發車時間
	LastBusTime string `json:"LastBusTime,omitempty" xml:"LastBusTime,omitempty"`

	// Array
	//
	// 營運業者代碼
	// Required: true
	OperatorIDs []string "json:\"OperatorIDs\" xml:\"List`1\""

	// String
	//
	// 地區既用中之附屬路線代碼(為原資料內碼)
	// Required: true
	SubRouteID *string `json:"SubRouteID" xml:"SubRouteID"`

	// NameType
	//
	// 附屬路線名稱
	// Required: true
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"SubRouteName" xml:"NameType"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	SubRouteUID *string `json:"SubRouteUID" xml:"SubRouteUID"`
}

// Validate validates this p t x service d t o tourism specification v2 bus sub route
func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorIDs(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) validateOperatorIDs(formats strfmt.Registry) error {

	if err := validate.Required("OperatorIDs", "body", m.OperatorIDs); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) validateSubRouteID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteID", "body", m.SubRouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) validateSubRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) validateSubRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("SubRouteUID", "body", m.SubRouteUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 bus sub route based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusSubRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusSubRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
