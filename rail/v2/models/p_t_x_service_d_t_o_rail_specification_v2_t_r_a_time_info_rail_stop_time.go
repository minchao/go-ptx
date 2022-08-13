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

// PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime RailStopTime
//
// 台鐵停靠時間資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.TimeInfo.RailStopTime
type PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime struct {

	// String
	//
	// 到站時間(格式: HH:mm:ss)
	// Required: true
	ArrivalTime *string `json:"ArrivalTime" xml:"ArrivalTime"`

	// String
	//
	// 離站時間(格式: HH:mm:ss)
	// Required: true
	DepartureTime *string `json:"DepartureTime" xml:"DepartureTime"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName" xml:"NameType"`

	// Int32
	//
	// 跑法站序(由1開始)
	// Required: true
	StopSequence *int32 `json:"StopSequence"`

	// Int32
	//
	// 本站是否停駛 : [0:'否',1:'是']
	// Required: true
	SuspendedFlag *int64 `json:"SuspendedFlag"`
}

// Validate validates this p t x service d t o rail specification v2 t r a time info rail stop time
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSuspendedFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) validateStopSequence(formats strfmt.Registry) error {

	if err := validate.Required("StopSequence", "body", m.StopSequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) validateSuspendedFlag(formats strfmt.Registry) error {

	if err := validate.Required("SuspendedFlag", "body", m.SuspendedFlag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t r a time info rail stop time based on the context it is used
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRATimeInfoRailStopTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
