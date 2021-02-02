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

// PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod DiscountPeriod
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.RouteFare+StageFare+DiscountPeriod
type PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod struct {

	// String
	// Required: true
	EndTime *string `json:"EndTime"`

	// ServiceDay
	// Required: true
	ServiceDay struct {
		PTXServiceDTOBusSpecificationV3RouteFareStageFareServiceDay
	} `json:"ServiceDay"`

	// String
	// Required: true
	StartTime *string `json:"StartTime"`
}

// Validate validates this p t x service d t o bus specification v3 route fare stage fare discount period
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("EndTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) validateServiceDay(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o bus specification v3 route fare stage fare discount period based on the context it is used
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3RouteFareStageFareDiscountPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
