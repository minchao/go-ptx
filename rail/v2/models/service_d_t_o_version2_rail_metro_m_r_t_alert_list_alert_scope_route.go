// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDTOVersion2RailMetroMRTAlertListAlertScopeRoute AlertScopeRoute
//
// swagger:model Service.DTO.Version2.Rail.Metro.MRTAlertList.AlertScopeRoute
type ServiceDTOVersion2RailMetroMRTAlertListAlertScopeRoute struct {

	// 營運路線代碼
	RouteID string `json:"RouteID,omitempty"`

	// 營運路線名稱
	RouteName string `json:"RouteName,omitempty"`
}

// Validate validates this service d t o version2 rail metro m r t alert list alert scope route
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeRoute) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeRoute) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroMRTAlertListAlertScopeRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}