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

// PTXServiceDTOBusSpecificationV2BusStageFare BusStageFare
//
// 此計費方式以一路線內所有站牌分區收費。(公總稱之為計費站收費, Stage=計費站)
//
// swagger:model PTX.Service.DTO.Bus.Specification.V2.BusStageFare
type PTXServiceDTOBusSpecificationV2BusStageFare struct {

	// BusStage
	//
	// 訖點計費站
	// Required: true
	DestinationStage struct {
		PTXServiceDTOBusSpecificationV2BusStage
	} `json:"DestinationStage" xml:"BusStage"`

	// Int32
	//
	// 方向性描述 : [0:'去程',1:'返程',2:'迴圈',255:'未知']
	// Required: true
	Direction *int64 `json:"Direction"`

	// Array
	//
	// 票價內容
	// Required: true
	Fares []*PTXServiceDTOBusSpecificationV2BusFare "json:\"Fares\" xml:\"List`1\""

	// BusStage
	//
	// 起點計費站
	// Required: true
	OriginStage struct {
		PTXServiceDTOBusSpecificationV2BusStage
	} `json:"OriginStage" xml:"BusStage"`
}

// Validate validates this p t x service d t o bus specification v2 bus stage fare
func (m *PTXServiceDTOBusSpecificationV2BusStageFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) validateDestinationStage(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) validateFares(formats strfmt.Registry) error {

	if err := validate.Required("Fares", "body", m.Fares); err != nil {
		return err
	}

	for i := 0; i < len(m.Fares); i++ {
		if swag.IsZero(m.Fares[i]) { // not required
			continue
		}

		if m.Fares[i] != nil {
			if err := m.Fares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) validateOriginStage(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v2 bus stage fare based on the context it is used
func (m *PTXServiceDTOBusSpecificationV2BusStageFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginStage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) contextValidateDestinationStage(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) contextValidateFares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fares); i++ {

		if m.Fares[i] != nil {
			if err := m.Fares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV2BusStageFare) contextValidateOriginStage(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStageFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV2BusStageFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV2BusStageFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
