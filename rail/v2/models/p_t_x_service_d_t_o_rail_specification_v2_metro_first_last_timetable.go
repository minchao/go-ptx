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

// PTXServiceDTORailSpecificationV2MetroFirstLastTimetable FirstLastTimetable
//
// 首末班車時刻表資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.FirstLastTimetable
type PTXServiceDTORailSpecificationV2MetroFirstLastTimetable struct {

	// String
	//
	// 目的站車站代號
	// Required: true
	DestinationStaionID *string `json:"DestinationStaionID"`

	// NameType
	//
	// 目的站車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStationName"`

	// String
	//
	// 首班車時刻
	// Required: true
	FirstTrainTime *string `json:"FirstTrainTime"`

	// String
	//
	// 末班車時刻
	// Required: true
	LastTrainTime *string `json:"LastTrainTime"`

	// String
	//
	// 首末班車次之路線代碼
	// Required: true
	LineID *string `json:"LineID"`

	// String
	//
	// 首末班車次之路線代號
	LineNo string `json:"LineNo,omitempty"`

	// ServiceDay
	//
	// 服務日型態
	// Required: true
	ServiceDay struct {
		PTXServiceDTORailSpecificationV2MetroSubClassServiceDay
	} `json:"ServiceDay"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

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

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`

	// String
	//
	// 首末班車次之目的地方向描述
	TripHeadSign string `json:"TripHeadSign,omitempty"`

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

// Validate validates this p t x service d t o rail specification v2 metro first last timetable
func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) Validate(formats strfmt.Registry) error {
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

	if err := m.validateServiceDay(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateDestinationStaionID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStaionID", "body", m.DestinationStaionID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateFirstTrainTime(formats strfmt.Registry) error {

	if err := validate.Required("FirstTrainTime", "body", m.FirstTrainTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateLastTrainTime(formats strfmt.Registry) error {

	if err := validate.Required("LastTrainTime", "body", m.LastTrainTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateServiceDay(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro first last timetable based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceDay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) contextValidateServiceDay(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroFirstLastTimetable) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroFirstLastTimetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
