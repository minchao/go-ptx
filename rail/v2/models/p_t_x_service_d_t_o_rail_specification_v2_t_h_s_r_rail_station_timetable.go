// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV2THSRRailStationTimetable RailStationTimetable
//
// 高鐵車站站別時刻表資料型別
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.RailStationTimetable
type PTXServiceDTORailSpecificationV2THSRRailStationTimetable struct {

	// String
	//
	// 到站時間(格式: HH:mm:ss)
	// Required: true
	ArrivalTime *string `json:"ArrivalTime"`

	// String
	//
	// 離站時間(格式: HH:mm:ss)
	// Required: true
	DepartureTime *string `json:"DepartureTime"`

	// integer
	//
	// 順逆行 : [0:'南下',1:'北上']
	// Required: true
	Direction *int32 `json:"Direction"`

	// String
	//
	// 終點車站代號
	// Required: true
	EndingStationID *string `json:"EndingStationID"`

	// NameType
	//
	// 終點車站名稱
	// Required: true
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"EndingStationName"`

	// String
	//
	// 起點車站代號
	// Required: true
	StartingStationID *string `json:"StartingStationID"`

	// NameType
	//
	// 起點車站名稱
	// Required: true
	StartingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StartingStationName"`

	// String
	//
	// 車站代號
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName"`

	// String
	//
	// 時刻表日期(格式: yyyy-MM-dd)
	// Required: true
	TrainDate *string `json:"TrainDate"`

	// String
	//
	// 車次代號
	// Required: true
	TrainNo *string `json:"TrainNo"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 t h s r rail station timetable
func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateEndingStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationID", "body", m.EndingStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateEndingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateStartingStationID(formats strfmt.Registry) error {

	if err := validate.Required("StartingStationID", "body", m.StartingStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateStartingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateTrainDate(formats strfmt.Registry) error {

	if err := validate.Required("TrainDate", "body", m.TrainDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRRailStationTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRRailStationTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}