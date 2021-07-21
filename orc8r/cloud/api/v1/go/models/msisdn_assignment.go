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

// MsisdnAssignment msisdn assignment
//
// swagger:model msisdn_assignment
type MsisdnAssignment struct {

	// id
	// Required: true
	ID SubscriberID `json:"id"`

	// msisdn
	// Required: true
	Msisdn Msisdn `json:"msisdn"`
}

// Validate validates this msisdn assignment
func (m *MsisdnAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsisdn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsisdnAssignment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", SubscriberID(m.ID)); err != nil {
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

func (m *MsisdnAssignment) validateMsisdn(formats strfmt.Registry) error {

	if err := validate.Required("msisdn", "body", Msisdn(m.Msisdn)); err != nil {
		return err
	}

	if err := m.Msisdn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("msisdn")
		}
		return err
	}

	return nil
}

// ContextValidate validate this msisdn assignment based on the context it is used
func (m *MsisdnAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsisdn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsisdnAssignment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MsisdnAssignment) contextValidateMsisdn(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Msisdn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("msisdn")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsisdnAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsisdnAssignment) UnmarshalBinary(b []byte) error {
	var res MsisdnAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
