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

// PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews BusTaiwanTripNews
//
// 台灣好行最新消息資料
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.BusTaiwanTripNews
type PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews struct {

	// String
	//
	// 內容描述
	// Required: true
	Description *string `json:"Description" xml:"Description"`

	// 結束時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	// String
	//
	// 語系
	// Required: true
	Language *string `json:"Language" xml:"Language"`

	// String
	//
	// 消息類別
	NewsCategory string `json:"NewsCategory,omitempty" xml:"NewsCategory,omitempty"`

	// String
	//
	// 最新消息原單位發布代碼
	// Required: true
	NewsID *string `json:"NewsID" xml:"NewsID"`

	// 消息公告日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	PublishTime strfmt.DateTime `json:"PublishTime,omitempty"`

	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	SrcUpdateTime strfmt.DateTime `json:"SrcUpdateTime,omitempty"`

	// 開始時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// String
	//
	// 消息標題
	// Required: true
	Title *string `json:"Title" xml:"Title"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o tourism specification v2 bus taiwan trip news
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("Language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateNewsID(formats strfmt.Registry) error {

	if err := validate.Required("NewsID", "body", m.NewsID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validatePublishTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PublishTime", "body", "date-time", m.PublishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateSrcUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcUpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("Title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o tourism specification v2 bus taiwan trip news based on context it is used
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2BusTaiwanTripNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
