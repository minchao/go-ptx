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

// ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo RailDailyTrainInfo
//
// 高鐵車次資料型別(時刻表用)
// swagger:model Service.DTO.Version2.Rail.THSR.TimeInfo.RailDailyTrainInfo
type ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo struct {

	// integer
	//
	// 行駛方向 : [0:'南下',1:'北上']
	// Required: true
	Direction *int32 `json:"Direction"`

	// 列車終點車站代號
	EndingStationID string `json:"EndingStationID,omitempty"`

	// NameType
	//
	// 列車終點車站名稱
	EndingStationName *ServiceDTOVersion2BaseNameType `json:"EndingStationName,omitempty"`

	// NameType
	//
	// 附註說明
	Note *ServiceDTOVersion2BaseNameType `json:"Note,omitempty"`

	// 列車起點車站代號
	StartingStationID string `json:"StartingStationID,omitempty"`

	// NameType
	//
	// 列車起點車站名稱
	StartingStationName *ServiceDTOVersion2BaseNameType `json:"StartingStationName,omitempty"`

	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo"`
}

// Validate validates this service d t o version2 rail t h s r time info rail daily train info
func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) validateEndingStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.EndingStationName) { // not required
		return nil
	}

	if m.EndingStationName != nil {
		if err := m.EndingStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndingStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if m.Note != nil {
		if err := m.Note.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Note")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) validateStartingStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.StartingStationName) { // not required
		return nil
	}

	if m.StartingStationName != nil {
		if err := m.StartingStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StartingStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTHSRTimeInfoRailDailyTrainInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
