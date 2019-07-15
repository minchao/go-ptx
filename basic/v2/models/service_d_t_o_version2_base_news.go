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

// ServiceDTOVersion2BaseNews News
//
// 業管機關資料型別
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

	// 消息類別
	// Enum: [1 2 3 4 5 6 7 8 9 99]
	NewsCategory int64 `json:"NewsCategory,omitempty"`

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

	if err := m.validateNewsCategory(formats); err != nil {
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

var serviceDTOVersion2BaseNewsTypeNewsCategoryPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9,99]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BaseNewsTypeNewsCategoryPropEnum = append(serviceDTOVersion2BaseNewsTypeNewsCategoryPropEnum, v)
	}
}

// prop value enum
func (m *ServiceDTOVersion2BaseNews) validateNewsCategoryEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BaseNewsTypeNewsCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BaseNews) validateNewsCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.NewsCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateNewsCategoryEnum("NewsCategory", "body", m.NewsCategory); err != nil {
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
