// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3VehicleDepot VehicleDepot
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.VehicleDepot
type PTXServiceDTOBusSpecificationV3VehicleDepot struct {

	// String
	//
	// 營業所代碼
	// Required: true
	DepotID *string `json:"DepotID" xml:"DepotID"`

	// NameType
	//
	// 營業所名稱
	// Required: true
	DepotName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"DepotName" xml:"NameType"`

	// Array
	//
	// 車輛
	// Required: true
	Vehicles []*PTXServiceDTOBusSpecificationV3VehicleDepotVehicle "json:\"Vehicles\" xml:\"List`1\""
}

// Validate validates this p t x service d t o bus specification v3 vehicle depot
func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepotName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVehicles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) validateDepotID(formats strfmt.Registry) error {

	if err := validate.Required("DepotID", "body", m.DepotID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) validateDepotName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) validateVehicles(formats strfmt.Registry) error {

	if err := validate.Required("Vehicles", "body", m.Vehicles); err != nil {
		return err
	}

	for i := 0; i < len(m.Vehicles); i++ {
		if swag.IsZero(m.Vehicles[i]) { // not required
			continue
		}

		if m.Vehicles[i] != nil {
			if err := m.Vehicles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vehicles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Vehicles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 vehicle depot based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDepotName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVehicles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) contextValidateDepotName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) contextValidateVehicles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vehicles); i++ {

		if m.Vehicles[i] != nil {
			if err := m.Vehicles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Vehicles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Vehicles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3VehicleDepot) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3VehicleDepot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
