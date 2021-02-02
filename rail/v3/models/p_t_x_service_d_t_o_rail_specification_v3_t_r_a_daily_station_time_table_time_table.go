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

// PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable TimeTable
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.DailyStationTimeTable.TimeTable
type PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable struct {

	// String
	//
	// 到站時刻
	ArrivalTime string `json:"ArrivalTime,omitempty"`

	// String
	//
	// 發車時刻
	DepartureTime string `json:"DepartureTime,omitempty"`

	// String
	//
	// 目的站車站代號
	DestinationStationID string `json:"DestinationStationID,omitempty"`

	// NameType
	//
	// 目的站車站名稱
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"DestinationStationName,omitempty"`

	// Int32
	//
	// 發車順序
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// String
	//
	// 車次代碼
	TrainNo string `json:"TrainNo,omitempty"`

	// String
	//
	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	TrainTypeCode string `json:"TrainTypeCode,omitempty"`

	// String
	//
	// 車種代嗎
	TrainTypeID string `json:"TrainTypeID,omitempty"`

	// NameType
	//
	// 車種名稱
	TrainTypeName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"TrainTypeName,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a daily station time table time table
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) Validate(formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) validateDestinationStationName(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationStationName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) validateTrainTypeName(formats strfmt.Registry) error {
	if swag.IsZero(m.TrainTypeName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a daily station time table time table based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrainTypeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) contextValidateTrainTypeName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRADailyStationTimeTableTimeTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
