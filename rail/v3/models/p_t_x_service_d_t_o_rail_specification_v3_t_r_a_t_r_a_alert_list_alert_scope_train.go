// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain AlertScopeTrain
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRAAlertList.AlertScopeTrain
type PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain struct {

	// String
	//
	// 受影響的車次
	TrainNo string `json:"TrainNo,omitempty" xml:"TrainNo,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope train
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope train based on context it is used
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
