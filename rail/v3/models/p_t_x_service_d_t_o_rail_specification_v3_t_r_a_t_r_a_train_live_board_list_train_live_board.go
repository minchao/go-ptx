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

// PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard TrainLiveBoard
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRATrainLiveBoardList.TrainLiveBoard
type PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard struct {

	// Int32
	//
	// 延誤分鐘
	// Required: true
	DelayTime *int32 `json:"DelayTime"`

	// String
	//
	// 車站代號
	StationID string `json:"StationID,omitempty" xml:"StationID,omitempty"`

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
	TrainNo string `json:"TrainNo,omitempty" xml:"TrainNo,omitempty"`

	// Int32
	//
	// 列車目前所在之車站狀態 : [0:'進站中',1:'在站上',2:'已離站']
	TrainStationStatus int64 `json:"TrainStationStatus,omitempty"`

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

	// DateTime
	//
	// 本筆位置資料之更新日期時間
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a train live board list train live board
func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelayTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) validateDelayTime(formats strfmt.Registry) error {

	if err := validate.Required("DelayTime", "body", m.DelayTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) validateTrainTypeName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a t r a train live board list train live board based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) contextValidateTrainTypeName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRATrainLiveBoardListTrainLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
