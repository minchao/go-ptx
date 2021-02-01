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

// PTXServiceDTORailSpecificationV2TRARailLiveBoard RailLiveBoard
//
// 台鐵車次動態車次站別即時電子看板資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.RailLiveBoard
type PTXServiceDTORailSpecificationV2TRARailLiveBoard struct {

	// Int32
	//
	// 誤點時間(0:準點;>=1誤點)
	// Required: true
	DelayTime *int32 `json:"DelayTime"`

	// integer
	//
	// 順逆行 : [0:'順行',1:'逆行']
	// Required: true
	Direction *int32 `json:"Direction"`

	// String
	//
	// 車次終點車站代號
	// Required: true
	EndingStationID *string `json:"EndingStationID"`

	// NameType
	//
	// 車次終點車站名稱
	// Required: true
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"EndingStationName"`

	// String
	//
	// 表訂到站時間(格式: HH:mm:ss)
	// Required: true
	ScheduledArrivalTime *string `json:"ScheduledArrivalTime"`

	// String
	//
	// 表訂離站時間(格式: HH:mm:ss)
	// Required: true
	ScheduledDepartureTime *string `json:"ScheduledDepartureTime"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// String
	//
	// 車站代碼
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
	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo"`

	// String
	//
	// 列車車種簡碼
	TrainTypeCode string `json:"TrainTypeCode,omitempty"`

	// String
	//
	// 列車車種代碼
	TrainTypeID string `json:"TrainTypeID,omitempty"`

	// NameType
	//
	// 列車車種名稱
	TrainTypeName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"TrainTypeName,omitempty"`

	// integer
	//
	// 山海線類型 : [0:'不經山海線',1:'山線',2:'海線',3:'成追線']
	TripLine int32 `json:"TripLine,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v2 t r a rail live board
func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelayTime(formats); err != nil {
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

	if err := m.validateScheduledArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledDepartureTime(formats); err != nil {
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

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTypeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateDelayTime(formats strfmt.Registry) error {

	if err := validate.Required("DelayTime", "body", m.DelayTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateEndingStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationID", "body", m.EndingStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateEndingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateScheduledArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("ScheduledArrivalTime", "body", m.ScheduledArrivalTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateScheduledDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("ScheduledDepartureTime", "body", m.ScheduledDepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateTrainTypeName(formats strfmt.Registry) error {
	if swag.IsZero(m.TrainTypeName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 t r a rail live board based on the context it is used
func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndingStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationName(ctx, formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) contextValidateEndingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) contextValidateTrainTypeName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRARailLiveBoard) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRARailLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
