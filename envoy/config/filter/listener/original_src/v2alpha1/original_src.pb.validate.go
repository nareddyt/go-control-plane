// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/listener/original_src/v2alpha1/original_src.proto

package v2alpha1

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

// Validate checks the field values on OriginalSrc with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OriginalSrc) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BindPort

	// no validation rules for Mark

	return nil
}

// OriginalSrcValidationError is the validation error returned by
// OriginalSrc.Validate if the designated constraints aren't met.
type OriginalSrcValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OriginalSrcValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OriginalSrcValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OriginalSrcValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OriginalSrcValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OriginalSrcValidationError) ErrorName() string { return "OriginalSrcValidationError" }

// Error satisfies the builtin error interface
func (e OriginalSrcValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOriginalSrc.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OriginalSrcValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OriginalSrcValidationError{}
