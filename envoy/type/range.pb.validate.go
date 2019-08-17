// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/range.proto

package envoy_type

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

// Validate checks the field values on Int64Range with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Int64Range) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Start

	// no validation rules for End

	return nil
}

// Int64RangeValidationError is the validation error returned by
// Int64Range.Validate if the designated constraints aren't met.
type Int64RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int64RangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int64RangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int64RangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int64RangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int64RangeValidationError) ErrorName() string { return "Int64RangeValidationError" }

// Error satisfies the builtin error interface
func (e Int64RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt64Range.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int64RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int64RangeValidationError{}

// Validate checks the field values on DoubleRange with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DoubleRange) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Start

	// no validation rules for End

	return nil
}

// DoubleRangeValidationError is the validation error returned by
// DoubleRange.Validate if the designated constraints aren't met.
type DoubleRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleRangeValidationError) ErrorName() string { return "DoubleRangeValidationError" }

// Error satisfies the builtin error interface
func (e DoubleRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleRangeValidationError{}
