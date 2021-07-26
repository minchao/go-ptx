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

// PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews NewsList
//
// swagger:model PTX.API.Basic.Model.V2.BasicWrapper[PTX.Service.DTO.Shared.Specification.V2.Base.News]
type PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews struct {

	// count
	Count int64 `json:"Count,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	Newses []*PTXServiceDTOSharedSpecificationV2BaseNews "json:\"Newses\" xml:\"List`1\""

	// Int32
	//
	// 資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// 更新日期時間
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API basic model v2 basic wrapper p t x service d t o shared specification v2 base news
func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) validateNewses(formats strfmt.Registry) error {

	if err := validate.Required("Newses", "body", m.Newses); err != nil {
		return err
	}

	for i := 0; i < len(m.Newses); i++ {
		if swag.IsZero(m.Newses[i]) { // not required
			continue
		}

		if m.Newses[i] != nil {
			if err := m.Newses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Newses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API basic model v2 basic wrapper p t x service d t o shared specification v2 base news based on the context it is used
func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) contextValidateNewses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Newses); i++ {

		if m.Newses[i] != nil {
			if err := m.Newses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Newses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews) UnmarshalBinary(b []byte) error {
	var res PTXAPIBasicModelV2BasicWrapperPTXServiceDTOSharedSpecificationV2BaseNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
