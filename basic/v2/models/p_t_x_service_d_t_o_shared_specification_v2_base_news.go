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

// PTXServiceDTOSharedSpecificationV2BaseNews News
//
// 業管機關資料型別
//
// swagger:model PTX.Service.DTO.Shared.Specification.V2.Base.News
type PTXServiceDTOSharedSpecificationV2BaseNews struct {

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
	// 結束時間
	EndTime string `json:"EndTime,omitempty"`

	// String
	//
	// 語系
	// Required: true
	Language *string `json:"Language"`

	// integer
	//
	// 消息類別 : [1:'最新消息',2:'新聞稿',3:'營運資訊',4:'轉乘資訊',5:'活動訊息',6:'系統公告',7:'新服務上架',8:'API修正',9:'來源異常',99:'其他']
	NewsCategory string `json:"NewsCategory,omitempty"`

	// Guid
	//
	// 最新消息原單位發布代碼
	// Required: true
	// Format: uuid
	NewsID *strfmt.UUID `json:"NewsID"`

	// String
	//
	// 報導網站連結
	NewsURL string `json:"NewsURL,omitempty"`

	// DateTime
	//
	// 消息公告日期時間
	PublishTime string `json:"PublishTime,omitempty"`

	// DateTime
	//
	// 開始時間
	StartTime string `json:"StartTime,omitempty"`

	// String
	//
	// 消息標題
	Title string `json:"Title,omitempty"`

	// DateTime
	//
	// 消息更新時間
	UpdateTime string `json:"UpdateTime,omitempty"`
}

// Validate validates this p t x service d t o shared specification v2 base news
func (m *PTXServiceDTOSharedSpecificationV2BaseNews) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewsID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOSharedSpecificationV2BaseNews) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV2BaseNews) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("Language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOSharedSpecificationV2BaseNews) validateNewsID(formats strfmt.Registry) error {

	if err := validate.Required("NewsID", "body", m.NewsID); err != nil {
		return err
	}

	if err := validate.FormatOf("NewsID", "body", "uuid", m.NewsID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o shared specification v2 base news based on context it is used
func (m *PTXServiceDTOSharedSpecificationV2BaseNews) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV2BaseNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOSharedSpecificationV2BaseNews) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOSharedSpecificationV2BaseNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
