// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/matching/input_matchers/consistent_hashing/v3/consistent_hashing.proto

package consistent_hashingv3

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

// Validate checks the field values on ConsistentHashing with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConsistentHashing) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConsistentHashing with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConsistentHashingMultiError, or nil if none found.
func (m *ConsistentHashing) ValidateAll() error {
	return m.validate(true)
}

func (m *ConsistentHashing) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Threshold

	if m.GetModulo() <= 0 {
		err := ConsistentHashingValidationError{
			field:  "Modulo",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Seed

	if len(errors) > 0 {
		return ConsistentHashingMultiError(errors)
	}
	return nil
}

// ConsistentHashingMultiError is an error wrapping multiple validation errors
// returned by ConsistentHashing.ValidateAll() if the designated constraints
// aren't met.
type ConsistentHashingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsistentHashingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsistentHashingMultiError) AllErrors() []error { return m }

// ConsistentHashingValidationError is the validation error returned by
// ConsistentHashing.Validate if the designated constraints aren't met.
type ConsistentHashingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsistentHashingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsistentHashingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsistentHashingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsistentHashingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsistentHashingValidationError) ErrorName() string {
	return "ConsistentHashingValidationError"
}

// Error satisfies the builtin error interface
func (e ConsistentHashingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsistentHashing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsistentHashingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsistentHashingValidationError{}
