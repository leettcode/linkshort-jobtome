// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: link.proto

package link

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateLinkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLinkRequestMultiError, or nil if none found.
func (m *CreateLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetLongUri()) > 2048 {
		err := CreateLinkRequestValidationError{
			field:  "LongUri",
			reason: "value length must be at most 2048 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if uri, err := url.Parse(m.GetLongUri()); err != nil {
		err = CreateLinkRequestValidationError{
			field:  "LongUri",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := CreateLinkRequestValidationError{
			field:  "LongUri",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateLinkRequestMultiError(errors)
	}

	return nil
}

// CreateLinkRequestMultiError is an error wrapping multiple validation errors
// returned by CreateLinkRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLinkRequestMultiError) AllErrors() []error { return m }

// CreateLinkRequestValidationError is the validation error returned by
// CreateLinkRequest.Validate if the designated constraints aren't met.
type CreateLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLinkRequestValidationError) ErrorName() string {
	return "CreateLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLinkRequestValidationError{}

// Validate checks the field values on CreateLinkReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateLinkReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLinkReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLinkReplyMultiError, or nil if none found.
func (m *CreateLinkReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLinkReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShortUri

	if len(errors) > 0 {
		return CreateLinkReplyMultiError(errors)
	}

	return nil
}

// CreateLinkReplyMultiError is an error wrapping multiple validation errors
// returned by CreateLinkReply.ValidateAll() if the designated constraints
// aren't met.
type CreateLinkReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLinkReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLinkReplyMultiError) AllErrors() []error { return m }

// CreateLinkReplyValidationError is the validation error returned by
// CreateLinkReply.Validate if the designated constraints aren't met.
type CreateLinkReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLinkReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLinkReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLinkReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLinkReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLinkReplyValidationError) ErrorName() string { return "CreateLinkReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateLinkReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLinkReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLinkReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLinkReplyValidationError{}

// Validate checks the field values on RedirectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RedirectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedirectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RedirectRequestMultiError, or nil if none found.
func (m *RedirectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RedirectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetShortPath()); l < 1 || l > 2048 {
		err := RedirectRequestValidationError{
			field:  "ShortPath",
			reason: "value length must be between 1 and 2048 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_RedirectRequest_ShortPath_Pattern.MatchString(m.GetShortPath()) {
		err := RedirectRequestValidationError{
			field:  "ShortPath",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RedirectRequestMultiError(errors)
	}

	return nil
}

// RedirectRequestMultiError is an error wrapping multiple validation errors
// returned by RedirectRequest.ValidateAll() if the designated constraints
// aren't met.
type RedirectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedirectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedirectRequestMultiError) AllErrors() []error { return m }

// RedirectRequestValidationError is the validation error returned by
// RedirectRequest.Validate if the designated constraints aren't met.
type RedirectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedirectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedirectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedirectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedirectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedirectRequestValidationError) ErrorName() string { return "RedirectRequestValidationError" }

// Error satisfies the builtin error interface
func (e RedirectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedirectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedirectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedirectRequestValidationError{}

var _RedirectRequest_ShortPath_Pattern = regexp.MustCompile("^[a-zA-Z0-9_]*$")

// Validate checks the field values on RedirectReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RedirectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedirectReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RedirectReplyMultiError, or
// nil if none found.
func (m *RedirectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RedirectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LongUri

	if len(errors) > 0 {
		return RedirectReplyMultiError(errors)
	}

	return nil
}

// RedirectReplyMultiError is an error wrapping multiple validation errors
// returned by RedirectReply.ValidateAll() if the designated constraints
// aren't met.
type RedirectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedirectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedirectReplyMultiError) AllErrors() []error { return m }

// RedirectReplyValidationError is the validation error returned by
// RedirectReply.Validate if the designated constraints aren't met.
type RedirectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedirectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedirectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedirectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedirectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedirectReplyValidationError) ErrorName() string { return "RedirectReplyValidationError" }

// Error satisfies the builtin error interface
func (e RedirectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedirectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedirectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedirectReplyValidationError{}
