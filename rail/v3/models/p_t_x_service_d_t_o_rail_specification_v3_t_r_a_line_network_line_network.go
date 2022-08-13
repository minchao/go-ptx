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

// PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork LineNetwork
//
// 路線網路拓撲基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.LineNetwork.LineNetwork
type PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork struct {

	// String
	//
	// 路線編號
	// Required: true
	LineID *string `json:"LineID" xml:"LineID"`

	// NameType
	//
	// 路線名稱
	// Required: true
	LineName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"LineName" xml:"NameType"`

	// Array
	//
	// 路線站點間線段資訊
	LineSegments []*PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment "json:\"LineSegments\" xml:\"List`1\""
}

// Validate validates this p t x service d t o rail specification v3 t r a line network line network
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSegments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) validateLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) validateLineSegments(formats strfmt.Registry) error {
	if swag.IsZero(m.LineSegments) { // not required
		return nil
	}

	for i := 0; i < len(m.LineSegments); i++ {
		if swag.IsZero(m.LineSegments[i]) { // not required
			continue
		}

		if m.LineSegments[i] != nil {
			if err := m.LineSegments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LineSegments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("LineSegments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a line network line network based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) contextValidateLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) contextValidateLineSegments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LineSegments); i++ {

		if m.LineSegments[i] != nil {
			if err := m.LineSegments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LineSegments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("LineSegments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRALineNetworkLineNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
