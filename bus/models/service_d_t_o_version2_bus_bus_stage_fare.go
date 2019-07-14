// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2BusBusStageFare BusStageFare
//
// 此計費方式以一路線內所有站牌分區收費。(公總稱之為計費站收費, Stage=計費站)
// swagger:model Service.DTO.Version2.Bus.BusStageFare
type ServiceDTOVersion2BusBusStageFare struct {

	// BusStage
	//
	// 訖點計費站
	// Required: true
	DestinationStage *ServiceDTOVersion2BusBusStage `json:"DestinationStage"`

	// 方向性描述
	// Required: true
	// Enum: [0: 去程 1: 返程 2: 迴圈 255: 未知]
	Direction *string `json:"Direction"`

	// 票價內容
	// Required: true
	Fares []*ServiceDTOVersion2BusBusFare `json:"Fares"`

	// BusStage
	//
	// 起點計費站
	// Required: true
	OriginStage *ServiceDTOVersion2BusBusStage `json:"OriginStage"`
}

// Validate validates this service d t o version2 bus bus stage fare
func (m *ServiceDTOVersion2BusBusStageFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusStageFare) validateDestinationStage(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStage", "body", m.DestinationStage); err != nil {
		return err
	}

	if m.DestinationStage != nil {
		if err := m.DestinationStage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DestinationStage")
			}
			return err
		}
	}

	return nil
}

var serviceDTOVersion2BusBusStageFareTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["0: 去程","1: 返程","2: 迴圈","255: 未知"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusStageFareTypeDirectionPropEnum = append(serviceDTOVersion2BusBusStageFareTypeDirectionPropEnum, v)
	}
}

const (

	// ServiceDTOVersion2BusBusStageFareDirectionNr0去程 captures enum value "0: 去程"
	ServiceDTOVersion2BusBusStageFareDirectionNr0去程 string = "0: 去程"

	// ServiceDTOVersion2BusBusStageFareDirectionNr1返程 captures enum value "1: 返程"
	ServiceDTOVersion2BusBusStageFareDirectionNr1返程 string = "1: 返程"

	// ServiceDTOVersion2BusBusStageFareDirectionNr2迴圈 captures enum value "2: 迴圈"
	ServiceDTOVersion2BusBusStageFareDirectionNr2迴圈 string = "2: 迴圈"

	// ServiceDTOVersion2BusBusStageFareDirectionNr255未知 captures enum value "255: 未知"
	ServiceDTOVersion2BusBusStageFareDirectionNr255未知 string = "255: 未知"
)

// prop value enum
func (m *ServiceDTOVersion2BusBusStageFare) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusStageFareTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusStageFare) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("Direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusStageFare) validateFares(formats strfmt.Registry) error {

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

func (m *ServiceDTOVersion2BusBusStageFare) validateOriginStage(formats strfmt.Registry) error {

	if err := validate.Required("OriginStage", "body", m.OriginStage); err != nil {
		return err
	}

	if m.OriginStage != nil {
		if err := m.OriginStage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginStage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusStageFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusStageFare) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusBusStageFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
