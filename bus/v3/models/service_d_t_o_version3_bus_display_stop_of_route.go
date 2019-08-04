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

// ServiceDTOVersion3BusDisplayStopOfRoute DisplayStopOfRoute
// swagger:model Service.DTO.Version3.Bus.DisplayStopOfRoute
type ServiceDTOVersion3BusDisplayStopOfRoute struct {

	// integer
	//
	// 車輛去返程 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int32 `json:"Direction"`

	// 地區既用中之路線代碼(為原資料內碼)
	// Required: true
	RouteID *string `json:"RouteID"`

	// NameType
	//
	// 路線名稱
	// Required: true
	RouteName *ServiceDTOVersion3BaseNameType `json:"RouteName"`

	// 路線唯一識別代碼，規則為 {業管機關簡碼} + {RouteID}，其中 {業管機關簡碼} 可於Authority API中的AuthorityCode欄位查詢
	// Required: true
	RouteUID *string `json:"RouteUID"`

	// 路線站序資料
	// Required: true
	Stops []*ServiceDTOVersion3BusDisplayStopOfRouteStop `json:"Stops"`
}

// Validate validates this service d t o version3 bus display stop of route
func (m *ServiceDTOVersion3BusDisplayStopOfRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3BusDisplayStopOfRoute) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusDisplayStopOfRoute) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusDisplayStopOfRoute) validateRouteName(formats strfmt.Registry) error {

	if err := validate.Required("RouteName", "body", m.RouteName); err != nil {
		return err
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

func (m *ServiceDTOVersion3BusDisplayStopOfRoute) validateRouteUID(formats strfmt.Registry) error {

	if err := validate.Required("RouteUID", "body", m.RouteUID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3BusDisplayStopOfRoute) validateStops(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3BusDisplayStopOfRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3BusDisplayStopOfRoute) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3BusDisplayStopOfRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}