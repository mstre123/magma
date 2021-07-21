// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GatewayPoolID gateway pool id
// Example: pool1
//
// swagger:model gateway_pool_id
type GatewayPoolID string

// Validate validates this gateway pool id
func (m GatewayPoolID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.Pattern("", "body", string(m), `^[a-z][\da-z_-]+$`); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gateway pool id based on context it is used
func (m GatewayPoolID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
