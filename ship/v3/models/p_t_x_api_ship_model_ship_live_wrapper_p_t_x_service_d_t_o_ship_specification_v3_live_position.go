// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition ShipLivePositionList
//
// swagger:model PTX.API.Ship.Model.ShipLiveWrapper[PTX.Service.DTO.Ship.Specification.V3.LivePosition]
type PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"String"`

	// 資料總筆數<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為提供資料的總筆數  [該欄位由本平台自動產製]"></span>
	Count int64 `json:"Count,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	LivePositions []*PTXServiceDTOShipSpecificationV3LivePosition "json:\"LivePositions\" xml:\"List`1\""

	// Int32
	//
	// [來源端平臺]資料更新週期(秒)['-1: 不定期更新']<span class="emphasis fas fa-pen" rel="為來源Inbound XML中的UpdateInterval [該欄位由本平台自動產製]"></span>
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)<span class="emphasis fas fa-pen" rel="為來源Inbound XML中的UpdateTime [該欄位由本平台自動產製]"></span>
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Int32
	//
	// [平臺]資料更新週期(秒)<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為中心端更新資料的週期"></span>
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)<span class="emphasis fas fa-pen" rel="與來源Inbound XML不同，為中心端處理完後提供資料的更新時間"></span>
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API ship model ship live wrapper p t x service d t o ship specification v3 live position
func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLivePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateInterval(formats); err != nil {
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

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) validateLivePositions(formats strfmt.Registry) error {

	if err := validate.Required("LivePositions", "body", m.LivePositions); err != nil {
		return err
	}

	for i := 0; i < len(m.LivePositions); i++ {
		if swag.IsZero(m.LivePositions[i]) { // not required
			continue
		}

		if m.LivePositions[i] != nil {
			if err := m.LivePositions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LivePositions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API ship model ship live wrapper p t x service d t o ship specification v3 live position based on the context it is used
func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLivePositions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) contextValidateLivePositions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LivePositions); i++ {

		if m.LivePositions[i] != nil {
			if err := m.LivePositions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LivePositions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition) UnmarshalBinary(b []byte) error {
	var res PTXAPIShipModelShipLiveWrapperPTXServiceDTOShipSpecificationV3LivePosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
