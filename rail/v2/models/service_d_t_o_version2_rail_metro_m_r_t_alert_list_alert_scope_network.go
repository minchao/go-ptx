// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDTOVersion2RailMetroMRTAlertListAlertScopeNetwork AlertScopeNetwork
//
// swagger:model Service.DTO.Version2.Rail.Metro.MRTAlertList.AlertScopeNetwork
type ServiceDTOVersion2RailMetroMRTAlertListAlertScopeNetwork struct {

	// 路網代碼
	NetworkID string `json:"NetworkID,omitempty"`
}

// Validate validates this service d t o version2 rail metro m r t alert list alert scope network
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeNetwork) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroMRTAlertListAlertScopeNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
