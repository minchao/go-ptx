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

// PTXServiceDTOBusSpecificationV3RouteNetworkSegment Segment
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.RouteNetwork+Segment
type PTXServiceDTOBusSpecificationV3RouteNetworkSegment struct {

	// Single
	//
	// 站間線段距離
	// Required: true
	Distance *float32 `json:"Distance"`

	// String
	//
	// 表BusStop唯一碼(起點站牌)
	// Required: true
	FromStopID *string `json:"FromStopID" xml:"String"`

	// Single
	//
	// 站間線段序號
	// Required: true
	Sequence *float32 `json:"Sequence"`

	// String
	//
	// 表BusStop唯一碼(迄點站牌)
	// Required: true
	ToStopID *string `json:"ToStopID" xml:"String"`
}

// Validate validates this p t x service d t o bus specification v3 route network segment
func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStopID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStopID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("Distance", "body", m.Distance); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) validateFromStopID(formats strfmt.Registry) error {

	if err := validate.Required("FromStopID", "body", m.FromStopID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) validateToStopID(formats strfmt.Registry) error {

	if err := validate.Required("ToStopID", "body", m.ToStopID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v3 route network segment based on context it is used
func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteNetworkSegment) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3RouteNetworkSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
