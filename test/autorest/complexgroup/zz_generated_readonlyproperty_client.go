// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ReadonlypropertyClient contains the methods for the Readonlyproperty group.
// Don't use this type directly, use NewReadonlypropertyClient() instead.
type ReadonlypropertyClient struct {
	con *Connection
}

// NewReadonlypropertyClient creates a new instance of ReadonlypropertyClient with the specified values.
func NewReadonlypropertyClient(con *Connection) *ReadonlypropertyClient {
	return &ReadonlypropertyClient{con: con}
}

// GetValid - Get complex types that have readonly properties
// If the operation fails it returns the *Error error type.
func (client *ReadonlypropertyClient) GetValid(ctx context.Context, options *ReadonlypropertyGetValidOptions) (ReadonlypropertyGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return ReadonlypropertyGetValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ReadonlypropertyGetValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ReadonlypropertyGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *ReadonlypropertyClient) getValidCreateRequest(ctx context.Context, options *ReadonlypropertyGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *ReadonlypropertyClient) getValidHandleResponse(resp *azcore.Response) (ReadonlypropertyGetValidResponse, error) {
	result := ReadonlypropertyGetValidResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ReadonlyObj); err != nil {
		return ReadonlypropertyGetValidResponse{}, err
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *ReadonlypropertyClient) getValidHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// PutValid - Put complex types that have readonly properties
// If the operation fails it returns the *Error error type.
func (client *ReadonlypropertyClient) PutValid(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyPutValidOptions) (ReadonlypropertyPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return ReadonlypropertyPutValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ReadonlypropertyPutValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ReadonlypropertyPutValidResponse{}, client.putValidHandleError(resp)
	}
	return ReadonlypropertyPutValidResponse{RawResponse: resp.Response}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *ReadonlypropertyClient) putValidCreateRequest(ctx context.Context, complexBody ReadonlyObj, options *ReadonlypropertyPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *ReadonlypropertyClient) putValidHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
