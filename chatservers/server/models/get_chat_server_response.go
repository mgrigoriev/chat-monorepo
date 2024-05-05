// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetChatServerResponse get chat server response
//
// swagger:model GetChatServerResponse
type GetChatServerResponse struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name
	// Example: chatserver123
	Name string `json:"name,omitempty"`
}

// Validate validates this get chat server response
func (m *GetChatServerResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get chat server response based on context it is used
func (m *GetChatServerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetChatServerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetChatServerResponse) UnmarshalBinary(b []byte) error {
	var res GetChatServerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
