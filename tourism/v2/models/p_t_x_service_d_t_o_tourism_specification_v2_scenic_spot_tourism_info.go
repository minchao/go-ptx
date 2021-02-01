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

// PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo ScenicSpotTourismInfo
//
// 取得觀光景點資料
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.ScenicSpotTourismInfo
type PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo struct {

	// String
	//
	// 景點地址
	Address string `json:"Address,omitempty"`

	// String
	//
	// 所屬縣市
	City string `json:"City,omitempty"`

	// String
	//
	// 景點分類1
	Class1 string `json:"Class1,omitempty"`

	// String
	//
	// 景點分類2
	Class2 string `json:"Class2,omitempty"`

	// String
	//
	// 景點分類3
	Class3 string `json:"Class3,omitempty"`

	// String
	//
	// 景點特色精簡說明
	Description string `json:"Description,omitempty"`

	// String
	//
	// 景點特色詳細說明
	DescriptionDetail string `json:"DescriptionDetail,omitempty"`

	// String
	//
	// 景點代碼
	// Required: true
	ID *string `json:"ID"`

	// String
	//
	// 常用搜尋關鍵字
	Keyword string `json:"Keyword,omitempty"`

	// String
	//
	// 古蹟分級
	Level string `json:"Level,omitempty"`

	// String
	//
	// 景點地圖/簡圖介紹網址
	MapURL string `json:"MapUrl,omitempty"`

	// String
	//
	// 景點名稱
	Name string `json:"Name,omitempty"`

	// String
	//
	// 開放時間
	OpenTime string `json:"OpenTime,omitempty"`

	// String
	//
	// 停車資訊
	ParkingInfo string `json:"ParkingInfo,omitempty"`

	// PointType
	//
	// 景點主要停車場位置
	ParkingPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"ParkingPosition,omitempty"`

	// String
	//
	// 景點服務電話
	Phone string `json:"Phone,omitempty"`

	// TourismPicture
	//
	// 景點照片
	Picture struct {
		PTXServiceDTOTourismSpecificationV2TourismPicture
	} `json:"Picture,omitempty"`

	// PointType
	//
	// 景點位置
	Position struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"Position,omitempty"`

	// String
	//
	// 警告及注意事項
	Remarks string `json:"Remarks,omitempty"`

	// DateTime
	//
	// 觀光局檔案更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// String
	//
	// 票價資訊
	TicketInfo string `json:"TicketInfo,omitempty"`

	// String
	//
	// 交通資訊
	TravelInfo string `json:"TravelInfo,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// String
	//
	// 景點官方網站網址
	WebsiteURL string `json:"WebsiteUrl,omitempty"`

	// String
	//
	// 郵遞區號
	ZipCode string `json:"ZipCode,omitempty"`
}

// Validate validates this p t x service d t o tourism specification v2 scenic spot tourism info
func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParkingPosition(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) validateParkingPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.ParkingPosition) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) validatePicture(formats strfmt.Registry) error {
	if swag.IsZero(m.Picture) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 scenic spot tourism info based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParkingPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) contextValidateParkingPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) contextValidatePicture(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2ScenicSpotTourismInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
