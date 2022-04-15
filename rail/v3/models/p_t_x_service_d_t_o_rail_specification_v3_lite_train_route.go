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

// PTXServiceDTORailSpecificationV3LiteTrainRoute Route
//
// 營運路線基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.LiteTrain.Route
type PTXServiceDTORailSpecificationV3LiteTrainRoute struct {

	// String
	//
	// 營運路線迄站代號
	// Required: true
	EndStationID *string `json:"EndStationID" xml:"String"`

	// NameType
	//
	// 營運路線迄站名稱
	// Required: true
	EndStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"EndStationName" xml:"NameType"`

	// String
	//
	// 營運路線所屬之路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// String
	//
	// 營運路線所屬之路線編號
	LineNo string `json:"LineNo,omitempty" xml:"String,omitempty"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorCode *string `json:"OperatorCode" xml:"String"`

	// 路線旅行長度
	RouteDistance float32 `json:"RouteDistance,omitempty"`

	// String
	//
	// 營運路線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// NameType
	//
	// 營運路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// Int32
	//
	// 營運路線種類
	// Required: true
	RouteType *int32 `json:"RouteType"`

	// String
	//
	// 營運路線來源網址
	// Required: true
	RouteURL *string `json:"RouteURL" xml:"String"`

	// String
	//
	// 營運路線起站代號
	// Required: true
	StartStationID *string `json:"StartStationID" xml:"String"`

	// NameType
	//
	// 營運路線起站名稱
	// Required: true
	StartStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StartStationName" xml:"NameType"`
}

// Validate validates this p t x service d t o rail specification v3 lite train route
func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateEndStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndStationID", "body", m.EndStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateEndStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateOperatorCode(formats strfmt.Registry) error {

	if err := validate.Required("OperatorCode", "body", m.OperatorCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateRouteType(formats strfmt.Registry) error {

	if err := validate.Required("RouteType", "body", m.RouteType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateRouteURL(formats strfmt.Registry) error {

	if err := validate.Required("RouteURL", "body", m.RouteURL); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateStartStationID(formats strfmt.Registry) error {

	if err := validate.Required("StartStationID", "body", m.StartStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) validateStartStationName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 lite train route based on the context it is used
func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) contextValidateEndStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) contextValidateStartStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3LiteTrainRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}