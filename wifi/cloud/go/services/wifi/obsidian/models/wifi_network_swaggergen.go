// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	models1 "magma/orc8r/cloud/go/models"
	models2 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WifiNetwork Wifi Network spec
//
// swagger:model wifi_network
type WifiNetwork struct {

	// description
	// Required: true
	Description models1.NetworkDescription `json:"description"`

	// features
	Features *models2.NetworkFeatures `json:"features,omitempty"`

	// id
	// Required: true
	ID models1.NetworkID `json:"id"`

	// name
	// Required: true
	Name models1.NetworkName `json:"name"`

	// wifi
	// Required: true
	Wifi *NetworkWifiConfigs `json:"wifi"`
}

// Validate validates this wifi network
func (m *WifiNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWifi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiNetwork) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", NetworkDescription(m.Description)); err != nil {
		return err
	}

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *WifiNetwork) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", NetworkID(m.ID)); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", NetworkName(m.Name)); err != nil {
		return err
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) validateWifi(formats strfmt.Registry) error {

	if err := validate.Required("wifi", "body", m.Wifi); err != nil {
		return err
	}

	if m.Wifi != nil {
		if err := m.Wifi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wifi")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this wifi network based on the context it is used
func (m *WifiNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWifi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiNetwork) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Description.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	if m.Features != nil {
		if err := m.Features.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *WifiNetwork) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) contextValidateWifi(ctx context.Context, formats strfmt.Registry) error {

	if m.Wifi != nil {
		if err := m.Wifi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wifi")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WifiNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WifiNetwork) UnmarshalBinary(b []byte) error {
	var res WifiNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
