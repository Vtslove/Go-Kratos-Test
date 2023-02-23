// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/conf/registry.proto

package conf

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

// Validate checks the field values on Registry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Registry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegistryMultiError, or nil
// if none found.
func (m *Registry) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConsul()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConsul()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Consul",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEtcd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEtcd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Etcd",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetZookeeper()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Zookeeper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Zookeeper",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetZookeeper()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Zookeeper",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNacos()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Nacos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Nacos",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNacos()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Nacos",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKubernetes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Kubernetes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Kubernetes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKubernetes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Kubernetes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEureka()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Eureka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Eureka",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEureka()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Eureka",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPolaris()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Polaris",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Polaris",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPolaris()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Polaris",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetServicecomb()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Servicecomb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryValidationError{
					field:  "Servicecomb",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServicecomb()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryValidationError{
				field:  "Servicecomb",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegistryMultiError(errors)
	}

	return nil
}

// RegistryMultiError is an error wrapping multiple validation errors returned
// by Registry.ValidateAll() if the designated constraints aren't met.
type RegistryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistryMultiError) AllErrors() []error { return m }

// RegistryValidationError is the validation error returned by
// Registry.Validate if the designated constraints aren't met.
type RegistryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryValidationError) ErrorName() string { return "RegistryValidationError" }

// Error satisfies the builtin error interface
func (e RegistryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryValidationError{}

// Validate checks the field values on Consul with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Consul) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Consul with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConsulMultiError, or nil if none found.
func (m *Consul) ValidateAll() error {
	return m.validate(true)
}

func (m *Consul) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Scheme

	// no validation rules for Address

	// no validation rules for HealthCheck

	if len(errors) > 0 {
		return ConsulMultiError(errors)
	}

	return nil
}

// ConsulMultiError is an error wrapping multiple validation errors returned by
// Consul.ValidateAll() if the designated constraints aren't met.
type ConsulMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsulMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsulMultiError) AllErrors() []error { return m }

// ConsulValidationError is the validation error returned by Consul.Validate if
// the designated constraints aren't met.
type ConsulValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsulValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsulValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsulValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsulValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsulValidationError) ErrorName() string { return "ConsulValidationError" }

// Error satisfies the builtin error interface
func (e ConsulValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsul.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsulValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsulValidationError{}

// Validate checks the field values on Etcd with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Etcd) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Etcd with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EtcdMultiError, or nil if none found.
func (m *Etcd) ValidateAll() error {
	return m.validate(true)
}

func (m *Etcd) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EtcdMultiError(errors)
	}

	return nil
}

// EtcdMultiError is an error wrapping multiple validation errors returned by
// Etcd.ValidateAll() if the designated constraints aren't met.
type EtcdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EtcdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EtcdMultiError) AllErrors() []error { return m }

// EtcdValidationError is the validation error returned by Etcd.Validate if the
// designated constraints aren't met.
type EtcdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EtcdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EtcdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EtcdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EtcdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EtcdValidationError) ErrorName() string { return "EtcdValidationError" }

// Error satisfies the builtin error interface
func (e EtcdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEtcd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EtcdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EtcdValidationError{}

// Validate checks the field values on ZooKeeper with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ZooKeeper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ZooKeeper with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ZooKeeperMultiError, or nil
// if none found.
func (m *ZooKeeper) ValidateAll() error {
	return m.validate(true)
}

func (m *ZooKeeper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ZooKeeperValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ZooKeeperValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ZooKeeperValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ZooKeeperMultiError(errors)
	}

	return nil
}

// ZooKeeperMultiError is an error wrapping multiple validation errors returned
// by ZooKeeper.ValidateAll() if the designated constraints aren't met.
type ZooKeeperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ZooKeeperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ZooKeeperMultiError) AllErrors() []error { return m }

// ZooKeeperValidationError is the validation error returned by
// ZooKeeper.Validate if the designated constraints aren't met.
type ZooKeeperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZooKeeperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZooKeeperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZooKeeperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZooKeeperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZooKeeperValidationError) ErrorName() string { return "ZooKeeperValidationError" }

// Error satisfies the builtin error interface
func (e ZooKeeperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZooKeeper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZooKeeperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZooKeeperValidationError{}

// Validate checks the field values on Nacos with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Nacos) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Nacos with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NacosMultiError, or nil if none found.
func (m *Nacos) ValidateAll() error {
	return m.validate(true)
}

func (m *Nacos) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Port

	if len(errors) > 0 {
		return NacosMultiError(errors)
	}

	return nil
}

