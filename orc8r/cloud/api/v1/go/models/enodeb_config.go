// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ENODEBConfig Configuration for managed / unmanaged eNodeb
//
// swagger:model enodeb_config
type ENODEBConfig struct {

	// config type
	// Example: MANAGED
	// Required: true
	// Enum: [MANAGED UNMANAGED]
	ConfigType string `json:"config_type"`

	// managed config
	ManagedConfig *ENODEBConfiguration `json:"managed_config,omitempty"`

	// unmanaged config
	UnmanagedConfig *UnmanagedENODEBConfiguration `json:"unmanaged_config,omitempty"`
}

// Validate validates this enodeb config
func (m *ENODEBConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnmanagedConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var enodebConfigTypeConfigTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANAGED","UNMANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enodebConfigTypeConfigTypePropEnum = append(enodebConfigTypeConfigTypePropEnum, v)
	}
}

const (

	// ENODEBConfigConfigTypeMANAGED captures enum value "MANAGED"
	ENODEBConfigConfigTypeMANAGED string = "MANAGED"

	// ENODEBConfigConfigTypeUNMANAGED captures enum value "UNMANAGED"
	ENODEBConfigConfigTypeUNMANAGED string = "UNMANAGED"
)

// prop value enum
func (m *ENODEBConfig) validateConfigTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, enodebConfigTypeConfigTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ENODEBConfig) validateConfigType(formats strfmt.Registry) error {

	if err := validate.RequiredString("config_type", "body", m.ConfigType); err != nil {
		return err
	}

	// value enum
	if err := m.validateConfigTypeEnum("config_type", "body", m.ConfigType); err != nil {
		return err
	}

	return nil
}

func (m *ENODEBConfig) validateManagedConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagedConfig) { // not required
		return nil
	}

	if m.ManagedConfig != nil {
		if err := m.ManagedConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_config")
			}
			return err
		}
	}

	return nil
}

func (m *ENODEBConfig) validateUnmanagedConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.UnmanagedConfig) { // not required
		return nil
	}

	if m.UnmanagedConfig != nil {
		if err := m.UnmanagedConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unmanaged_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this enodeb config based on the context it is used
func (m *ENODEBConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManagedConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnmanagedConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ENODEBConfig) contextValidateManagedConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagedConfig != nil {
		if err := m.ManagedConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managed_config")
			}
			return err
		}
	}

	return nil
}

func (m *ENODEBConfig) contextValidateUnmanagedConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.UnmanagedConfig != nil {
		if err := m.UnmanagedConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unmanaged_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ENODEBConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ENODEBConfig) UnmarshalBinary(b []byte) error {
	var res ENODEBConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
