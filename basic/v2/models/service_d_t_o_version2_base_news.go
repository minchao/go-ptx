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

// ServiceDTOVersion2BaseNews News
//
// 業管機關資料型別
//
// swagger:model Service.DTO.Version2.Base.News
type ServiceDTOVersion2BaseNews struct {

	// 相關網站連結
	AttachmentURL string `json:"AttachmentURL,omitempty"`

	// 發布單位
	Department string `json:"Department,omitempty"`

	// 內容描述
	// Required: true
	Description *string `json:"Description"`

	// DateTime
	//
	// 結束時間
	EndTime string `json:"EndTime,omitempty"`

	// 語系
	// Required: true
	Language *string `json:"Language"`

	// integer
	//
	// 消息類別 : [1:'最新消息',2:'新聞稿',3:'營運資訊',4:'轉乘資訊',5:'活動訊息',6:'系統公告',7:'新服務上架',8:'API修正',9:'來源異常',99:'其他']
	NewsCategory int32 `json:"NewsCategory,omitempty"`

	// 最新消息原單位發布代碼
	// Required: true
	// Format: uuid
	NewsID *strfmt.UUID `json:"NewsID"`

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

	// 消息標題
	Title string `json:"Title,omitempty"`

	// DateTime
	//
	// 消息更新時間
	UpdateTime string `json:"UpdateTime,omitempty"`
}

// Validate validates this service d t o version2 base news
func (m *ServiceDTOVersion2BaseNews) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion2BaseNews) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BaseNews) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("Language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BaseNews) validateNewsID(formats strfmt.Registry) error {

	if err := validate.Required("NewsID", "body", m.NewsID); err != nil {
		return err
	}

	if err := validate.FormatOf("NewsID", "body", "uuid", m.NewsID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BaseNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BaseNews) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BaseNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
