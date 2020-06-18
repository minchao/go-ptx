// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOShipSpecificationV3RouteFare RouteFare
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.RouteFare
type PTXServiceDTOShipSpecificationV3RouteFare struct {

	// PortInfomation
	//
	// 迄點港口資訊
	DestinationPort struct {
		PTXServiceDTOShipSpecificationV3PortInfomation
	} `json:"DestinationPort,omitempty"`

	// String
	//
	// 票價代碼
	FareID string `json:"FareID,omitempty"`

	// Array
	//
	// 票價內容
	Fares []*PTXServiceDTOShipSpecificationV3Fare `json:"Fares"`

	// PortInfomation
	//
	// 起始港口資訊
	OriginPort struct {
		PTXServiceDTOShipSpecificationV3PortInfomation
	} `json:"OriginPort,omitempty"`

	// String
	//
	// 航線代碼
	RouteID string `json:"RouteID,omitempty"`

	// NameType
	//
	// 航線名稱
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName,omitempty"`
}

// Validate validates this p t x service d t o ship specification v3 route fare
func (m *PTXServiceDTOShipSpecificationV3RouteFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3RouteFare) validateDestinationPort(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationPort) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3RouteFare) validateFares(formats strfmt.Registry) error {

	if swag.IsZero(m.Fares) { // not required
		return nil
	}

	for i := 0; i < len(m.Fares); i++ {
		if swag.IsZero(m.Fares[i]) { // not required
			continue
		}

		if m.Fares[i] != nil {
			if err := m.Fares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3RouteFare) validateOriginPort(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginPort) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3RouteFare) validateRouteName(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteName) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3RouteFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3RouteFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3RouteFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
