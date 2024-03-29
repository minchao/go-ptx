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

// PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard StationLiveBoard
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRAStationLiveBoardList.StationLiveBoard
type PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard struct {

	// Int32
	//
	// 延誤分鐘
	// Required: true
	DelayTime *int32 `json:"DelayTime"`

	// Int32
	//
	// 行駛方向 : [0:'順行',1:'逆行']
	Direction int64 `json:"Direction,omitempty"`

	// String
	//
	// 終點站代碼
	EndingStationID string `json:"EndingStationID,omitempty" xml:"EndingStationID,omitempty"`

	// NameType
	//
	// 終點站名稱
	// Required: true
	EndingStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"EndingStationName" xml:"NameType"`

	// String
	//
	// 停靠月台(00代表當時尚未確定停靠的月台，待確定好停靠的月台後，就會更新Platfrom。)
	Platform string `json:"Platform,omitempty" xml:"Platform,omitempty"`

	// Int32
	//
	// 列車狀態 : [0:'準點',1:'誤點',2:'取消']
	RunningStatus int64 `json:"RunningStatus,omitempty"`

	// String
	//
	// 表訂到站時刻
	ScheduleArrivalTime string `json:"ScheduleArrivalTime,omitempty" xml:"ScheduleArrivalTime,omitempty"`

	// String
	//
	// 表定離站時刻
	ScheduleDepartureTime string `json:"ScheduleDepartureTime,omitempty" xml:"ScheduleDepartureTime,omitempty"`

	// String
	//
	// 車站代號
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StationName" xml:"NameType"`

	// String
	//
	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo" xml:"TrainNo"`

	// String
	//
	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	TrainTypeCode string `json:"TrainTypeCode,omitempty" xml:"TrainTypeCode,omitempty"`

	// String
	//
	// 車種代嗎
	TrainTypeID string `json:"TrainTypeID,omitempty" xml:"TrainTypeID,omitempty"`

	// NameType
	//
	// 車種名稱
	// Required: true
	TrainTypeName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"TrainTypeName" xml:"NameType"`

	// Int32
	//
	// 山海線類型 : [0:'不經山海線',1:'山線',2:'海線',3:'成追線']
	TripLine int64 `json:"TripLine,omitempty"`

	// DateTime
	//
	// 本筆資料之更新日期時間
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a station live board list station live board
func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelayTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateDelayTime(formats strfmt.Registry) error {

	if err := validate.Required("DelayTime", "body", m.DelayTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateEndingStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateTrainTypeName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a t r a station live board list station live board based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) contextValidateEndingStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) contextValidateTrainTypeName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
