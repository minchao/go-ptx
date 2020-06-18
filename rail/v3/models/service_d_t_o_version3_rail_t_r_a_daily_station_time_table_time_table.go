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

// ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable TimeTable
//
// swagger:model Service.DTO.Version3.Rail.TRA.DailyStationTimeTable.TimeTable
type ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable struct {

	// 到站時刻
	ArrivalTime string `json:"ArrivalTime,omitempty"`

	// 發車時刻
	DepartureTime string `json:"DepartureTime,omitempty"`

	// 目的站車站代號
	DestinationStationID string `json:"DestinationStationID,omitempty"`

	// NameType
	//
	// 目的站車站名稱
	DestinationStationName *ServiceDTOVersion3BaseNameType `json:"DestinationStationName,omitempty"`

	// 發車順序
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// 車次代碼
	TrainNo string `json:"TrainNo,omitempty"`

	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	TrainTypeCode string `json:"TrainTypeCode,omitempty"`

	// 車種代嗎
	TrainTypeID string `json:"TrainTypeID,omitempty"`

	// NameType
	//
	// 車種名稱
	TrainTypeName *ServiceDTOVersion3BaseNameType `json:"TrainTypeName,omitempty"`
}

// Validate validates this service d t o version3 rail t r a daily station time table time table
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTypeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable) validateDestinationStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationStationName) { // not required
		return nil
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

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable) validateTrainTypeName(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRADailyStationTimeTableTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
