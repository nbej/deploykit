package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Poller20PartialPoller A poller for periodic collection of telemetry data
// swagger:model poller.2.0_PartialPoller
type Poller20PartialPoller struct {

	// Poller configuration object
	Config interface{} `json:"config,omitempty"`

	// Asserted if poller is paused
	Paused bool `json:"paused,omitempty"`

	// Interval at which poller will run
	PollInterval float64 `json:"pollInterval,omitempty"`

	// Type of poller
	Type string `json:"type,omitempty"`
}

// Validate validates this poller 2 0 partial poller
func (m *Poller20PartialPoller) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var poller20PartialPollerTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipmi","snmp","redfish"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		poller20PartialPollerTypeTypePropEnum = append(poller20PartialPollerTypeTypePropEnum, v)
	}
}

const (
	// Poller20PartialPollerTypeIPMI captures enum value "ipmi"
	Poller20PartialPollerTypeIPMI string = "ipmi"
	// Poller20PartialPollerTypeSnmp captures enum value "snmp"
	Poller20PartialPollerTypeSnmp string = "snmp"
	// Poller20PartialPollerTypeRedfish captures enum value "redfish"
	Poller20PartialPollerTypeRedfish string = "redfish"
)

// prop value enum
func (m *Poller20PartialPoller) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, poller20PartialPollerTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Poller20PartialPoller) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
