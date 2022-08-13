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

// PTXServiceDTOBusSpecificationV2ScopeStop Stop
//
// 站牌資料
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.Scope+Stop
type PTXServiceDTOBusSpecificationV2ScopeStop struct {

	// String
	//
	// 站位代碼
	StationID string `json:"StationID,omitempty" xml:"StationID,omitempty"`

	// String
	//
	// 地區既用中之站牌代碼(為原資料內碼)
	StopID string `json:"StopID,omitempty" xml:"StopID,omitempty"`

	// NameType
	//
	// 站牌名稱
	StopName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StopName,omitempty" xml:"NameType,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 scope stop
func (m *PTXServiceDTOBusSpecificationV2ScopeStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ScopeStop) validateStopName(formats strfmt.Registry) error {
	if swag.IsZero(m.StopName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 scope stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2ScopeStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2ScopeStop) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ScopeStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2ScopeStop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2ScopeStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
