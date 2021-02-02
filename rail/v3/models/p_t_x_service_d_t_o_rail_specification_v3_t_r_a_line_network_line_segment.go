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

// PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment LineSegment
//
// 路線站點間線段資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.LineNetwork.LineSegment
type PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment struct {

	// Single
	//
	// 站間距離
	// Required: true
	Distance *float32 `json:"Distance"`

	// String
	//
	// 線段起點站代碼
	// Required: true
	FromStationID *string `json:"FromStationID"`

	// String
	//
	// 線段代碼
	// Required: true
	LineSegmentID *string `json:"LineSegmentID"`

	// NameType
	//
	// 路段名稱
	// Required: true
	LineSegmentName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"LineSegmentName"`

	// String
	//
	// 線段種類 = ['M: 主路線', 'B: 分支路線']
	// Required: true
	SegmentType *string `json:"SegmentType"`

	// String
	//
	// 線段迄點站代碼
	// Required: true
	ToStationID *string `json:"ToStationID"`
}

// Validate validates this p t x service d t o rail specification v3 t r a line network line segment
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSegmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineSegmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentType(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) validateDistance(formats strfmt.Registry) error {

	if err := validate.Required("Distance", "body", m.Distance); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) validateFromStationID(formats strfmt.Registry) error {

	if err := validate.Required("FromStationID", "body", m.FromStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) validateLineSegmentID(formats strfmt.Registry) error {

	if err := validate.Required("LineSegmentID", "body", m.LineSegmentID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) validateLineSegmentName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) validateSegmentType(formats strfmt.Registry) error {

	if err := validate.Required("SegmentType", "body", m.SegmentType); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) validateToStationID(formats strfmt.Registry) error {

	if err := validate.Required("ToStationID", "body", m.ToStationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a line network line segment based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineSegmentName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) contextValidateLineSegmentName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRALineNetworkLineSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