// NacosMultiError is an error wrapping multiple validation errors returned by
// Nacos.ValidateAll() if the designated constraints aren't met.
type NacosMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NacosMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NacosMultiError) AllErrors() []error { return m }

// NacosValidationError is the validation error returned by Nacos.Validate if
// the designated constraints aren't met.
type NacosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NacosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NacosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NacosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NacosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NacosValidationError) ErrorName() string { return "NacosValidationError" }

// Error satisfies the builtin error interface
func (e NacosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNacos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NacosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NacosValidationError{}

// Validate checks the field values on Kubernetes with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Kubernetes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Kubernetes with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KubernetesMultiError, or
// nil if none found.
func (m *Kubernetes) ValidateAll() error {
	return m.validate(true)
}

func (m *Kubernetes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return KubernetesMultiError(errors)
	}

	return nil
}

// KubernetesMultiError is an error wrapping multiple validation errors
// returned by Kubernetes.ValidateAll() if the designated constraints aren't met.
type KubernetesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KubernetesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KubernetesMultiError) AllErrors() []error { return m }

// KubernetesValidationError is the validation error returned by
// Kubernetes.Validate if the designated constraints aren't met.
type KubernetesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KubernetesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KubernetesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KubernetesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KubernetesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KubernetesValidationError) ErrorName() string { return "KubernetesValidationError" }

// Error satisfies the builtin error interface
func (e KubernetesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKubernetes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KubernetesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KubernetesValidationError{}

// Validate checks the field values on Eureka with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Eureka) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Eureka with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EurekaMultiError, or nil if none found.
func (m *Eureka) ValidateAll() error {
	return m.validate(true)
}

func (m *Eureka) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHeartbeatInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EurekaValidationError{
					field:  "HeartbeatInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EurekaValidationError{
					field:  "HeartbeatInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeartbeatInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EurekaValidationError{
				field:  "HeartbeatInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRefreshInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EurekaValidationError{
					field:  "RefreshInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EurekaValidationError{
					field:  "RefreshInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EurekaValidationError{
				field:  "RefreshInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Path

	if len(errors) > 0 {
		return EurekaMultiError(errors)
	}

	return nil
}

// EurekaMultiError is an error wrapping multiple validation errors returned by
// Eureka.ValidateAll() if the designated constraints aren't met.
type EurekaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EurekaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EurekaMultiError) AllErrors() []error { return m }

// EurekaValidationError is the validation error returned by Eureka.Validate if
// the designated constraints aren't met.
type EurekaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EurekaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EurekaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EurekaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EurekaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EurekaValidationError) ErrorName() string { return "EurekaValidationError" }

// Error satisfies the builtin error interface
func (e EurekaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEureka.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EurekaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EurekaValidationError{}

// Validate checks the field values on Polaris with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Polaris) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Polaris with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PolarisMultiError, or nil if none found.
func (m *Polaris) ValidateAll() error {
	return m.validate(true)
}

func (m *Polaris) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Port

	// no validation rules for InstanceCount

	// no validation rules for Namespace

	// no validation rules for Service

	// no validation rules for Token

	if len(errors) > 0 {
		return PolarisMultiError(errors)
	}

	return nil
}

// PolarisMultiError is an error wrapping multiple validation errors returned
// by Polaris.ValidateAll() if the designated constraints aren't met.
type PolarisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PolarisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PolarisMultiError) AllErrors() []error { return m }

// PolarisValidationError is the validation error returned by Polaris.Validate
// if the designated constraints aren't met.
type PolarisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PolarisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PolarisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PolarisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PolarisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PolarisValidationError) ErrorName() string { return "PolarisValidationError" }

// Error satisfies the builtin error interface
func (e PolarisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPolaris.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PolarisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PolarisValidationError{}

// Validate checks the field values on Servicecomb with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Servicecomb) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Servicecomb with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServicecombMultiError, or
// nil if none found.
func (m *Servicecomb) ValidateAll() error {
	return m.validate(true)
}

func (m *Servicecomb) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ServicecombMultiError(errors)
	}

	return nil
}

// ServicecombMultiError is an error wrapping multiple validation errors
// returned by Servicecomb.ValidateAll() if the designated constraints aren't met.
type ServicecombMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServicecombMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServicecombMultiError) AllErrors() []error { return m }

// ServicecombValidationError is the validation error returned by
// Servicecomb.Validate if the designated constraints aren't met.
type ServicecombValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServicecombValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServicecombValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServicecombValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServicecombValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServicecombValidationError) ErrorName() string { return "ServicecombValidationError" }

// Error satisfies the builtin error interface
func (e ServicecombValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServicecomb.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServicecombValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServicecombValidationError{}
