// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3BusRouteFare RouteFare
// swagger:model Service.DTO.Version3.Bus.RouteFare
type ServiceDTOVersion3BusRouteFare struct {

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

	// 起迄站間計費
	ODFares []*ServiceDTOVersion3BusODFare `json:"ODFares"`

	// 營運業者代碼
	OperatorID string `json:"OperatorID,omitempty"`

	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID"`

	// NameType
	//
	// 路線名稱
	RouteName *ServiceDTOVersion3BaseNameType `json:"RouteName,omitempty"`

	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID"`

	// 段次計費
	SectionFares []*ServiceDTOVersion3BusSectionFare `json:"SectionFares"`

	// 計費站區間計費
	StageFares []*ServiceDTOVersion3BusStageFare `json:"StageFares"`

	// 地區既用中之附屬路線代碼(為原資料內碼)
	SubRouteID string `json:"SubRouteID,omitempty"`

	// NameType
	//
	// 附屬路線名稱
	SubRouteName *ServiceDTOVersion3BaseNameType `json:"SubRouteName,omitempty"`

	// 附屬路線唯一識別代碼，規則為 {業管機關簡碼} + {SubRouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	SubRouteUID string `json:"SubRouteUID,omitempty"`
}

// Validate validates this service d t o version3 bus route fare
func (m *ServiceDTOVersion3BusRouteFare) Validate(formats strfmt.Registry) error {
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

	if err := m.validateODFares(formats); err != nil {
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

	if err := m.validateSectionFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageFares(formats); err != nil {
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

func (m *ServiceDTOVersion3BusRouteFare) validateFarePricingType(formats strfmt.Registry) error {

	if err := validate.Required("FarePricingType", "body", m.FarePricingType); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateIsForAllSubRoutes(formats strfmt.Registry) error {

	if err := validate.Required("IsForAllSubRoutes", "body", m.IsForAllSubRoutes); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateIsFreeBus(formats strfmt.Registry) error {

	if err := validate.Required("IsFreeBus", "body", m.IsFreeBus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateODFares(formats strfmt.Registry) error {

	if swag.IsZero(m.ODFares) { // not required
		return nil
	}

	for i := 0; i < len(m.ODFares); i++ {
		if swag.IsZero(m.ODFares[i]) { // not required
			continue
		}

		if m.ODFares[i] != nil {
			if err := m.ODFares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ODFares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	if m.RouteName != nil {
		if err := m.RouteName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RouteName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusRouteFare) validateSectionFares(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3BusRouteFare) validateStageFares(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion3BusRouteFare) validateSubRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.SubRouteName) { // not required
		return nil
	}

	if m.SubRouteName != nil {
		if err := m.SubRouteName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubRouteName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusRouteFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusRouteFare) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusRouteFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
