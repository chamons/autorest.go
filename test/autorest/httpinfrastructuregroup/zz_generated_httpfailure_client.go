//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPFailureClient contains the methods for the HTTPFailure group.
// Don't use this type directly, use NewHTTPFailureClient() instead.
type HTTPFailureClient struct {
	con *Connection
}

// NewHTTPFailureClient creates a new instance of HTTPFailureClient with the specified values.
func NewHTTPFailureClient(con *Connection) *HTTPFailureClient {
	return &HTTPFailureClient{con: con}
}

// GetEmptyError - Get empty error form server
// If the operation fails it returns the *Error error type.
func (client *HTTPFailureClient) GetEmptyError(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (HTTPFailureGetEmptyErrorResponse, error) {
	req, err := client.getEmptyErrorCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureGetEmptyErrorResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPFailureGetEmptyErrorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPFailureGetEmptyErrorResponse{}, client.getEmptyErrorHandleError(resp)
	}
	return client.getEmptyErrorHandleResponse(resp)
}

// getEmptyErrorCreateRequest creates the GetEmptyError request.
func (client *HTTPFailureClient) getEmptyErrorCreateRequest(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (*policy.Request, error) {
	urlPath := "/http/failure/emptybody/error"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyErrorHandleResponse handles the GetEmptyError response.
func (client *HTTPFailureClient) getEmptyErrorHandleResponse(resp *http.Response) (HTTPFailureGetEmptyErrorResponse, error) {
	result := HTTPFailureGetEmptyErrorResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPFailureGetEmptyErrorResponse{}, err
	}
	return result, nil
}

// getEmptyErrorHandleError handles the GetEmptyError error response.
func (client *HTTPFailureClient) getEmptyErrorHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetNoModelEmpty - Get empty response from server
// If the operation fails it returns a generic error.
func (client *HTTPFailureClient) GetNoModelEmpty(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (HTTPFailureGetNoModelEmptyResponse, error) {
	req, err := client.getNoModelEmptyCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureGetNoModelEmptyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPFailureGetNoModelEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPFailureGetNoModelEmptyResponse{}, client.getNoModelEmptyHandleError(resp)
	}
	return client.getNoModelEmptyHandleResponse(resp)
}

// getNoModelEmptyCreateRequest creates the GetNoModelEmpty request.
func (client *HTTPFailureClient) getNoModelEmptyCreateRequest(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (*policy.Request, error) {
	urlPath := "/http/failure/nomodel/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNoModelEmptyHandleResponse handles the GetNoModelEmpty response.
func (client *HTTPFailureClient) getNoModelEmptyHandleResponse(resp *http.Response) (HTTPFailureGetNoModelEmptyResponse, error) {
	result := HTTPFailureGetNoModelEmptyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPFailureGetNoModelEmptyResponse{}, err
	}
	return result, nil
}

// getNoModelEmptyHandleError handles the GetNoModelEmpty error response.
func (client *HTTPFailureClient) getNoModelEmptyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetNoModelError - Get empty error form server
// If the operation fails it returns a generic error.
func (client *HTTPFailureClient) GetNoModelError(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (HTTPFailureGetNoModelErrorResponse, error) {
	req, err := client.getNoModelErrorCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureGetNoModelErrorResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPFailureGetNoModelErrorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPFailureGetNoModelErrorResponse{}, client.getNoModelErrorHandleError(resp)
	}
	return client.getNoModelErrorHandleResponse(resp)
}

// getNoModelErrorCreateRequest creates the GetNoModelError request.
func (client *HTTPFailureClient) getNoModelErrorCreateRequest(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (*policy.Request, error) {
	urlPath := "/http/failure/nomodel/error"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNoModelErrorHandleResponse handles the GetNoModelError response.
func (client *HTTPFailureClient) getNoModelErrorHandleResponse(resp *http.Response) (HTTPFailureGetNoModelErrorResponse, error) {
	result := HTTPFailureGetNoModelErrorResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPFailureGetNoModelErrorResponse{}, err
	}
	return result, nil
}

// getNoModelErrorHandleError handles the GetNoModelError error response.
func (client *HTTPFailureClient) getNoModelErrorHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
