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

// PTXServiceDTOBusSpecificationV3SubRouteStop Stop
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.SubRoute+Stop
type PTXServiceDTOBusSpecificationV3SubRouteStop struct {

	// String
	//
	// 站牌代碼
	// Required: true
	StopID *string `json:"StopID" xml:"StopID"`

	// NameType
	//
	// 站牌名稱
	// Required: true
	StopName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StopName" xml:"NameType"`
}

// Validate validates this p t x service d t o bus specification v3 sub route stop
func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) validateStopID(formats strfmt.Registry) error {

	if err := validate.Required("StopID", "body", m.StopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) validateStopName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 sub route stop based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) contextValidateStopName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3SubRouteStop) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3SubRouteStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
