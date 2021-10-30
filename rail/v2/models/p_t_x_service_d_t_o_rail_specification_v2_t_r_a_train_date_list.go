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

// PTXServiceDTORailSpecificationV2TRATrainDateList TrainDateList
//
// 臺鐵每日時刻表供應日期資料型別
//
// swagger:model PTX.Service.DTO.Rail.Specification.V2.TRA.TrainDateList
type PTXServiceDTORailSpecificationV2TRATrainDateList struct {

	// String
	//
	// 每日時刻表供應最終日期(格式: yyyy-MM-dd)
	// Required: true
	EndDate *string `json:"EndDate" xml:"String"`

	// String
	//
	// 每日時刻表供應起始日期(格式: yyyy-MM-dd)
	// Required: true
	StartDate *string `json:"StartDate" xml:"String"`

	// Array
	//
	// 每日時刻表供應日期清單資料
	// Required: true
	TrainDates []string "json:\"TrainDates\" xml:\"List`1\""

	// DateTime
	//
	// 資料更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"UpdateTime"`
}

// Validate validates this p t x service d t o rail specification v2 t r a train date list
func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainDates(formats); err != nil {
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

func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("EndDate", "body", m.EndDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("StartDate", "body", m.StartDate); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) validateTrainDates(formats strfmt.Registry) error {

	if err := validate.Required("TrainDates", "body", m.TrainDates); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) validateUpdateTime(formats strfmt.Registry) error {

	if err := validate.Required("UpdateTime", "body", m.UpdateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdateTime", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p t x service d t o rail specification v2 t r a train date list based on context it is used
func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV2TRATrainDateList) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV2TRATrainDateList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
