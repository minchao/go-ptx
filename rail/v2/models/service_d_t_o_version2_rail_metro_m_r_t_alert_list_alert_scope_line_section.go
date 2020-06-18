// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLineSection AlertScopeLineSection
//
// swagger:model Service.DTO.Version2.Rail.Metro.MRTAlertList.AlertScopeLineSection
type ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLineSection struct {

	// 影響區間輔助描述
	Description string `json:"Description,omitempty"`

	// 區間迄站車站代碼
	EndingStationID string `json:"EndingStationID,omitempty"`

	// 區間迄站車站名稱
	EndingStationName string `json:"EndingStationName,omitempty"`

	// 路線區間所在路線代碼
	LineID string `json:"LineID,omitempty"`

	// 區間起站車站代碼
	StartingStationID string `json:"StartingStationID,omitempty"`

	// 區間起站車站名稱
	StartingStationName string `json:"StartingStationName,omitempty"`
}

// Validate validates this service d t o version2 rail metro m r t alert list alert scope line section
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLineSection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLineSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLineSection) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroMRTAlertListAlertScopeLineSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
