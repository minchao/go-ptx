// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard TRAStationLiveBoardList
//
// swagger:model PTX.API.Rail.Model.TRARealTimeWrapper[PTX.Service.DTO.Rail.Specification.V3.TRA.TRAStationLiveBoardList.StationLiveBoard]
type PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	StationLiveBoards []*PTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard "json:\"StationLiveBoards\" xml:\"List`1\""

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API rail model t r a real time wrapper p t x service d t o rail specification v3 t r a t r a station live board list station live board
func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStationLiveBoards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
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

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateStationLiveBoards(formats strfmt.Registry) error {

	if err := validate.Required("StationLiveBoards", "body", m.StationLiveBoards); err != nil {
		return err
	}

	for i := 0; i < len(m.StationLiveBoards); i++ {
		if swag.IsZero(m.StationLiveBoards[i]) { // not required
			continue
		}

		if m.StationLiveBoards[i] != nil {
			if err := m.StationLiveBoards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationLiveBoards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StationLiveBoards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model t r a real time wrapper p t x service d t o rail specification v3 t r a t r a station live board list station live board based on the context it is used
func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationLiveBoards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) contextValidateStationLiveBoards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StationLiveBoards); i++ {

		if m.StationLiveBoards[i] != nil {
			if err := m.StationLiveBoards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StationLiveBoards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StationLiveBoards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelTRARealTimeWrapperPTXServiceDTORailSpecificationV3TRATRAStationLiveBoardListStationLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
