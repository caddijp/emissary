// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/cluster/aggregate/v2alpha/cluster.proto

package v2alpha

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

// Validate checks the field values on ClusterConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClusterConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClusterConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClusterConfigMultiError, or
// nil if none found.
func (m *ClusterConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ClusterConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetClusters()) < 1 {
		err := ClusterConfigValidationError{
			field:  "Clusters",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ClusterConfigMultiError(errors)
	}
	return nil
}

// ClusterConfigMultiError is an error wrapping multiple validation errors
// returned by ClusterConfig.ValidateAll() if the designated constraints
// aren't met.
type ClusterConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterConfigMultiError) AllErrors() []error { return m }

// ClusterConfigValidationError is the validation error returned by
// ClusterConfig.Validate if the designated constraints aren't met.
type ClusterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterConfigValidationError) ErrorName() string { return "ClusterConfigValidationError" }

// Error satisfies the builtin error interface
func (e ClusterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterConfigValidationError{}
