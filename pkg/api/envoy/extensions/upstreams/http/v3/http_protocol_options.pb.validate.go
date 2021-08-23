// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/upstreams/http/v3/http_protocol_options.proto

package envoy_extensions_upstreams_http_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on HttpProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "CommonHttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "UpstreamHttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.UpstreamProtocolOptions.(type) {

	case *HttpProtocolOptions_ExplicitHttpConfig_:

		if v, ok := interface{}(m.GetExplicitHttpConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptionsValidationError{
					field:  "ExplicitHttpConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_UseDownstreamProtocolConfig:

		if v, ok := interface{}(m.GetUseDownstreamProtocolConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptionsValidationError{
					field:  "UseDownstreamProtocolConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_AutoConfig:

		if v, ok := interface{}(m.GetAutoConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptionsValidationError{
					field:  "AutoConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return HttpProtocolOptionsValidationError{
			field:  "UpstreamProtocolOptions",
			reason: "value is required",
		}

	}

	return nil
}

// HttpProtocolOptionsValidationError is the validation error returned by
// HttpProtocolOptions.Validate if the designated constraints aren't met.
type HttpProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptionsValidationError) ErrorName() string {
	return "HttpProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptionsValidationError{}

// Validate checks the field values on HttpProtocolOptions_ExplicitHttpConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *HttpProtocolOptions_ExplicitHttpConfig) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ProtocolConfig.(type) {

	case *HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions:

		if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptions_ExplicitHttpConfigValidationError{
					field:  "HttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions:

		if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptions_ExplicitHttpConfigValidationError{
					field:  "Http2ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return HttpProtocolOptions_ExplicitHttpConfigValidationError{
			field:  "ProtocolConfig",
			reason: "value is required",
		}

	}

	return nil
}

// HttpProtocolOptions_ExplicitHttpConfigValidationError is the validation
// error returned by HttpProtocolOptions_ExplicitHttpConfig.Validate if the
// designated constraints aren't met.
type HttpProtocolOptions_ExplicitHttpConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) ErrorName() string {
	return "HttpProtocolOptions_ExplicitHttpConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions_ExplicitHttpConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptions_ExplicitHttpConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptions_ExplicitHttpConfigValidationError{}

// Validate checks the field values on
// HttpProtocolOptions_UseDownstreamHttpConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpProtocolOptions_UseDownstreamHttpConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
				field:  "HttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
				field:  "Http2ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpProtocolOptions_UseDownstreamHttpConfigValidationError is the validation
// error returned by HttpProtocolOptions_UseDownstreamHttpConfig.Validate if
// the designated constraints aren't met.
type HttpProtocolOptions_UseDownstreamHttpConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) ErrorName() string {
	return "HttpProtocolOptions_UseDownstreamHttpConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions_UseDownstreamHttpConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptions_UseDownstreamHttpConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptions_UseDownstreamHttpConfigValidationError{}

// Validate checks the field values on HttpProtocolOptions_AutoHttpConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *HttpProtocolOptions_AutoHttpConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_AutoHttpConfigValidationError{
				field:  "HttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_AutoHttpConfigValidationError{
				field:  "Http2ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpProtocolOptions_AutoHttpConfigValidationError is the validation error
// returned by HttpProtocolOptions_AutoHttpConfig.Validate if the designated
// constraints aren't met.
type HttpProtocolOptions_AutoHttpConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) ErrorName() string {
	return "HttpProtocolOptions_AutoHttpConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions_AutoHttpConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptions_AutoHttpConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptions_AutoHttpConfigValidationError{}
