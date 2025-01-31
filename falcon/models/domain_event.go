// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainEvent domain event
//
// swagger:model domain.Event
type DomainEvent struct {

	// The raw body of the event
	// Required: true
	Body *string `json:"body"`

	// By default, event bodies are truncated to 64kb and bodyIsTruncated is set to True. For event bodies larger than 64kb, call the /events-full-body endpoint with the respective eventId
	// Required: true
	BodyIsTruncated *bool `json:"body_is_truncated"`

	// Link to the event, can be missing
	BodyLink string `json:"body_link,omitempty"`

	// botnet config source
	BotnetConfigSource *DomainBotnetConfigSource `json:"botnet_config_source,omitempty"`

	// The date the event was created (in UTC format)
	// Required: true
	CreatedDate *string `json:"created_date"`

	// ddos attack source
	DdosAttackSource *DomainDDOSAttackSource `json:"ddos_attack_source,omitempty"`

	// The type of event. One of `TweetEvent`, `CodePasteEvent`, `BotnetConfigEvent`, `DdosAttackEvent`
	// Required: true
	EventType *string `json:"event_type"`

	// The event's fingerprint
	// Required: true
	Fingerprint *string `json:"fingerprint"`

	// The unique event ID
	// Required: true
	ID *string `json:"id"`

	// List of objects with rules that matched the event
	MatchedRules []*DomainMatchedRule `json:"matched_rules"`

	// pastebin text source
	PastebinTextSource *DomainPastebinTextSource `json:"pastebin_text_source,omitempty"`

	// A list of tags summarizing event content
	Tags []string `json:"tags"`

	// tweet source
	TweetSource *DomainTweetSource `json:"tweet_source,omitempty"`

	// The date the event was last updated (in UTC format)
	// Required: true
	UpdatedDate *string `json:"updated_date"`
}

// Validate validates this domain event
func (m *DomainEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBodyIsTruncated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBotnetConfigSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDdosAttackSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchedRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePastebinTextSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTweetSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainEvent) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *DomainEvent) validateBodyIsTruncated(formats strfmt.Registry) error {

	if err := validate.Required("body_is_truncated", "body", m.BodyIsTruncated); err != nil {
		return err
	}

	return nil
}

func (m *DomainEvent) validateBotnetConfigSource(formats strfmt.Registry) error {
	if swag.IsZero(m.BotnetConfigSource) { // not required
		return nil
	}

	if m.BotnetConfigSource != nil {
		if err := m.BotnetConfigSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("botnet_config_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("botnet_config_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("created_date", "body", m.CreatedDate); err != nil {
		return err
	}

	return nil
}

func (m *DomainEvent) validateDdosAttackSource(formats strfmt.Registry) error {
	if swag.IsZero(m.DdosAttackSource) { // not required
		return nil
	}

	if m.DdosAttackSource != nil {
		if err := m.DdosAttackSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ddos_attack_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ddos_attack_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", m.EventType); err != nil {
		return err
	}

	return nil
}

func (m *DomainEvent) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *DomainEvent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainEvent) validateMatchedRules(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchedRules) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchedRules); i++ {
		if swag.IsZero(m.MatchedRules[i]) { // not required
			continue
		}

		if m.MatchedRules[i] != nil {
			if err := m.MatchedRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matched_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matched_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainEvent) validatePastebinTextSource(formats strfmt.Registry) error {
	if swag.IsZero(m.PastebinTextSource) { // not required
		return nil
	}

	if m.PastebinTextSource != nil {
		if err := m.PastebinTextSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pastebin_text_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pastebin_text_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) validateTweetSource(formats strfmt.Registry) error {
	if swag.IsZero(m.TweetSource) { // not required
		return nil
	}

	if m.TweetSource != nil {
		if err := m.TweetSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tweet_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tweet_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) validateUpdatedDate(formats strfmt.Registry) error {

	if err := validate.Required("updated_date", "body", m.UpdatedDate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain event based on the context it is used
func (m *DomainEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBotnetConfigSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDdosAttackSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatchedRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePastebinTextSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTweetSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainEvent) contextValidateBotnetConfigSource(ctx context.Context, formats strfmt.Registry) error {

	if m.BotnetConfigSource != nil {
		if err := m.BotnetConfigSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("botnet_config_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("botnet_config_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) contextValidateDdosAttackSource(ctx context.Context, formats strfmt.Registry) error {

	if m.DdosAttackSource != nil {
		if err := m.DdosAttackSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ddos_attack_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ddos_attack_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) contextValidateMatchedRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchedRules); i++ {

		if m.MatchedRules[i] != nil {
			if err := m.MatchedRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matched_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matched_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainEvent) contextValidatePastebinTextSource(ctx context.Context, formats strfmt.Registry) error {

	if m.PastebinTextSource != nil {
		if err := m.PastebinTextSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pastebin_text_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pastebin_text_source")
			}
			return err
		}
	}

	return nil
}

func (m *DomainEvent) contextValidateTweetSource(ctx context.Context, formats strfmt.Registry) error {

	if m.TweetSource != nil {
		if err := m.TweetSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tweet_source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tweet_source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainEvent) UnmarshalBinary(b []byte) error {
	var res DomainEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
