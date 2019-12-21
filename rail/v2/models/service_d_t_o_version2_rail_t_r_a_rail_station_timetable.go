// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailTRARailStationTimetable RailStationTimetable
//
// 台鐵車站站別時刻表資料型別
// swagger:model Service.DTO.Version2.Rail.TRA.RailStationTimetable
type ServiceDTOVersion2RailTRARailStationTimetable struct {

	// 到站時間(格式: HH:mm:ss)
	// Required: true
	ArrivalTime *string `json:"ArrivalTime"`

	// 離站時間(格式: HH:mm:ss)
	// Required: true
	DepartureTime *string `json:"DepartureTime"`

	// integer
	//
	// 順逆行 : [0:'順行',1:'逆行']
	// Required: true
	Direction *int32 `json:"Direction"`

	// 終點車站代號
	// Required: true
	EndingStationID *string `json:"EndingStationID"`

	// 終點車站名稱
	// Required: true
	EndingStationName *string `json:"EndingStationName"`

	// 起點車站代號
	// Required: true
	StartingStationID *string `json:"StartingStationID"`

	// 起點車站名稱
	// Required: true
	StartingStationName *string `json:"StartingStationName"`

	// 車站代號
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName"`

	// 時刻表日期(格式: yyyy-MM-dd)
	// Required: true
	TrainDate *string `json:"TrainDate"`

	// 車次代號
	// Required: true
	TrainNo *string `json:"TrainNo"`

	// 列車車種簡碼
	TrainTypeCode string `json:"TrainTypeCode,omitempty"`

	// 列車車種代碼
	TrainTypeID string `json:"TrainTypeID,omitempty"`

	// NameType
	//
	// 列車車種名稱
	TrainTypeName *ServiceDTOVersion2BaseNameType `json:"TrainTypeName,omitempty"`

	// integer
	//
	// 山海線類型 : [0:'不經山海線',1:'山線',2:'海線']
	TripLine int32 `json:"TripLine,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 rail t r a rail station timetable
func (m *ServiceDTOVersion2RailTRARailStationTimetable) Validate(formats strfmt.Registry) error {
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

	if err := m.validateTrainTypeName(formats); err != nil {
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

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalTime", "body", m.ArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateEndingStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationID", "body", m.EndingStationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateEndingStationName(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationName", "body", m.EndingStationName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateStartingStationID(formats strfmt.Registry) error {

	if err := validate.Required("StartingStationID", "body", m.StartingStationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateStartingStationName(formats strfmt.Registry) error {

	if err := validate.Required("StartingStationName", "body", m.StartingStationName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
	}

	if m.StationName != nil {
		if err := m.StationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateTrainDate(formats strfmt.Registry) error {

	if err := validate.Required("TrainDate", "body", m.TrainDate); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateTrainTypeName(formats strfmt.Registry) error {

	if swag.IsZero(m.TrainTypeName) { // not required
		return nil
	}

	if m.TrainTypeName != nil {
		if err := m.TrainTypeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TrainTypeName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailStationTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailStationTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailStationTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTRARailStationTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
