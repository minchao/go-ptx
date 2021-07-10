// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOBusSpecificationV2N1Estimate Estimate
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.N1.Estimate
type PTXServiceDTOBusSpecificationV2N1Estimate struct {

	// 車輛之到站時間預估(秒)
	EstimateTime int32 `json:"EstimateTime,omitempty"`

	// 是否為末班車
	IsLastBus bool `json:"IsLastBus,omitempty"`

	// String
	//
	// 車輛車牌號碼
	PlateNumb string `json:"PlateNumb,omitempty" xml:"String,omitempty"`

	// Int32
	//
	// 車輛於該站之進離站狀態 : [0:'離站',1:'進站']
	VehicleStopStatus int64 `json:"VehicleStopStatus,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 n1 estimate
func (m *PTXServiceDTOBusSpecificationV2N1Estimate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o bus specification v2 n1 estimate based on context it is used
func (m *PTXServiceDTOBusSpecificationV2N1Estimate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2N1Estimate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2N1Estimate) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2N1Estimate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
