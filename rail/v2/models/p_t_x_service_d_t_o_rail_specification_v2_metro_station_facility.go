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

// PTXServiceDTORailSpecificationV2MetroStationFacility StationFacility
//
// 捷運車站設施資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.Metro.StationFacility
type PTXServiceDTORailSpecificationV2MetroStationFacility struct {

	// Array
	//
	// 飲水機位置資訊
	// Required: true
	DrinkingFountains []*PTXServiceDTORailSpecificationV2MetroDrinkingFountain "json:\"DrinkingFountains\" xml:\"List`1\""

	// Array
	//
	// 無障礙電梯位置資訊
	// Required: true
	Elevators []*PTXServiceDTORailSpecificationV2MetroElevator "json:\"Elevators\" xml:\"List`1\""

	// Array
	//
	// 詢問處位置資訊
	// Required: true
	InformationSpots []*PTXServiceDTORailSpecificationV2MetroInformationSpot "json:\"InformationSpots\" xml:\"List`1\""

	// DateTime
	//
	// 來源端平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcUpdateTime *string `json:"SrcUpdateTime"`

	// String
	//
	// 車站代號
	// Required: true
	StationID *string `json:"StationID" xml:"String"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"StationName" xml:"NameType"`

	// Array
	//
	// 廁所位置資訊
	// Required: true
	Toilets []*PTXServiceDTORailSpecificationV2MetroToilet "json:\"Toilets\" xml:\"List`1\""

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

// Validate validates this p t x service d t o rail specification v2 metro station facility
func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrinkingFountains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElevators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInformationSpots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToilets(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateDrinkingFountains(formats strfmt.Registry) error {

	if err := validate.Required("DrinkingFountains", "body", m.DrinkingFountains); err != nil {
		return err
	}

	for i := 0; i < len(m.DrinkingFountains); i++ {
		if swag.IsZero(m.DrinkingFountains[i]) { // not required
			continue
		}

		if m.DrinkingFountains[i] != nil {
			if err := m.DrinkingFountains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DrinkingFountains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateElevators(formats strfmt.Registry) error {

	if err := validate.Required("Elevators", "body", m.Elevators); err != nil {
		return err
	}

	for i := 0; i < len(m.Elevators); i++ {
		if swag.IsZero(m.Elevators[i]) { // not required
			continue
		}

		if m.Elevators[i] != nil {
			if err := m.Elevators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Elevators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateInformationSpots(formats strfmt.Registry) error {

	if err := validate.Required("InformationSpots", "body", m.InformationSpots); err != nil {
		return err
	}

	for i := 0; i < len(m.InformationSpots); i++ {
		if swag.IsZero(m.InformationSpots[i]) { // not required
			continue
		}

		if m.InformationSpots[i] != nil {
			if err := m.InformationSpots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InformationSpots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateSrcUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcUpdateTime", "body", m.SrcUpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateToilets(formats strfmt.Registry) error {

	if err := validate.Required("Toilets", "body", m.Toilets); err != nil {
		return err
	}

	for i := 0; i < len(m.Toilets); i++ {
		if swag.IsZero(m.Toilets[i]) { // not required
			continue
		}

		if m.Toilets[i] != nil {
			if err := m.Toilets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Toilets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v2 metro station facility based on the context it is used
func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrinkingFountains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElevators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInformationSpots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToilets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) contextValidateDrinkingFountains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DrinkingFountains); i++ {

		if m.DrinkingFountains[i] != nil {
			if err := m.DrinkingFountains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DrinkingFountains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) contextValidateElevators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Elevators); i++ {

		if m.Elevators[i] != nil {
			if err := m.Elevators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Elevators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) contextValidateInformationSpots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InformationSpots); i++ {

		if m.InformationSpots[i] != nil {
			if err := m.InformationSpots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InformationSpots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) contextValidateToilets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Toilets); i++ {

		if m.Toilets[i] != nil {
			if err := m.Toilets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Toilets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2MetroStationFacility) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2MetroStationFacility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
