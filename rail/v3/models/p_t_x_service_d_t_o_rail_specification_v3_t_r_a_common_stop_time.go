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

// PTXServiceDTORailSpecificationV3TRACommonStopTime StopTime
//
// 台鐵停靠時間資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.Common.StopTime
type PTXServiceDTORailSpecificationV3TRACommonStopTime struct {

	// String
	//
	// 到站時間(格式: HH:mm)
	// Required: true
	ArrivalTime *string `json:"ArrivalTime" xml:"String"`

	// String
	//
	// 離站時間(格式: HH:mm)
	// Required: true
	DepartureTime *string `json:"DepartureTime" xml:"String"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"String"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StationName" xml:"NameType"`

	// Int32
	//
	// 停靠站序(由1開始)
	// Required: true
	StopSequence *int32 `json:"StopSequence"`
}

// Validate validates this p t x service d t o rail specification v3 t r a common stop time
func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a common stop time based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRACommonStopTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRACommonStopTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
