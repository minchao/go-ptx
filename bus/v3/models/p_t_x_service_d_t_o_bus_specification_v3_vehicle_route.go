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

// PTXServiceDTOBusSpecificationV3VehicleRoute VehicleRoute
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.VehicleRoute
type PTXServiceDTOBusSpecificationV3VehicleRoute struct {

	// String
	//
	// 車牌號碼
	// Required: true
	PlateNumb *string `json:"PlateNumb" xml:"PlateNumb"`

	// String
	//
	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID" xml:"RouteID"`

	// NameType
	//
	// 路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// String
	//
	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID" xml:"RouteUID"`

	// String
	//
	// 地區既用中之附屬路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty" xml:"SubRouteID,omitempty"`

	// NameType
	//
	// 路線名稱
	SubRouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"SubRouteName,omitempty" xml:"NameType,omitempty"`

	// String
	//
	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty" xml:"SubRouteUID,omitempty"`
}

// Validate validates this p t x service d t o bus specification v3 vehicle route
func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlateNumb(formats); err != nil {
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

	if err := m.validateSubRouteName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) validatePlateNumb(formats strfmt.Registry) error {

	if err := validate.Required("PlateNumb", "body", m.PlateNumb); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) validateSubRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 vehicle route based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) contextValidateSubRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3VehicleRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3VehicleRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
