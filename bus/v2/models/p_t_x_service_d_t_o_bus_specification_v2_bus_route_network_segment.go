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

// PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment Segment
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusRouteNetwork+Segment
type PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment struct {

	// Single
	//
	// 站間線段距離
	// Required: true
	Distance *float32 `json:"Distance"`

	// String
	//
	// 表BusStation唯一碼(起點站牌)
	// Required: true
	FromStationID *string `json:"FromStationID" xml:"FromStationID"`

	// Single
	//
	// 站間線段序號
	// Required: true
	Sequence *float32 `json:"Sequence"`

	// String
	//
	// 表BusStation唯一碼(迄點站牌)
	// Required: true
	ToStationID *string `json:"ToStationID" xml:"ToStationID"`
}

// Validate validates this p t x service d t o bus specification v2 bus route network segment
func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("Distance", "body", m.Distance); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) validateFromStationID(formats strfmt.Registry) error {

	if err := validate.Required("FromStationID", "body", m.FromStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) validateToStationID(formats strfmt.Registry) error {

	if err := validate.Required("ToStationID", "body", m.ToStationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bus specification v2 bus route network segment based on context it is used
func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusRouteNetworkSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
