// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailTHSRAvailableSeat AvailableSeat
//
// 高鐵對號座位狀態資訊
//
// swagger:model Service.DTO.Version2.Rail.THSR.AvailableSeat
type ServiceDTOVersion2RailTHSRAvailableSeat struct {

	// 發車時間(格式: HH:mm)
	// Required: true
	DepartureTime *string `json:"DepartureTime"`

	// integer
	//
	// 方向 : [0:'南下',1:'北上']
	// Required: true
	Direction *int32 `json:"Direction"`

	// 終點車站代碼
	// Required: true
	EndingStationID *string `json:"EndingStationID"`

	// NameType
	//
	// 終點車站名稱
	// Required: true
	EndingStationName *ServiceDTOVersion2BaseNameType `json:"EndingStationName"`

	// DateTime
	//
	// 來源端平台接收時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	SrcRecTime *string `json:"SrcRecTime"`

	// 查詢車站代碼
	// Required: true
	StationID *string `json:"StationID"`

	// NameType
	//
	// 查詢車站名稱
	// Required: true
	StationName *ServiceDTOVersion2BaseNameType `json:"StationName"`

	// 車次停靠站點組合
	// Required: true
	StopStations []*ServiceDTOVersion2RailTHSRStopStation `json:"StopStations"`

	// 車次號碼
	// Required: true
	TrainNo *string `json:"TrainNo"`

	// DateTime
	//
	// 本平台資料更新時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`
}

// Validate validates this service d t o version2 rail t h s r available seat
func (m *ServiceDTOVersion2RailTHSRAvailableSeat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcRecTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopStations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
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

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateDepartureTime(formats strfmt.Registry) error {

	if err := validate.Required("DepartureTime", "body", m.DepartureTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateEndingStationID(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationID", "body", m.EndingStationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateEndingStationName(formats strfmt.Registry) error {

	if err := validate.Required("EndingStationName", "body", m.EndingStationName); err != nil {
		return err
	}

	if m.EndingStationName != nil {
		if err := m.EndingStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndingStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateSrcRecTime(formats strfmt.Registry) error {

	if err := validate.Required("SrcRecTime", "body", m.SrcRecTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("StationName", "body", m.StationName); err != nil {
		return err
	}

	if m.StationName != nil {
		if err := m.StationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateStopStations(formats strfmt.Registry) error {

	if err := validate.Required("StopStations", "body", m.StopStations); err != nil {
		return err
	}

	for i := 0; i < len(m.StopStations); i++ {
		if swag.IsZero(m.StopStations[i]) { // not required
			continue
		}

		if m.StopStations[i] != nil {
			if err := m.StopStations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StopStations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTHSRAvailableSeat) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSRAvailableSeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTHSRAvailableSeat) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTHSRAvailableSeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
