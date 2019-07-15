// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape BusTaiwanTripShape
//
// 台灣好行線型資料
// swagger:model Service.DTO.Version2.TaiwanTripBus.BusTaiwanTripShape
type ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape struct {

	// 去返程，若無值則表示來源尚無區分去返程
	// Required: true
	// Enum: [0 1 2 255]
	Direction *int64 `json:"Direction"`

	// well-known text，為路線軌跡資料
	// Required: true
	Geometry *string `json:"Geometry"`

	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID"`

	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// NameType
	//
	// 台灣好行路線名稱
	TaiwanTripName *ServiceDTOVersion2BaseNameType `json:"TaiwanTripName,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 taiwan trip bus bus taiwan trip shape
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeometry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaiwanTripName(formats); err != nil {
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

var serviceDTOVersion2TaiwanTripBusBusTaiwanTripShapeTypeDirectionPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,255]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2TaiwanTripBusBusTaiwanTripShapeTypeDirectionPropEnum = append(serviceDTOVersion2TaiwanTripBusBusTaiwanTripShapeTypeDirectionPropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateDirectionEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2TaiwanTripBusBusTaiwanTripShapeTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("Direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateGeometry(formats strfmt.Registry) error {

	if err := validate.Required("Geometry", "body", m.Geometry); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateTaiwanTripName(formats strfmt.Registry) error {

	if swag.IsZero(m.TaiwanTripName) { // not required
		return nil
	}

	if m.TaiwanTripName != nil {
		if err := m.TaiwanTripName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaiwanTripName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2TaiwanTripBusBusTaiwanTripShape
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
