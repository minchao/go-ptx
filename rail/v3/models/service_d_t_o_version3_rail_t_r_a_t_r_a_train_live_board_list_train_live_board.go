// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard TrainLiveBoard
// swagger:model Service.DTO.Version3.Rail.TRA.TRATrainLiveBoardList.TrainLiveBoard
type ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard struct {

	// 延誤分鐘
	// Required: true
	DelayTime *int32 `json:"DelayTime"`

	// 車站代號
	StationID string `json:"StationID,omitempty"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName *ServiceDTOVersion3BaseNameType `json:"StationName"`

	// 車次代碼
	TrainNo string `json:"TrainNo,omitempty"`

	// 列車目前所在之車站狀態
	// Enum: [0 1 2]
	TrainStationStatus int64 `json:"TrainStationStatus,omitempty"`

	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	TrainTypeCode string `json:"TrainTypeCode,omitempty"`

	// 車種代嗎
	TrainTypeID string `json:"TrainTypeID,omitempty"`

	// NameType
	//
	// 車種名稱
	// Required: true
	TrainTypeName *ServiceDTOVersion3BaseNameType `json:"TrainTypeName"`

	// DateTime
	//
	// 本筆位置資料之更新日期時間
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version3 rail t r a t r a train live board list train live board
func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelayTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainStationStatus(formats); err != nil {
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

func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateDelayTime(formats strfmt.Registry) error {

	if err := validate.Required("DelayTime", "body", m.DelayTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateStationName(formats strfmt.Registry) error {

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

var serviceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoardTypeTrainStationStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoardTypeTrainStationStatusPropEnum = append(serviceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoardTypeTrainStationStatusPropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateTrainStationStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoardTypeTrainStationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateTrainStationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TrainStationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrainStationStatusEnum("TrainStationStatus", "body", m.TrainStationStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateTrainTypeName(formats strfmt.Registry) error {

	if err := validate.Required("TrainTypeName", "body", m.TrainTypeName); err != nil {
		return err
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

func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
