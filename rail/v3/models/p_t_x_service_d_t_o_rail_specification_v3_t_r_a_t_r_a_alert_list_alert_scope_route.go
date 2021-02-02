// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute AlertScopeRoute
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRAAlertList.AlertScopeRoute
type PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute struct {

	// String
	//
	// 營運路線代碼
	RouteID string `json:"RouteID,omitempty"`

	// String
	//
	// 營運路線名稱
	RouteName string `json:"RouteName,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope route
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope route based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
