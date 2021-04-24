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

// PTXServiceDTORailSpecificationV2MetroODFare ODFare
//
// 起迄站間票價資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.ODFare
type PTXServiceDTORailSpecificationV2MetroODFare struct {

	// String
	//
	// 迄站車站代碼
	// Required: true
	DestinationStationID *string `json:"DestinationStationID" xml:"String"`

	// NameType
	//
	// 迄站車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"DestinationStationName" xml:"NameType"`

	// Array
	//
	// 票價資訊
	// Required: true
	Fares []*PTXServiceDTORailSpecificationV2MetroSubClassFare "json:\"Fares\" xml:\"List`1\""

	// String
	//
	// 起站車站代碼
	// Required: true
	OriginStationID *string `json:"OriginStationID" xml:"String"`

	// NameType
	//
	// 起站車站名稱
	// Required: true
	OriginStationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"OriginStationName" xml:"NameType"`

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// 車種(0:不分車種, 1:普通車, 2:直達車)
	TrainType int32 `json:"TrainType,omitempty"`

	// 起迄站間乘車距離(公里)
	TravelDistance float32 `json:"TravelDistance,omitempty"`

	// 起迄站間乘車時間(分)
	TravelTime int32 `json:"TravelTime,omitempty"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// Int32
	//
	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`
}

// Validate validates this p t x service d t o rail specification v2 metro o d fare
func (m *PTXServiceDTORailSpecificationV2MetroODFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateDestinationStationID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationID", "body", m.DestinationStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateFares(formats strfmt.Registry) error {

	if err := validate.Required("Fares", "body", m.Fares); err != nil {
		return err
	}

	for i := 0; i < len(m.Fares); i++ {
		if swag.IsZero(m.Fares[i]) { // not required
			continue
		}

		if m.Fares[i] != nil {
			if err := m.Fares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateOriginStationID(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationID", "body", m.OriginStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateOriginStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro o d fare based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroODFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) contextValidateFares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fares); i++ {

		if m.Fares[i] != nil {
			if err := m.Fares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroODFare) contextValidateOriginStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroODFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroODFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroODFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
