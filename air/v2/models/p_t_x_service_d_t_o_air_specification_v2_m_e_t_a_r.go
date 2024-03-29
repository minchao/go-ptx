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

// PTXServiceDTOAirSpecificationV2METAR METAR
//
// swagger:model PTX.Service.DTO.Air.Specification.V2.METAR
type PTXServiceDTOAirSpecificationV2METAR struct {

	// String
	//
	// 機場代碼(IATA)
	// Required: true
	AirportID *string `json:"AirportID" xml:"AirportID"`

	// NameType
	//
	// 機場名稱
	// Required: true
	AirportName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"AirportName" xml:"NameType"`

	// String
	//
	// 雲冪(含單位)
	// Required: true
	Ceiling *string `json:"Ceiling" xml:"Ceiling"`

	// NameType
	//
	// 機場所屬城市
	// Required: true
	CityName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"CityName" xml:"NameType"`

	// NameType
	//
	// 國家名稱
	// Required: true
	CountryName struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"CountryName" xml:"NameType"`

	// String
	//
	// 機場天氣報告(METAR/SPECI)
	// Required: true
	MetarText *string `json:"MetarText" xml:"MetarText"`

	// String
	//
	// 機場天氣報告時間(地方時)
	// Required: true
	MetarTime *string `json:"MetarTime" xml:"MetarTime"`

	// DateTime
	//
	// 觀測時間
	// Required: true
	// Format: date-time
	ObservationTime *strfmt.DateTime `json:"ObservationTime"`

	// String
	//
	// 機場氣象觀測站代碼(ICAO)
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// PointType
	//
	// 機場氣象觀測站座標
	// Required: true
	StationPosition struct {
		PTXServiceDTOSharedSpecificationV2BasePointType
	} `json:"StationPosition" xml:"PointType"`

	// String
	//
	// 溫度(含單位)
	// Required: true
	Temperature *string `json:"Temperature" xml:"Temperature"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`

	// String
	//
	// 盛行能見度(含單位，9999以10公里以上表示)
	// Required: true
	Visibility *string `json:"Visibility" xml:"Visibility"`

	// NameType
	//
	// 天氣描述
	// Required: true
	WeatherDescription struct {
		PTXServiceDTOSharedSpecificationV2BaseNameType
	} `json:"WeatherDescription" xml:"NameType"`

	// String
	//
	// 風向(含單位)
	// Required: true
	WindDirection *string `json:"WindDirection" xml:"WindDirection"`

	// String
	//
	// 風速(含單位)
	// Required: true
	WindSpeed *string `json:"WindSpeed" xml:"WindSpeed"`
}

// Validate validates this p t x service d t o air specification v2 m e t a r
func (m *PTXServiceDTOAirSpecificationV2METAR) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAirportName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCeiling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCityName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetarText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetarTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemperature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeatherDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindSpeed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateAirportID(formats strfmt.Registry) error {

	if err := validate.Required("AirportID", "body", m.AirportID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateAirportName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateCeiling(formats strfmt.Registry) error {

	if err := validate.Required("Ceiling", "body", m.Ceiling); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateCityName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateCountryName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateMetarText(formats strfmt.Registry) error {

	if err := validate.Required("MetarText", "body", m.MetarText); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateMetarTime(formats strfmt.Registry) error {

	if err := validate.Required("MetarTime", "body", m.MetarTime); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateObservationTime(formats strfmt.Registry) error {

	if err := validate.Required("ObservationTime", "body", m.ObservationTime); err != nil {
		return err
	}

	if err := validate.FormatOf("ObservationTime", "body", "date-time", m.ObservationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateStationPosition(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateTemperature(formats strfmt.Registry) error {

	if err := validate.Required("Temperature", "body", m.Temperature); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("Visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateWeatherDescription(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateWindDirection(formats strfmt.Registry) error {

	if err := validate.Required("WindDirection", "body", m.WindDirection); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) validateWindSpeed(formats strfmt.Registry) error {

	if err := validate.Required("WindSpeed", "body", m.WindSpeed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o air specification v2 m e t a r based on the context it is used
func (m *PTXServiceDTOAirSpecificationV2METAR) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAirportName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCityName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountryName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStationPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeatherDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) contextValidateAirportName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) contextValidateCityName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) contextValidateCountryName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) contextValidateStationPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOAirSpecificationV2METAR) contextValidateWeatherDescription(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2METAR) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOAirSpecificationV2METAR) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOAirSpecificationV2METAR
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
