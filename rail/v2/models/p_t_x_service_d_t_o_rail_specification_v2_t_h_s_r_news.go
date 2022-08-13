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

// PTXServiceDTORailSpecificationV2THSRNews News
//
// 高鐵最新消息
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.THSR.News
type PTXServiceDTORailSpecificationV2THSRNews struct {

	// Array
	//
	// 消息附件網址連結
	// Required: true
	AttachmentURLList []string "json:\"AttachmentUrlList\" xml:\"List`1\""

	// String
	//
	// 消息內容描述
	// Required: true
	Description *string `json:"Description" xml:"Description"`

	// 結束時間
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
	// Required: true
	NewsCategory *string `json:"NewsCategory" xml:"NewsCategory"`

	// Guid
	//
	// 消息編號
	// Required: true
	// Format: uuid
	NewsID *strfmt.UUID `json:"NewsID"`

	// String
	//
	// 消息網址連結
	// Required: true
	NewsURL *string `json:"NewsUrl" xml:"NewsUrl"`

	// DateTime
	//
	// 消息發布日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	PublishTime *strfmt.DateTime `json:"PublishTime"`

	// DateTime
	//
	// 開始時間
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"StartTime"`

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

// Validate validates this p t x service d t o rail specification v2 t h s r news
func (m *PTXServiceDTORailSpecificationV2THSRNews) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentURLList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
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

	if err := m.validateNewsURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishTime(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateAttachmentURLList(formats strfmt.Registry) error {

	if err := validate.Required("AttachmentUrlList", "body", m.AttachmentURLList); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("Language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateNewsCategory(formats strfmt.Registry) error {

	if err := validate.Required("NewsCategory", "body", m.NewsCategory); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateNewsID(formats strfmt.Registry) error {

	if err := validate.Required("NewsID", "body", m.NewsID); err != nil {
		return err
	}

	if err := validate.FormatOf("NewsID", "body", "uuid", m.NewsID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateNewsURL(formats strfmt.Registry) error {

	if err := validate.Required("NewsUrl", "body", m.NewsURL); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validatePublishTime(formats strfmt.Registry) error {

	if err := validate.Required("PublishTime", "body", m.PublishTime); err != nil {
		return err
	}

	if err := validate.FormatOf("PublishTime", "body", "date-time", m.PublishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("StartTime", "body", m.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("Title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2THSRNews) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 t h s r news based on context it is used
func (m *PTXServiceDTORailSpecificationV2THSRNews) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2THSRNews) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2THSRNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
