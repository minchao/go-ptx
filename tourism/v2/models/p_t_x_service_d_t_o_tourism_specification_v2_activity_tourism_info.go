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

// PTXServiceDTOTourismSpecificationV2ActivityTourismInfo ActivityTourismInfo
//
// 取得觀光活動資料
//
// swagger:model PTX.Service.DTO.Tourism.Specification.V2.ActivityTourismInfo
type PTXServiceDTOTourismSpecificationV2ActivityTourismInfo struct {

	// String
	//
	// 活動訊息代碼
	// Required: true
	ActivityID *string `json:"ActivityID" xml:"String"`

	// String
	//
	// 活動名稱
	ActivityName string `json:"ActivityName,omitempty" xml:"String,omitempty"`

	// String
	//
	// 主要活動地點地址
	Address string `json:"Address,omitempty" xml:"String,omitempty"`

	// String
	//
	// 費用標示
	Charge string `json:"Charge,omitempty" xml:"String,omitempty"`

	// String
	//
	// 所屬縣市
	City string `json:"City,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動分類1
	Class1 string `json:"Class1,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動分類2
	Class2 string `json:"Class2,omitempty" xml:"String,omitempty"`

	// String
	//
	// 週期性活動執行時間
	Cycle string `json:"Cycle,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動簡述
	Description string `json:"Description,omitempty" xml:"String,omitempty"`

	// 活動結束時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	// String
	//
	// 活動訊息代碼
	// Required: true
	ID *string `json:"ID" xml:"String"`

	// String
	//
	// 主要活動地點名稱
	Location string `json:"Location,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動地圖/簡圖連結網址
	MapURL string `json:"MapUrl,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動名稱
	Name string `json:"Name,omitempty" xml:"String,omitempty"`

	// String
	//
	// 非週期性活動執行時間
	NonCycle string `json:"NonCycle,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動主辦單位
	Organizer string `json:"Organizer,omitempty" xml:"String,omitempty"`

	// String
	//
	// 停車資訊
	ParkingInfo string `json:"ParkingInfo,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動參與對象
	Particpation string `json:"Particpation,omitempty" xml:"String,omitempty"`

	// String
	//
	// 活動聯絡電話
	Phone string `json:"Phone,omitempty" xml:"String,omitempty"`

	// TourismPicture
	//
	// 活動照片
	Picture struct {
		PTXServiceDTOTourismSpecificationV2TourismPicture
	} `json:"Picture,omitempty" xml:"TourismPicture,omitempty"`

	// PointType
	//
	// 活動位置
	Position struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"Position,omitempty" xml:"PointType,omitempty"`

	// String
	//
	// 備註(其他活動相關事項)
	Remarks string `json:"Remarks,omitempty" xml:"String,omitempty"`

	// 觀光局檔案更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// 活動開始時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// String
	//
	// 交通資訊
	TravelInfo string `json:"TravelInfo,omitempty" xml:"String,omitempty"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// String
	//
	// 活動網址
	WebsiteURL string `json:"WebsiteUrl,omitempty" xml:"String,omitempty"`
}

// Validate validates this p t x service d t o tourism specification v2 activity tourism info
func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateStartTime(formats); err != nil {
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

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validateActivityID(formats strfmt.Registry) error {

	if err := validate.Required("ActivityID", "body", m.ActivityID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validatePicture(formats strfmt.Registry) error {
	if swag.IsZero(m.Picture) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o tourism specification v2 activity tourism info based on the context it is used
func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) contextValidatePicture(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOTourismSpecificationV2ActivityTourismInfo) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOTourismSpecificationV2ActivityTourismInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
