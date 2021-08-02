// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	models2 "magma/lte/cloud/go/services/lte/obsidian/models"
	models4 "magma/orc8r/cloud/go/models"
	models5 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FegLteNetwork Federated LTE Network spec
//
// swagger:model feg_lte_network
type FegLteNetwork struct {

	// cellular
	// Required: true
	Cellular *models2.NetworkCellularConfigs `json:"cellular"`

	// description
	// Required: true
	Description models4.NetworkDescription `json:"description"`

	// dns
	// Required: true
	DNS *models5.NetworkDNSConfig `json:"dns"`

	// features
	Features *models5.NetworkFeatures `json:"features,omitempty"`

	// federation
	// Required: true
	Federation *FederatedNetworkConfigs `json:"federation"`

	// id
	// Required: true
	ID models4.NetworkID `json:"id"`

	// name
	// Required: true
	Name models4.NetworkName `json:"name"`
}

// Validate validates this feg lte network
func (m *FegLteNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCellular(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFederation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FegLteNetwork) validateCellular(formats strfmt.Registry) error {

	if err := validate.Required("cellular", "body", m.Cellular); err != nil {
		return err
	}

	if m.Cellular != nil {
		if err := m.Cellular.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellular")
			}
			return err
		}
	}

	return nil
}

func (m *FegLteNetwork) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", models4.NetworkDescription(m.Description)); err != nil {
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

func (m *FegLteNetwork) validateDNS(formats strfmt.Registry) error {

	if err := validate.Required("dns", "body", m.DNS); err != nil {
		return err
	}

	if m.DNS != nil {
		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *FegLteNetwork) validateFeatures(formats strfmt.Registry) error {
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

func (m *FegLteNetwork) validateFederation(formats strfmt.Registry) error {

	if err := validate.Required("federation", "body", m.Federation); err != nil {
		return err
	}

	if m.Federation != nil {
		if err := m.Federation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("federation")
			}
			return err
		}
	}

	return nil
}

func (m *FegLteNetwork) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", models4.NetworkID(m.ID)); err != nil {
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

func (m *FegLteNetwork) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", models4.NetworkName(m.Name)); err != nil {
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

// ContextValidate validate this feg lte network based on the context it is used
func (m *FegLteNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCellular(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFederation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FegLteNetwork) contextValidateCellular(ctx context.Context, formats strfmt.Registry) error {

	if m.Cellular != nil {
		if err := m.Cellular.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellular")
			}
			return err
		}
	}

	return nil
}

func (m *FegLteNetwork) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Description.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *FegLteNetwork) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	if m.DNS != nil {
		if err := m.DNS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *FegLteNetwork) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FegLteNetwork) contextValidateFederation(ctx context.Context, formats strfmt.Registry) error {

	if m.Federation != nil {
		if err := m.Federation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("federation")
			}
			return err
		}
	}

	return nil
}

func (m *FegLteNetwork) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *FegLteNetwork) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FegLteNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FegLteNetwork) UnmarshalBinary(b []byte) error {
	var res FegLteNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
