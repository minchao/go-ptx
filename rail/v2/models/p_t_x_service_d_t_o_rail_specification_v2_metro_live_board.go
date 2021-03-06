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

// PTXServiceDTORailSpecificationV2MetroLiveBoard LiveBoard
//
// 捷運列車到離站動態資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.LiveBoard
type PTXServiceDTORailSpecificationV2MetroLiveBoard struct {

	// String
	//
	// 目的地車站代號
	// Required: true
	DestinationStaionID *string `json:"DestinationStaionID" xml:"String"`

	// String
	//
	// 目的地車站代號
	// Required: true
	DestinationStationID *string `json:"DestinationStationID" xml:"String"`

	// NameType
	//
	// 目的地車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStationName" xml:"NameType"`

	// Int32
	//
	// 下班車次抵達時間預估(分)
	// Required: true
	EstimateTime *int32 `json:"EstimateTime"`

	// String
	//
	// 路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// String
	//
	// 路線編號
	LineNO string `json:"LineNO,omitempty" xml:"String,omitempty"`

	// NameType
	//
	// 路線名稱
	// Required: true
	LineName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"LineName" xml:"NameType"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// String
	//
	// 所在車站代號
	// Required: true
	StationID *string `json:"StationID" xml:"String"`

	// NameType
	//
	// 所在車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName" xml:"NameType"`

	// String
	//
	// 下班車次方向描述
	// Required: true
	TripHeadSign *string `json:"TripHeadSign" xml:"String"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v2 metro live board
func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStaionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineName(formats); err != nil {
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

	if err := m.validateTripHeadSign(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateDestinationStaionID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStaionID", "body", m.DestinationStaionID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateDestinationStationID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationID", "body", m.DestinationStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateEstimateTime(formats strfmt.Registry) error {

	if err := validate.Required("EstimateTime", "body", m.EstimateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateLineName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateTripHeadSign(formats strfmt.Registry) error {

	if err := validate.Required("TripHeadSign", "body", m.TripHeadSign); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro live board based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineName(ctx, formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) contextValidateLineName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroLiveBoard) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroLiveBoard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
