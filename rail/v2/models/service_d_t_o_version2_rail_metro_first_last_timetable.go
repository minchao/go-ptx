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

// ServiceDTOVersion2RailMetroFirstLastTimetable FirstLastTimetable
//
// 首末班車時刻表資料
//
// swagger:model Service.DTO.Version2.Rail.Metro.FirstLastTimetable
type ServiceDTOVersion2RailMetroFirstLastTimetable struct {

	// 目的站車站代號
	// Required: true
	DestinationStaionID *string `json:"DestinationStaionID"`

	// NameType
	//
	// 目的站車站名稱
	// Required: true
	DestinationStationName *ServiceDTOVersion2BaseNameType `json:"DestinationStationName"`

	// 首班車時刻
	// Required: true
	FirstTrainTime *string `json:"FirstTrainTime"`

	// 末班車時刻
	// Required: true
	LastTrainTime *string `json:"LastTrainTime"`

	// 首末班車次之路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// 首末班車次之路線代號
	LineNo string `json:"LineNo,omitempty"`

	// ServiceDays
	//
	// 服務日型態
	// Required: true
	ServiceDays *ServiceDTOVersion2RailMetroSubClassServiceDays `json:"ServiceDays"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 車站代號
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName"`

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`

	// 首末班車次之目的地方向描述
	TripHeadSign string `json:"TripHeadSign,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this service d t o version2 rail metro first last timetable
func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStaionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstTrainTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTrainTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
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

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateDestinationStaionID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStaionID", "body", m.DestinationStaionID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateDestinationStationName(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationName", "body", m.DestinationStationName); err != nil {
		return err
	}

	if m.DestinationStationName != nil {
		if err := m.DestinationStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DestinationStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateFirstTrainTime(formats strfmt.Registry) error {

	if err := validate.Required("FirstTrainTime", "body", m.FirstTrainTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateLastTrainTime(formats strfmt.Registry) error {

	if err := validate.Required("LastTrainTime", "body", m.LastTrainTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateServiceDays(formats strfmt.Registry) error {

	if err := validate.Required("ServiceDays", "body", m.ServiceDays); err != nil {
		return err
	}

	if m.ServiceDays != nil {
		if err := m.ServiceDays.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServiceDays")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateStationName(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailMetroFirstLastTimetable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailMetroFirstLastTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
