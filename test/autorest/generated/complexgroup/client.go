// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/url"
)

type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// Client - Test Infrastructure for AutoRest
type Client struct {
	u *url.URL
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.
func NewDefaultClient(options *ClientOptions) (*Client, error) {
	return NewClient(DefaultEndpoint, options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, options *ClientOptions) (*Client, error) {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)
	}
	return &Client{u: u, p: p}, nil
}

// BasicOperations returns the BasicOperations associated with this client.
func (client *Client) BasicOperations() BasicOperations {
	return &basicOperations{Client: client}
}

// PrimitiveOperations returns the PrimitiveOperations associated with this client.
func (client *Client) PrimitiveOperations() PrimitiveOperations {
	return &primitiveOperations{Client: client}
}

// ArrayOperations returns the ArrayOperations associated with this client.
func (client *Client) ArrayOperations() ArrayOperations {
	return &arrayOperations{Client: client}
}

// DictionaryOperations returns the DictionaryOperations associated with this client.
func (client *Client) DictionaryOperations() DictionaryOperations {
	return &dictionaryOperations{Client: client}
}

// InheritanceOperations returns the InheritanceOperations associated with this client.
func (client *Client) InheritanceOperations() InheritanceOperations {
	return &inheritanceOperations{Client: client}
}

// PolymorphismOperations returns the PolymorphismOperations associated with this client.
func (client *Client) PolymorphismOperations() PolymorphismOperations {
	return &polymorphismOperations{Client: client}
}

// PolymorphicrecursiveOperations returns the PolymorphicrecursiveOperations associated with this client.
func (client *Client) PolymorphicrecursiveOperations() PolymorphicrecursiveOperations {
	return &polymorphicrecursiveOperations{Client: client}
}

// ReadonlypropertyOperations returns the ReadonlypropertyOperations associated with this client.
func (client *Client) ReadonlypropertyOperations() ReadonlypropertyOperations {
	return &readonlypropertyOperations{Client: client}
}

// FlattencomplexOperations returns the FlattencomplexOperations associated with this client.
func (client *Client) FlattencomplexOperations() FlattencomplexOperations {
	return &flattencomplexOperations{Client: client}
}
