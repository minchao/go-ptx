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

// PTXServiceDTOShipSpecificationV3Stops Stops
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.Stops
type PTXServiceDTOShipSpecificationV3Stops struct {

	// String
	//
	// 港口代碼
	// Required: true
	PortID *string `json:"PortID" xml:"PortID"`

	// NameType
	//
	// 港口名稱
	// Required: true
	PortName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"PortName" xml:"NameType"`

	// Int32
	//
	// 路線站序
	// Required: true
	StopSequence *int32 `json:"StopSequence"`
}

// Validate validates this p t x service d t o ship specification v3 stops
func (m *PTXServiceDTOShipSpecificationV3Stops) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopSequence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stops) validatePortID(formats strfmt.Registry) error {

	if err := validate.Required("PortID", "body", m.PortID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stops) validatePortName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stops) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 stops based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3Stops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePortName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3Stops) contextValidatePortName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Stops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3Stops) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3Stops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
