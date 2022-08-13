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

// PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork LiteTrainNetworkList
//
// swagger:model PTX.API.Rail.Model.LiteTrain.LiteTrainBaseWrapper[PTX.Service.DTO.Rail.Specification.V3.LiteTrain.Network.Network]
type PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork struct {

	// String
	//
	// 業管機關簡碼
	// Required: true
	AuthorityCode *string `json:"AuthorityCode" xml:"AuthorityCode"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// Array
	//
	// 資料(陣列)
	// Required: true
	Networks []*PTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork "json:\"Networks\" xml:\"List`1\""

	// Int32
	//
	// 來源端平台資料更新週期(秒)['-1: 不定期更新']
	// Required: true
	SrcUpdateInterval *int32 `json:"SrcUpdateInterval"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	SrcUpdateTime *strfmt.DateTime `json:"SrcUpdateTime"`

	// Int32
	//
	// 本平台資料更新週期(秒)
	// Required: true
	UpdateInterval *int32 `json:"UpdateInterval"`

	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x API rail model lite train lite train base wrapper p t x service d t o rail specification v3 lite train network network
func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
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

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) validateAuthorityCode(formats strfmt.Registry) error {

	if err := validate.Required("AuthorityCode", "body", m.AuthorityCode); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("Networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) validateSrcUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateInterval", "body", m.SrcUpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("SrcUpdateTime", "body", "date-time", m.SrcUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) validateUpdateInterval(formats strfmt.Registry) error {

	if err := validate.Required("UpdateInterval", "body", m.UpdateInterval); err != nil {
		return err
	}

	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x API rail model lite train lite train base wrapper p t x service d t o rail specification v3 lite train network network based on the context it is used
func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {
			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelLiteTrainLiteTrainBaseWrapperPTXServiceDTORailSpecificationV3LiteTrainNetworkNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
