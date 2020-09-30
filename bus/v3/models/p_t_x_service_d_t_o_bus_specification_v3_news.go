// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTOBusSpecificationV3News News
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.News
type PTXServiceDTOBusSpecificationV3News struct {

	// String
	//
	// 相關網站連結
	AttachmentURL string `json:"AttachmentURL,omitempty"`

	// String
	//
	// 發布單位
	Department string `json:"Department,omitempty"`

	// String
	//
	// 內容描述
	// Required: true
	Description *string `json:"Description"`

	// DateTime
	//
	// 結束時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	EndTime string `json:"EndTime,omitempty"`

	// String
	//
	// 語系
	// Required: true
	Language *string `json:"Language"`

	// String
	//
	// 消息類別= ['1: 最新消息', '2: 新聞稿', '3: 營運資訊', '4: 轉乘資訊', '5: 活動訊息', '6: 系統公告', '99: 其他']
	// Required: true
	NewsCategory *string `json:"NewsCategory"`

	// String
	//
	// 最新消息原單位發布代碼
	// Required: true
	NewsID *string `json:"NewsID"`

	// String
	//
	// 報導網站連結
	NewsURL string `json:"NewsURL,omitempty"`

	// DateTime
	//
	// 消息公告日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	PublishTime *string `json:"PublishTime"`

	// DateTime
	//
	// [來源端平臺]此筆資料最後更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	SrcUpdateTime string `json:"SrcUpdateTime,omitempty"`

	// DateTime
	//
	// 開始時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	StartTime string `json:"StartTime,omitempty"`

	// String
	//
	// 消息標題
	// Required: true
	Title *string `json:"Title"`
}

// Validate validates this p t x service d t o bus specification v3 news
func (m *PTXServiceDTOBusSpecificationV3News) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewsCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOBusSpecificationV3News) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3News) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("Language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3News) validateNewsCategory(formats strfmt.Registry) error {

	if err := validate.Required("NewsCategory", "body", m.NewsCategory); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3News) validateNewsID(formats strfmt.Registry) error {

	if err := validate.Required("NewsID", "body", m.NewsID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3News) validatePublishTime(formats strfmt.Registry) error {

	if err := validate.Required("PublishTime", "body", m.PublishTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBusSpecificationV3News) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("Title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3News) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3News) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3News
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}