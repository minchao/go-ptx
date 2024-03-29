// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV2BusStage BusStage
//
// 計費站
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusStage
type PTXServiceDTOBusSpecificationV2BusStage struct {

	// String
	//
	// 站牌代碼
	// Required: true
	StopID *string `json:"StopID" xml:"StopID"`

	// String
	//
	// 站牌名稱
	StopName string `json:"StopName,omitempty" xml:"StopName,omitempty"`
}

// Validate validates this p t x service d t o bus specification v2 bus stage
func (m *PTXServiceDTOBusSpecificationV2BusStage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStage) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v2 bus stage based on context it is used
func (m *PTXServiceDTOBusSpecificationV2BusStage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStage) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusStage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
