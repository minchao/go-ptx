// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2ApplicationMETAR METAR
//
// swagger:model Service.DTO.Version2.Application.METAR
type ServiceDTOVersion2ApplicationMETAR struct {

	// 機場代碼(IATA)
	// Required: true
	AirportID *string `json:"AirportID"`

	// NameType
	//
	// 機場名稱
	// Required: true
	AirportName *ServiceDTOVersion2BaseNameType `json:"AirportName"`

	// 雲冪(含單位)
	// Required: true
	Ceiling *string `json:"Ceiling"`

	// NameType
	//
	// 機場所屬城市
	// Required: true
	CityName *ServiceDTOVersion2BaseNameType `json:"CityName"`

	// NameType
	//
	// 國家名稱
	// Required: true
	CountryName *ServiceDTOVersion2BaseNameType `json:"CountryName"`

	// 機場天氣報告(METAR/SPECI)
	// Required: true
	MetarText *string `json:"MetarText"`

	// 機場天氣報告時間(地方時)
	// Required: true
	MetarTime *string `json:"MetarTime"`

	// DateTime
	//
	// 觀測時間
	// Required: true
	ObservationTime *string `json:"ObservationTime"`

	// 機場氣象觀測站代碼(ICAO)
	// Required: true
	StationID *string `json:"StationID"`

	// PointType
	//
	// 機場氣象觀測站座標
	// Required: true
	StationPosition *ServiceDTOVersion2BasePointType `json:"StationPosition"`

	// 溫度(含單位)
	// Required: true
	Temperature *string `json:"Temperature"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 盛行能見度(含單位，9999以10公里以上表示)
	// Required: true
	Visibility *string `json:"Visibility"`

	// NameType
	//
	// 天氣描述
	// Required: true
	WeatherDescription *ServiceDTOVersion2BaseNameType `json:"WeatherDescription"`

	// 風向(含單位)
	// Required: true
	WindDirection *string `json:"WindDirection"`

	// 風速(含單位)
	// Required: true
	WindSpeed *string `json:"WindSpeed"`
}

// Validate validates this service d t o version2 application m e t a r
func (m *ServiceDTOVersion2ApplicationMETAR) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion2ApplicationMETAR) validateAirportID(formats strfmt.Registry) error {

	if err := validate.Required("AirportID", "body", m.AirportID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateAirportName(formats strfmt.Registry) error {

	if err := validate.Required("AirportName", "body", m.AirportName); err != nil {
		return err
	}

	if m.AirportName != nil {
		if err := m.AirportName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AirportName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateCeiling(formats strfmt.Registry) error {

	if err := validate.Required("Ceiling", "body", m.Ceiling); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateCityName(formats strfmt.Registry) error {

	if err := validate.Required("CityName", "body", m.CityName); err != nil {
		return err
	}

	if m.CityName != nil {
		if err := m.CityName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CityName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateCountryName(formats strfmt.Registry) error {

	if err := validate.Required("CountryName", "body", m.CountryName); err != nil {
		return err
	}

	if m.CountryName != nil {
		if err := m.CountryName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CountryName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateMetarText(formats strfmt.Registry) error {

	if err := validate.Required("MetarText", "body", m.MetarText); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateMetarTime(formats strfmt.Registry) error {

	if err := validate.Required("MetarTime", "body", m.MetarTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateObservationTime(formats strfmt.Registry) error {

	if err := validate.Required("ObservationTime", "body", m.ObservationTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateStationPosition(formats strfmt.Registry) error {

	if err := validate.Required("StationPosition", "body", m.StationPosition); err != nil {
		return err
	}

	if m.StationPosition != nil {
		if err := m.StationPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationPosition")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateTemperature(formats strfmt.Registry) error {

	if err := validate.Required("Temperature", "body", m.Temperature); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("Visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateWeatherDescription(formats strfmt.Registry) error {

	if err := validate.Required("WeatherDescription", "body", m.WeatherDescription); err != nil {
		return err
	}

	if m.WeatherDescription != nil {
		if err := m.WeatherDescription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WeatherDescription")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateWindDirection(formats strfmt.Registry) error {

	if err := validate.Required("WindDirection", "body", m.WindDirection); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2ApplicationMETAR) validateWindSpeed(formats strfmt.Registry) error {

	if err := validate.Required("WindSpeed", "body", m.WindSpeed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2ApplicationMETAR) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2ApplicationMETAR) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2ApplicationMETAR
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
