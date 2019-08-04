// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MOTCAPIBusDALBusServiceDTOVersion3BusDepot BusDepotList
// swagger:model MOTC.API.Bus.DAL.Bus[Service.DTO.Version3.Bus.Depot]
type MOTCAPIBusDALBusServiceDTOVersion3BusDepot struct {

	// 業管機關簡碼
	AuthorityCode string `json:"AuthorityCode,omitempty"`

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// 資料(陣列)
	Depots []*ServiceDTOVersion3BusDepot `json:"Depots"`

	// [來源端平臺]資料更新週期
	SrcUpdateInterval int32 `json:"SrcUpdateInterval,omitempty"`

	// DateTime
	//
	// [來源端平臺]資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	SrcUpdateTime string `json:"SrcUpdateTime,omitempty"`

	// [平臺]資料更新週期(秒)
	UpdateInterval int32 `json:"UpdateInterval,omitempty"`

	// DateTime
	//
	// [平臺]資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	UpdateTime string `json:"UpdateTime,omitempty"`

	// 資料版本編號
	VersionID int32 `json:"VersionID,omitempty"`
}

// Validate validates this m o t c API bus d a l bus service d t o version3 bus depot
func (m *MOTCAPIBusDALBusServiceDTOVersion3BusDepot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MOTCAPIBusDALBusServiceDTOVersion3BusDepot) validateDepots(formats strfmt.Registry) error {

	if swag.IsZero(m.Depots) { // not required
		return nil
	}

	for i := 0; i < len(m.Depots); i++ {
		if swag.IsZero(m.Depots[i]) { // not required
			continue
		}

		if m.Depots[i] != nil {
			if err := m.Depots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Depots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MOTCAPIBusDALBusServiceDTOVersion3BusDepot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MOTCAPIBusDALBusServiceDTOVersion3BusDepot) UnmarshalBinary(b []byte) error {
	var res MOTCAPIBusDALBusServiceDTOVersion3BusDepot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
