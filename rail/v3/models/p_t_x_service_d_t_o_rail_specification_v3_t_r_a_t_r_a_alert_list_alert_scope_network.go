// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork AlertScopeNetwork
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRAAlertList.AlertScopeNetwork
type PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork struct {

	// String
	//
	// 路網代碼
	NetworkID string `json:"NetworkID,omitempty" xml:"NetworkID,omitempty"`

	// String
	//
	// 路網名稱
	NetworkName string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope network
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope network based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
