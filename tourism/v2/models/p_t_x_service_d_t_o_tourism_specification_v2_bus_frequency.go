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

// PTXServiceDTOTourismSpecificationV2BusFrequency BusFrequency
//
// 發車班距資料
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusFrequency
type PTXServiceDTOTourismSpecificationV2BusFrequency struct {

	// String
	//
	// 發車班距結束適用時間，格式為: HH:mm
	// Required: true
	EndTime *string `json:"EndTime" xml:"EndTime"`

	// Int32
	//
	// 最大班距時間(分鐘)
	// Required: true
	MaxHeadwayMins *int32 `json:"MaxHeadwayMins"`

	// Int32
	//
	// 最小班距時間(分鐘)
	// Required: true
	MinHeadwayMins *int32 `json:"MinHeadwayMins"`

	// ServiceDay
	//
	// 週內營運日
	ServiceDay struct {
		PTXServiceDTOTourismSpecificationV2ServiceDay
	} `json:"ServiceDay,omitempty" xml:"ServiceDay,omitempty"`

	// Array
	//
	// 特殊營運日
	SpecialDays []*PTXServiceDTOTourismSpecificationV2SpecialDay "json:\"SpecialDays\" xml:\"List`1\""

	// String
	//
	// 發車班距起始適用時間，格式為: HH:mm
	// Required: true
	StartTime *string `json:"StartTime" xml:"StartTime"`
}

// Validate validates this p t x service d t o tourism specification v2 bus frequency
func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxHeadwayMins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinHeadwayMins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecialDays(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("EndTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) validateMaxHeadwayMins(formats strfmt.Registry) error {

	if err := validate.Required("MaxHeadwayMins", "body", m.MaxHeadwayMins); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) validateMinHeadwayMins(formats strfmt.Registry) error {

	if err := validate.Required("MinHeadwayMins", "body", m.MinHeadwayMins); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) validateServiceDay(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceDay) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) validateSpecialDays(formats strfmt.Registry) error {
	if swag.IsZero(m.SpecialDays) { // not required
		return nil
	}

	for i := 0; i < len(m.SpecialDays); i++ {
		if swag.IsZero(m.SpecialDays[i]) { // not required
			continue
		}

		if m.SpecialDays[i] != nil {
			if err := m.SpecialDays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 bus frequency based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpecialDays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) contextValidateSpecialDays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SpecialDays); i++ {

		if m.SpecialDays[i] != nil {
			if err := m.SpecialDays[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SpecialDays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusFrequency) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusFrequency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
