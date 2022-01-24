// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/listener/http_inspector/v3/http_inspector.proto

package http_inspectorv3

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

// Validate checks the field values on HttpInspector with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HttpInspector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpInspector with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HttpInspectorMultiError, or
// nil if none found.
func (m *HttpInspector) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpInspector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return HttpInspectorMultiError(errors)
	}
	return nil
}

// HttpInspectorMultiError is an error wrapping multiple validation errors
// returned by HttpInspector.ValidateAll() if the designated constraints
// aren't met.
type HttpInspectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpInspectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpInspectorMultiError) AllErrors() []error { return m }

// HttpInspectorValidationError is the validation error returned by
// HttpInspector.Validate if the designated constraints aren't met.
type HttpInspectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpInspectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpInspectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpInspectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpInspectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpInspectorValidationError) ErrorName() string { return "HttpInspectorValidationError" }

// Error satisfies the builtin error interface
func (e HttpInspectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpInspector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpInspectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpInspectorValidationError{}
