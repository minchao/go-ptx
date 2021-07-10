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

// PTXServiceDTOBikeSpecificationV2BikeAvailability BikeAvailability
//
// 自行車站點剩餘數量資訊
//
// swagger:model PTX.Service.DTO.Bike.Specification.V2.BikeAvailability
type PTXServiceDTOBikeSpecificationV2BikeAvailability struct {

	// Int32
	//
	// 可租借車數
	AvailableRentBikes int32 `json:"AvailableRentBikes,omitempty"`

	// Int32
	//
	// 可歸還車數
	AvailableReturnBikes int32 `json:"AvailableReturnBikes,omitempty"`

	// Int32
	//
	// 服務狀態 : [0:'停止營運',1:'正常營運']
	ServiceAvailable int64 `json:"ServiceAvailable,omitempty"`

	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Format: date-time
	SrcUpdateTime strfmt.DateTime `json:"SrcUpdateTime,omitempty"`

	// String
	//
	// 站點代碼
	StationID string `json:"StationID,omitempty" xml:"String,omitempty"`

	// String
	//
	// 站點唯一識別代碼，規則為 {業管機關代碼} + {StationID}，其中 {業管機關代碼} 可於Authority API中的AuthorityCode欄位查詢
	StationUID string `json:"StationUID,omitempty" xml:"String,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o bike specification v2 bike availability
func (m *PTXServiceDTOBikeSpecificationV2BikeAvailability) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *PTXServiceDTOBikeSpecificationV2BikeAvailability) validateSrcUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcUpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOBikeSpecificationV2BikeAvailability) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o bike specification v2 bike availability based on context it is used
func (m *PTXServiceDTOBikeSpecificationV2BikeAvailability) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBikeSpecificationV2BikeAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBikeSpecificationV2BikeAvailability) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBikeSpecificationV2BikeAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
