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

// PTXServiceDTOTourismSpecificationV2HotelTourismInfo HotelTourismInfo
//
// 取得觀光旅宿資料
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.HotelTourismInfo
type PTXServiceDTOTourismSpecificationV2HotelTourismInfo struct {

	// String
	//
	// 旅館民宿地址
	Address string `json:"Address,omitempty" xml:"Address,omitempty"`

	// String
	//
	// 所屬縣市
	City string `json:"City,omitempty" xml:"City,omitempty"`

	// String
	//
	// 旅館民宿分類
	Class string `json:"Class,omitempty" xml:"Class,omitempty"`

	// String
	//
	// 旅館民宿簡述
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// String
	//
	// 旅館民宿傳真
	Fax string `json:"Fax,omitempty" xml:"Fax,omitempty"`

	// String
	//
	// 觀光旅館星級
	Grade string `json:"Grade,omitempty" xml:"Grade,omitempty"`

	// String
	//
	// 旅館民宿代碼
	// Required: true
	HotelID *string `json:"HotelID" xml:"HotelID"`

	// String
	//
	// 旅館民宿名稱
	HotelName string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`

	// String
	//
	// 旅館民宿地點簡圖連結網址
	MapURL string `json:"MapUrl,omitempty" xml:"MapUrl,omitempty"`

	// String
	//
	// 停車資訊
	ParkingInfo string `json:"ParkingInfo,omitempty" xml:"ParkingInfo,omitempty"`

	// String
	//
	// 旅館民宿電話
	Phone string `json:"Phone,omitempty" xml:"Phone,omitempty"`

	// TourismPicture
	//
	// 旅館民宿照片
	Picture struct {
		PTXServiceDTOTourismSpecificationV2TourismPicture
	} `json:"Picture,omitempty" xml:"TourismPicture,omitempty"`

	// PointType
	//
	// 旅館民宿位置
	Position struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"Position,omitempty" xml:"PointType,omitempty"`

	// String
	//
	// 服務內容介紹
	ServiceInfo string `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty"`

	// String
	//
	// 房型、價目及數量說明
	Spec string `json:"Spec,omitempty" xml:"Spec,omitempty"`

	// DateTime
	//
	// 觀光局檔案更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// String
	//
	// 旅館民宿網站網址
	WebsiteURL string `json:"WebsiteUrl,omitempty" xml:"WebsiteUrl,omitempty"`

	// String
	//
	// 郵遞區號
	ZipCode string `json:"ZipCode,omitempty" xml:"ZipCode,omitempty"`
}

// Validate validates this p t x service d t o tourism specification v2 hotel tourism info
func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHotelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePicture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) validateHotelID(formats strfmt.Registry) error {

	if err := validate.Required("HotelID", "body", m.HotelID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) validatePicture(formats strfmt.Registry) error {
	if swag.IsZero(m.Picture) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 hotel tourism info based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePicture(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) contextValidatePicture(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2HotelTourismInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2HotelTourismInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
