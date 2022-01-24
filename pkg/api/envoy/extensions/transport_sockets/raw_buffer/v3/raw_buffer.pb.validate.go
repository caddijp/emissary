// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/raw_buffer/v3/raw_buffer.proto

package raw_bufferv3

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

// Validate checks the field values on RawBuffer with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RawBuffer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RawBuffer with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RawBufferMultiError, or nil
// if none found.
func (m *RawBuffer) ValidateAll() error {
	return m.validate(true)
}

func (m *RawBuffer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RawBufferMultiError(errors)
	}
	return nil
}

// RawBufferMultiError is an error wrapping multiple validation errors returned
// by RawBuffer.ValidateAll() if the designated constraints aren't met.
type RawBufferMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RawBufferMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RawBufferMultiError) AllErrors() []error { return m }

// RawBufferValidationError is the validation error returned by
// RawBuffer.Validate if the designated constraints aren't met.
type RawBufferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RawBufferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RawBufferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RawBufferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RawBufferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RawBufferValidationError) ErrorName() string { return "RawBufferValidationError" }

// Error satisfies the builtin error interface
func (e RawBufferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRawBuffer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RawBufferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RawBufferValidationError{}
