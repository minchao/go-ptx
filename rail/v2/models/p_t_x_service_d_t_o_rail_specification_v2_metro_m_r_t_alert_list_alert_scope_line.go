// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLine AlertScopeLine
//
// 受影響的實體路線
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.MRTAlertList.AlertScopeLine
type PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLine struct {

	// String
	//
	// 實體路線代碼
	LineID string `json:"LineID,omitempty"`

	// String
	//
	// 實體路線名稱
	LineName string `json:"LineName,omitempty"`
}

// Validate validates this p t x service d t o rail specification v2 metro m r t alert list alert scope line
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLine) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLine) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
