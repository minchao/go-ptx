// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOBusSpecificationV2ScopeRoute Route
//
// 路線資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.Scope+Route
type PTXServiceDTOBusSpecificationV2ScopeRoute struct {

	// integer
	//
	// 影響方向 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	Direction int32 `json:"Direction,omitempty"`

	// String
	//
	// 地區既用中之班次代碼(為原資料內碼)
	RouteID string `json:"RouteID,omitempty"`

	// NameType
	//
	// 路線名稱
	RouteName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"RouteName,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 scope route
func (m *PTXServiceDTOBusSpecificationV2ScopeRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ScopeRoute) validateRouteName(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 scope route based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2ScopeRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ScopeRoute) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ScopeRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ScopeRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2ScopeRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}