// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection AlertScopeLineSection
//
// 受影響的路線區間
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.MRTAlertList.AlertScopeLineSection
type PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection struct {

	// String
	//
	// 影響區間輔助描述
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// String
	//
	// 區間迄站車站代碼
	EndingStationID string `json:"EndingStationID,omitempty" xml:"EndingStationID,omitempty"`

	// String
	//
	// 區間迄站車站名稱
	EndingStationName string `json:"EndingStationName,omitempty" xml:"EndingStationName,omitempty"`

	// String
	//
	// 路線區間所在路線代碼
	LineID string `json:"LineID,omitempty" xml:"LineID,omitempty"`

	// String
	//
	// 區間起站車站代碼
	StartingStationID string `json:"StartingStationID,omitempty" xml:"StartingStationID,omitempty"`

	// String
	//
	// 區間起站車站名稱
	StartingStationName string `json:"StartingStationName,omitempty" xml:"StartingStationName,omitempty"`
}

// Validate validates this p t x service d t o rail specification v2 metro m r t alert list alert scope line section
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 metro m r t alert list alert scope line section based on context it is used
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroMRTAlertListAlertScopeLineSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
