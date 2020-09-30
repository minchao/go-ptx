// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation AlertScopeStation
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRAAlertList.AlertScopeStation
type PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation struct {

	// String
	//
	// 車站代碼
	StationID string `json:"StationID,omitempty"`

	// String
	//
	// 車站名稱
	StationName string `json:"StationName,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope station
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}