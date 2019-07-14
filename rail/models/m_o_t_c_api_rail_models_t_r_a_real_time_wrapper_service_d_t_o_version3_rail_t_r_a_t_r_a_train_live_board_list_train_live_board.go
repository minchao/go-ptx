// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard TRATrainLiveBoardList
// swagger:model MOTC.API.Rail.Models.TRARealTimeWrapper[Service.DTO.Version3.Rail.TRA.TRATrainLiveBoardList.TrainLiveBoard]
type MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard struct {

	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// 來源端平台資料更新週期(秒)
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 資料(陣列)
	// Required: true
	TrainLiveBoards []*ServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard `json:"TrainLiveBoards"`

	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	UpdateTime string `json:"UpdateTime,omitempty"`
}

// Validate validates this m o t c API rail models t r a real time wrapper service d t o version3 rail t r a t r a train live board list train live board
func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainLiveBoards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateTrainLiveBoards(formats strfmt.Registry) error {

	if err := validate.Required("TrainLiveBoards", "body", m.TrainLiveBoards); err != nil {
		return err
	}

	for i := 0; i < len(m.TrainLiveBoards); i++ {
		if swag.IsZero(m.TrainLiveBoards[i]) { // not required
			continue
		}

		if m.TrainLiveBoards[i] != nil {
			if err := m.TrainLiveBoards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TrainLiveBoards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard) UnmarshalBinary(b []byte) error {
	var res MOTCAPIRailModelsTRARealTimeWrapperServiceDTOVersion3RailTRATRATrainLiveBoardListTrainLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
