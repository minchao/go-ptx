// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV2BusRouteFare BusRouteFare
//
// 路線票價資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusRouteFare
type PTXServiceDTOBusSpecificationV2BusRouteFare struct {

	// integer
	//
	// 描述該路線計費方式 : [0:'段次計費',1:'起迄站間計費',2:'計費站區間計費']
	// Required: true
	FarePricingType *int32 `json:"FarePricingType"`

	// integer
	//
	// 該收費方式是否應用到所有附屬路線 : [0:'否',1:'是']
	// Required: true
	IsForAllSubRoutes *int32 `json:"IsForAllSubRoutes"`

	// integer
	//
	// 是否為免費公車 : [0:'否',1:'是']
	// Required: true
	IsFreeBus *int32 `json:"IsFreeBus"`

	// String
	//
	// 營運業者代碼
	// Required: true
	OperatorID *string `json:"OperatorID"`

	// String
	//
	// 機關定義路線代號
	// Required: true
	RouteID *string `json:"RouteID"`

	// String
	//
	// 路線名稱
	RouteName string `json:"RouteName,omitempty"`

	// Array
	//
	// 段次計費
	SectionFares []*PTXServiceDTOBusSpecificationV2SectionFare `json:"SectionFares"`

	// Array
	//
	// 計費站區間計費
	StageFares []*PTXServiceDTOBusSpecificationV2BusStageFare `json:"StageFares"`

	// String
	//
	// 機關定義附屬路線代碼
	SubRouteID string `json:"SubRouteID,omitempty"`

	// String
	//
	// 附屬路線名稱
	SubRouteName string `json:"SubRouteName,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bus specification v2 bus route fare
func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFarePricingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsForAllSubRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsFreeBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSectionFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageFares(formats); err != nil {
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

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateFarePricingType(formats strfmt.Registry) error {

	if err := validate.Required("FarePricingType", "body", m.FarePricingType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateIsForAllSubRoutes(formats strfmt.Registry) error {

	if err := validate.Required("IsForAllSubRoutes", "body", m.IsForAllSubRoutes); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateIsFreeBus(formats strfmt.Registry) error {

	if err := validate.Required("IsFreeBus", "body", m.IsFreeBus); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateOperatorID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorID", "body", m.OperatorID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateSectionFares(formats strfmt.Registry) error {

	if swag.IsZero(m.SectionFares) { // not required
		return nil
	}

	for i := 0; i < len(m.SectionFares); i++ {
		if swag.IsZero(m.SectionFares[i]) { // not required
			continue
		}

		if m.SectionFares[i] != nil {
			if err := m.SectionFares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SectionFares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateStageFares(formats strfmt.Registry) error {

	if swag.IsZero(m.StageFares) { // not required
		return nil
	}

	for i := 0; i < len(m.StageFares); i++ {
		if swag.IsZero(m.StageFares[i]) { // not required
			continue
		}

		if m.StageFares[i] != nil {
			if err := m.StageFares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StageFares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusRouteFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusRouteFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}