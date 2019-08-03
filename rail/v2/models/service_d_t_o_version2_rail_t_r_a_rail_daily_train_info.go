// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2RailTRARailDailyTrainInfo RailDailyTrainInfo
//
// 臺鐵車次資料型別
// swagger:model Service.DTO.Version2.Rail.TRA.RailDailyTrainInfo
type ServiceDTOVersion2RailTRARailDailyTrainInfo struct {

	// integer
	//
	// 是否人車同行班次(置於攜車袋之自行車各級列車均可乘車) : [0:'否',1:'是']
	// Required: true
	BikeFlag *int32 `json:"BikeFlag"`

	// integer
	//
	// 是否設有哺(集)乳室車廂 : [0:'否',1:'是']
	// Required: true
	BreastFeedingFlag *int32 `json:"BreastFeedingFlag"`

	// integer
	//
	// 是否每日行駛 : [0:'否',1:'是']
	// Required: true
	DailyFlag *int32 `json:"DailyFlag"`

	// integer
	//
	// 是否提供餐車服務 : [0:'否',1:'是']
	// Required: true
	DiningFlag *int32 `json:"DiningFlag"`

	// integer
	//
	// 順逆行 : [0:'順行',1:'逆行']
	// Required: true
	Direction *int32 `json:"Direction"`

	// 列車終點車站代號
	EndingStationID string `json:"EndingStationID,omitempty"`

	// NameType
	//
	// 列車終點車站名稱
	EndingStationName *ServiceDTOVersion2BaseNameType `json:"EndingStationName,omitempty"`

	// NameType
	//
	// 附註說明
	Note *ServiceDTOVersion2BaseNameType `json:"Note,omitempty"`

	// 跨夜車站代碼
	OverNightStationID string `json:"OverNightStationID,omitempty"`

	// integer
	//
	// 是否提供行李服務 : [0:'否',1:'是']
	// Required: true
	PackageServiceFlag *int32 `json:"PackageServiceFlag"`

	// integer
	//
	// 是否為加班車 : [0:'否',1:'是']
	// Required: true
	ServiceAddedFlag *int32 `json:"ServiceAddedFlag"`

	// 列車起點車站代號
	StartingStationID string `json:"StartingStationID,omitempty"`

	// NameType
	//
	// 列車起點車站名稱
	StartingStationName *ServiceDTOVersion2BaseNameType `json:"StartingStationName,omitempty"`

	// 車次代碼
	// Required: true
	TrainNo *string `json:"TrainNo"`

	// 列車車種簡碼
	// Required: true
	TrainTypeCode *string `json:"TrainTypeCode"`

	// 列車車種代碼
	// Required: true
	TrainTypeID *string `json:"TrainTypeID"`

	// NameType
	//
	// 列車車種名稱
	// Required: true
	TrainTypeName *ServiceDTOVersion2BaseNameType `json:"TrainTypeName"`

	// 車次車頭文字描述(通用以"往"+ 迄站中文站名")
	TripHeadsign string `json:"TripHeadsign,omitempty"`

	// integer
	//
	// 山海線類型 : [0:'不經山海線',1:'山線',2:'海線']
	TripLine int32 `json:"TripLine,omitempty"`

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	UpdateTime *string `json:"UpdateTime"`

	// 資料版本編號
	// Required: true
	VersionID *int32 `json:"VersionID"`

	// integer
	//
	// 是否設身障旅客專用座位車 : [0:'否',1:'是']
	// Required: true
	WheelchairFlag *int32 `json:"WheelchairFlag"`
}

// Validate validates this service d t o version2 rail t r a rail daily train info
func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBikeFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBreastFeedingFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDailyFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiningFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageServiceFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAddedFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTypeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainTypeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWheelchairFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateBikeFlag(formats strfmt.Registry) error {

	if err := validate.Required("BikeFlag", "body", m.BikeFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateBreastFeedingFlag(formats strfmt.Registry) error {

	if err := validate.Required("BreastFeedingFlag", "body", m.BreastFeedingFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateDailyFlag(formats strfmt.Registry) error {

	if err := validate.Required("DailyFlag", "body", m.DailyFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateDiningFlag(formats strfmt.Registry) error {

	if err := validate.Required("DiningFlag", "body", m.DiningFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateEndingStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.EndingStationName) { // not required
		return nil
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

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if m.Note != nil {
		if err := m.Note.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Note")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validatePackageServiceFlag(formats strfmt.Registry) error {

	if err := validate.Required("PackageServiceFlag", "body", m.PackageServiceFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateServiceAddedFlag(formats strfmt.Registry) error {

	if err := validate.Required("ServiceAddedFlag", "body", m.ServiceAddedFlag); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateStartingStationName(formats strfmt.Registry) error {

	if swag.IsZero(m.StartingStationName) { // not required
		return nil
	}

	if m.StartingStationName != nil {
		if err := m.StartingStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StartingStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateTrainNo(formats strfmt.Registry) error {

	if err := validate.Required("TrainNo", "body", m.TrainNo); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateTrainTypeCode(formats strfmt.Registry) error {

	if err := validate.Required("TrainTypeCode", "body", m.TrainTypeCode); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateTrainTypeID(formats strfmt.Registry) error {

	if err := validate.Required("TrainTypeID", "body", m.TrainTypeID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateTrainTypeName(formats strfmt.Registry) error {

	if err := validate.Required("TrainTypeName", "body", m.TrainTypeName); err != nil {
		return err
	}

	if m.TrainTypeName != nil {
		if err := m.TrainTypeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TrainTypeName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateVersionID(formats strfmt.Registry) error {

	if err := validate.Required("VersionID", "body", m.VersionID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) validateWheelchairFlag(formats strfmt.Registry) error {

	if err := validate.Required("WheelchairFlag", "body", m.WheelchairFlag); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2RailTRARailDailyTrainInfo) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2RailTRARailDailyTrainInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
