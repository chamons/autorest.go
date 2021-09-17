//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dictionarygroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ConnectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type ConnectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient policy.Transporter
	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []policy.Policy
}

// Connection - Test Infrastructure for AutoRest Swagger BAT
type Connection struct {
	u string
	p runtime.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultConnection creates an instance of the Connection type using the DefaultEndpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewDefaultConnection(options *ConnectionOptions) *Connection {
	return NewConnection(DefaultEndpoint, options)
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewConnection(endpoint string, options *ConnectionOptions) *Connection {
	if options == nil {
		options = &ConnectionOptions{}
	}
	policies := []policy.Policy{}
	if !options.Telemetry.Disabled {
		policies = append(policies, runtime.NewTelemetryPolicy(module, version, &options.Telemetry))
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, runtime.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies, runtime.NewLogPolicy(&options.Logging))
	return &Connection{u: endpoint, p: runtime.NewPipeline(options.HTTPClient, policies...)}
}

// Endpoint returns the connection's endpoint.
func (c *Connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() runtime.Pipeline {
	return c.p
}
