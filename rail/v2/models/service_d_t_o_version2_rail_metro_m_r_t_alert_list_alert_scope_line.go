// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLine AlertScopeLine
//
// swagger:model Service.DTO.Version2.Rail.Metro.MRTAlertList.AlertScopeLine
type ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLine struct {

	// 實體路線代碼
	LineID string `json:"LineID,omitempty"`

	// 實體路線名稱
	LineName string `json:"LineName,omitempty"`
}

// Validate validates this service d t o version2 rail metro m r t alert list alert scope line
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLine) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLine) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
