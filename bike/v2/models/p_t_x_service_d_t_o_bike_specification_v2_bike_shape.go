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

// PTXServiceDTOBikeSpecificationV2BikeShape BikeShape
//
// swagger:model PTX.Service.DTO.Bike.Specification.V2.BikeShape
type PTXServiceDTOBikeSpecificationV2BikeShape struct {

	// String
	//
	// 業管機關名稱（可能包含多個業管機關）
	AuthorityName string `json:"AuthorityName,omitempty" xml:"String,omitempty"`

	// String
	//
	// 路線所在縣市名稱
	// Required: true
	City *string `json:"City" xml:"String"`

	// String
	//
	// 路線所在縣市代碼
	// Required: true
	CityCode *string `json:"CityCode" xml:"String"`

	// 自行車道長度
	CyclingLength float64 `json:"CyclingLength,omitempty"`

	// String
	//
	// 自行車道類型
	CyclingType string `json:"CyclingType,omitempty" xml:"String,omitempty"`

	// String
	//
	// 自行車道車行方向
	Direction string `json:"Direction,omitempty" xml:"String,omitempty"`

	// String
	//
	// 路線軌跡編碼(encoded polyline)
	// Required: true
	EncodedPolyline *string `json:"EncodedPolyline" xml:"String"`

	// String
	//
	// 自行車道完工日期時間
	FinishedTime string `json:"FinishedTime,omitempty" xml:"String,omitempty"`

	// String
	//
	// well-known text，為路線軌跡資料
	// Required: true
	Geometry *string `json:"Geometry" xml:"String"`

	// String
	//
	// 路線迄點描述
	RoadSectionEnd string `json:"RoadSectionEnd,omitempty" xml:"String,omitempty"`

	// String
	//
	// 路線起點描述
	RoadSectionStart string `json:"RoadSectionStart,omitempty" xml:"String,omitempty"`

	// String
	//
	// 路線名稱
	// Required: true
	RouteName *string `json:"RouteName" xml:"String"`

	// String
	//
	// 路線所在鄉鎮名稱（可能包含多個鄉鎮）
	Town string `json:"Town,omitempty" xml:"String,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bike specification v2 bike shape
func (m *PTXServiceDTOBikeSpecificationV2BikeShape) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncodedPolyline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeometry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
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

func (m *PTXServiceDTOBikeSpecificationV2BikeShape) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("City", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeShape) validateCityCode(formats strfmt.Registry) error {

	if err := validate.Required("CityCode", "body", m.CityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeShape) validateEncodedPolyline(formats strfmt.Registry) error {

	if err := validate.Required("EncodedPolyline", "body", m.EncodedPolyline); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeShape) validateGeometry(formats strfmt.Registry) error {

	if err := validate.Required("Geometry", "body", m.Geometry); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeShape) validateRouteName(formats strfmt.Registry) error {

	if err := validate.Required("RouteName", "body", m.RouteName); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeShape) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bike specification v2 bike shape based on context it is used
func (m *PTXServiceDTOBikeSpecificationV2BikeShape) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBikeSpecificationV2BikeShape) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBikeSpecificationV2BikeShape) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBikeSpecificationV2BikeShape
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
